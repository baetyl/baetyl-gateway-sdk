package httpin

import (
	"bytes"
	"context"
	"io"
	"net/http"
	"strconv"
	"time"

	"github.com/baetyl/baetyl-go/v2/utils"
	"github.com/gin-gonic/gin"
)

type Server struct {
	cfg    *ServerConfig
	svr    *http.Server
	router *gin.Engine
	api    *API
}

func NewServer(cfg *ServerConfig) (*Server, error) {
	L().Debug("NewServer function")

	s := &Server{
		cfg: cfg,
	}
	defer L().Debug("NewServer function end", s.cfg, s.svr, s)

	router := gin.New()
	server := &http.Server{
		Addr:           cfg.Port,
		Handler:        router,
		ReadTimeout:    cfg.ReadTimeout,
		WriteTimeout:   cfg.WriteTimeout,
		MaxHeaderBytes: 1 << 20,
	}
	if cfg.Certificate.Cert != "" &&
		cfg.Certificate.Key != "" {
		t, errTLS := utils.NewTLSConfigServer(utils.Certificate{
			Cert: cfg.Certificate.Cert,
			Key:  cfg.Certificate.Key,
		})
		if errTLS != nil {
			return nil, errTLS
		}
		server.TLSConfig = t
	}
	return s, nil
}

func (s *Server) Run() {
	if s.svr.TLSConfig == nil {
		if err := s.svr.ListenAndServe(); err != nil {
			L().Debug("init server http stopped", err)
		}
	} else {
		if err := s.svr.ListenAndServeTLS("", ""); err != nil {
			L().Debug("init server https stopped", err)
		}
	}
}

func (s *Server) InitRoute() {
	s.router.GET("/health", func(c *gin.Context) {
		c.JSON(200, true)
	})

	s.router.Use(s.LoggerHandlerFunc)
	v1 := s.router.Group("v1")
	{
		report := v1.Group("/report")
		report.POST("", Wrapper(s.api.Report))
	}
}

func (s *Server) GetRoute() *gin.Engine {
	return s.router
}

func (s *Server) SetAPI(api *API) {
	s.api = api
}

func (s *Server) LoggerHandlerFunc(c *gin.Context) {
	L().Debug("logger handler start request",
		c.Request.Method,
		c.Request.URL.Path,
		c.Request.Host,
		c.Request.Header,
		c.ClientIP(),
	)
	if c.Request.Header.Get("Content-type") == "application/json" && c.Request.Body != nil {
		if buf, err := io.ReadAll(c.Request.Body); err == nil {
			c.Request.Body = io.NopCloser(bytes.NewReader(buf[:]))
			L().Debug("logger handler request body",
				string(buf),
			)
		}
	}
	start := time.Now()
	c.Next()
	L().Debug("logger handler finish request",
		strconv.Itoa(c.Writer.Status()),
		time.Since(start),
		c.Writer.Size(),
	)
}

func (s *Server) Close() error {
	if s.svr != nil {
		ctx, cancel := context.WithTimeout(context.Background(), s.cfg.ShutdownTime)
		err := s.svr.Shutdown(ctx)
		cancel()
		return err
	}
	return nil
}
