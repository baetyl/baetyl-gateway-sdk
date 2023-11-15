package httpin

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

type HandlerFunc func(c *gin.Context) (any, error)

func Wrapper(handler HandlerFunc) func(c *gin.Context) {
	return func(cc *gin.Context) {
		defer func() {
			if r := recover(); r != nil {
				err, ok := r.(error)
				if !ok {
					err = fmt.Errorf("unknown error: %s", err.Error())
				}
				fmt.Println("handle a panic", err.Error())
				cc.JSON(http.StatusInternalServerError, err.Error())
			}
		}()
		res, err := handler(cc)
		if err != nil {
			fmt.Println("handle a panic", err.Error())
			cc.JSON(http.StatusInternalServerError, err.Error())
			return
		}
		fmt.Println("process success", string(toJson(res)))
		cc.JSON(http.StatusOK, res)
	}
}

func toJson(obj interface{}) []byte {
	data, _ := json.Marshal(obj)
	return data
}
