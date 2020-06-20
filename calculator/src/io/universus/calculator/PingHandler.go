package calculator

import (
	"fmt"
	"log"
	"net/http"
)

type PingHandler struct {
}

func NewPingHandler() *PingHandler {
	return &PingHandler{}
}

func (h *PingHandler) Handle(writer http.ResponseWriter, _ *http.Request) {

	log.Print("[PingHandler] Handle() >")

	_, _ = fmt.Fprint(writer, "{Ok}")

	log.Print("[PingHandler] Handle() <")
}
