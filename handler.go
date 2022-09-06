package autowebhook

import (
	"io"
	"net/http"
)

type HandlerFrame struct {
	Exec     func(interface{})
	Parser   func(io.ReadCloser) (interface{}, error)
	HookPath string
}

func (hf HandlerFrame) GetEndHandler() func(w http.ResponseWriter, r *http.Request) { //TODO: add logger callout
	return func(w http.ResponseWriter, r *http.Request) {
		parsed, err := hf.Parser(r.Body)
		if err != nil {
			w.WriteHeader(500)
			return
		}
		go hf.Exec(parsed)
		w.WriteHeader(200)
	}
}

func NewHandlerFram(path string, exec func(interface{}), parser func(io.ReadCloser) (interface{}, error)) HandlerFrame {
	return HandlerFrame{
		HookPath: path,
		Parser:   parser,
		Exec:     exec,
	}
}
