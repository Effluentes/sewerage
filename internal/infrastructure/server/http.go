package server

import (
	"log"
	"log/slog"
	"net/http"
	"strconv"
	"strings"
)

type HttpMethod string

type HTTPHandler http.HandlerFunc

type HTTPServer struct {
	mux *http.ServeMux
}

func NewHTTPServer() *HTTPServer {
	return &HTTPServer{
		mux: http.NewServeMux(),
	}
}

func (server *HTTPServer) HandleFunc(url string, handler HTTPHandler) {
	server.mux.HandleFunc(url, handler)
}

func (server *HTTPServer) RunServer(port int) {
	err := http.ListenAndServe(":"+strconv.Itoa(port), server.mux)
	log.Fatal(err)
}

func (server *HTTPServer) CombineServer(prefixToSubMux string, toAddMux *HTTPServer) {
	slog.Info("Add new submux", "prefixToSubMux", prefixToSubMux)
	server.mux.Handle(prefixToSubMux, http.StripPrefix(strings.TrimRight(prefixToSubMux, "/"), toAddMux.mux))
}
