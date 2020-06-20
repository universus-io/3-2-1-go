package calculator

import (
	"log"
	"net/http"
	"strconv"
	"sync"
)

type Server struct {
	Port        int
	pingHandler *PingHandler
	expHandler  *ExpHandler
	latch       *sync.WaitGroup
}

func NewServer(port int) *Server {

	calculator := NewCalculator()
	expHandler := NewExpHandler(calculator)
	pingHandler := NewPingHandler()

	return &Server{port, pingHandler, expHandler, &sync.WaitGroup{}}
}

func (server *Server) Start() *sync.WaitGroup {

	server.latch.Add(1)

	http.HandleFunc("/ping", server.pingHandler.Handle)

	http.HandleFunc("/exp/", server.expHandler.Handle)

	address := ":" + strconv.Itoa(server.Port)

	go func() {
		_ = http.ListenAndServe(address, nil)
	}()

	log.Print("[Server] Started")

	return server.latch
}

func (server *Server) Stop() {

	server.latch.Done()

	log.Print("[Server] Stopped")
}
