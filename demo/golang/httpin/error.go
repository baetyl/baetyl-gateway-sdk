package httpin

import (
	"bytes"
	"errors"
	"net/http"
	"strings"
	"text/template"
)

var (
	ErrEventTypeNotSupported = errors.New("event type not supported yet")
)

type Status interface {
	Status() int
}

// StatusError error with http status
type StatusError struct {
	e error
	s int
}

func NewStatusError(status, msg string) error {
	return &StatusError{
		e: errors.New(msg),
		s: Code2Status(Code(status)),
	}
}

func (s *StatusError) Status() int {
	return s.s
}

func (s *StatusError) Error() string {
	return s.e.Error()
}

// Code error code
type Code string

func (c Code) String() string {
	if msg, ok := templates[c]; ok {
		return msg
	}
	return templates[ErrUnknown]
}

const (
	ErrUnknown  = "ErrUnknown"
	ErrNotFound = "ErrNotFound"
	ErrDenied   = "ErrDenied"
	ErrInvalid  = "ErrInvalid"
	ErrRunning  = "ErrRunning"
)

var (
	templates = map[Code]string{
		ErrUnknown:  "There is a unknown error. {{if .error}} {{.error}}{{end}}",
		ErrNotFound: "The {{if .type}} {{.type}}{{end}} resource {{if .name}} {{.name}}{{end}} is not found",
		ErrDenied:   "The request access is denied. {{if .error}} {{.error}}{{end}}",
		ErrInvalid:  "The request is invalid. {{if .name}} {{.name}}{{end}} {{if .error}} {{.error}}{{end}}",
		ErrRunning:  "The execution process encountered an error. {{if .name}} {{.name}}{{end}} {{if .error}} {{.error}}{{end}}",
	}
)

func Code2Status(c Code) int {
	switch c {
	case ErrDenied:
		return http.StatusUnauthorized
	case ErrNotFound:
		return http.StatusNotFound
	case ErrInvalid, ErrRunning:
		return http.StatusBadRequest
	default:
		return http.StatusInternalServerError
	}
}

type Filed struct {
	k string
	v any
}

func F(k string, v any) *Filed {
	return &Filed{k, v}
}

func E(c Code, fs ...*Filed) error {
	m := c.String()
	if strings.Contains(m, "{{") {
		vs := map[string]any{}
		for _, f := range fs {
			vs[f.k] = f.v
		}
		t, err := template.New(string(c)).Option("missingkey=zero").Parse(m)
		if err != nil {
			return err
		}
		b := bytes.NewBuffer(nil)
		err = t.Execute(b, vs)
		if err != nil {
			return err
		}
		m = b.String()
	}
	return NewStatusError(string(c), m)
}
