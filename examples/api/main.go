package main

import (
	"encoding/json"
	"io"
	"log"

	"github.com/erice5005/autowebhook"
)

type SampleMessage struct {
	Content string `json:"content"`
}

type SampleHandler struct {
	Messages []string
}

func (sh *SampleHandler) AddMessage(ix SampleMessage) {
	sh.Messages = append(sh.Messages, ix.Content)
	log.Printf("Messages: %+v\n", sh.Messages)
}

func main() {
	sh := &SampleHandler{
		Messages: make([]string, 0),
	}
	handlers := []autowebhook.HandlerFrame{
		{
			HookPath: "/test1",
			Parser:   ParseToSampleMessage,
			Exec: func(i interface{}) {
				sh.AddMessage(i.(SampleMessage))
			},
		},
	}
	ws := autowebhook.NewWebhookService(handlers)
	ws.Init()
	ws.Serve(":8989")
}

func ParseToSampleMessage(rc io.ReadCloser) (interface{}, error) {
	var out SampleMessage
	if err := json.NewDecoder(rc).Decode(&out); err != nil {
		return SampleMessage{}, err
	}
	return out, nil
}
