package main

import (
	"encoding/json"
	"io"
	"log"

	"github.com/erice5005/autowebhook"
)

type ExpectedRecv struct {
	Hello string `json:"hello"`
}

func main() {

	ws := autowebhook.NewWebhookService([]autowebhook.HandlerFrame{
		{
			HookPath: "/test1",
			Parser: func(bd io.ReadCloser) (interface{}, error) {
				var out ExpectedRecv
				if err := json.NewDecoder(bd).Decode(&out); err != nil {
					return nil, err
				}
				return out, nil
			},
			Exec: func(i interface{}) {
				ix := i.(ExpectedRecv)
				log.Printf("Hello: %v\n", ix.Hello)
			},
		},
	})
	ws.Init()
	ws.Serve(":8989")
}
