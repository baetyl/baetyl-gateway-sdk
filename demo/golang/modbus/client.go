package modbus

import (
	"fmt"

	"github.com/baetyl/baetyl-go/v2/errors"
	md "github.com/goburrow/modbus"
)

type Mode string

type Handler interface {
	md.ClientHandler
	Connect() error
	Close() error
}

type mbClient struct {
	md.Client
	Handler
}

func NewClient(cfg SlaveConfig) (*mbClient, error) {
	var cli mbClient
	switch Mode(cfg.Mode) {
	case ModeTCP:
		// Modbus TCP
		h := md.NewTCPClientHandler(cfg.Address)
		h.SlaveId = cfg.Id
		h.Timeout = cfg.Timeout
		h.IdleTimeout = cfg.IdleTimeout
		cli.Handler = h
	case ModeRTU:
		// Modbus RTU
		h := md.NewRTUClientHandler(cfg.Address)
		h.BaudRate = cfg.BaudRate
		h.DataBits = cfg.DataBits
		h.Parity = cfg.Parity
		h.StopBits = cfg.StopBits
		h.SlaveId = cfg.Id
		h.Timeout = cfg.Timeout
		h.IdleTimeout = cfg.IdleTimeout
		cli.Handler = h
	default:
		return nil, errors.Errorf("method not supported")
	}
	return &cli, nil
}

func (m *mbClient) Reconnect() error {
	if err := m.Close(); err != nil {
		return err
	}
	return m.Connect()
}

func (m *mbClient) Connect() error {
	err := m.Handler.Connect()
	if err != nil {
		return fmt.Errorf("failed to connect: %w", err)
	}
	m.Client = md.NewClient(m.Handler)
	return nil
}

func (m *mbClient) Close() error {
	err := m.Handler.Close()
	if err != nil {
		return fmt.Errorf("failed to close client: %w", err)
	}
	return nil
}
