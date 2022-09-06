package autowebhook

import (
	"net/http"

	"github.com/gorilla/mux"
)

type WebhookService struct {
	Handlers []HandlerFrame
	r        *mux.Router
}

func NewWebhookService(handlers []HandlerFrame) WebhookService {
	return WebhookService{
		r:        mux.NewRouter(),
		Handlers: handlers,
	}
}

func (ws WebhookService) Init() {
	for _, hx := range ws.Handlers {
		ws.r.HandleFunc(hx.HookPath, hx.GetEndHandler())
	}
}

func (ws WebhookService) Serve(addr string) error {
	if err := http.ListenAndServe(addr, ws.r); err != nil {
		return err
	}
	return nil
}
