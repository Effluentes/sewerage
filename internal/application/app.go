package application

import (
	"sewerage/internal/telemetry"
	"sewerage/internal/infrastructure/server"
	"sewerage/internal/handlers"
)

func Run() {
	initTelemetry()

	topmux := createMainServer()
	for _, endpoint := range createEndpoints() {
		topmux.CombineServer("/api/", endpoint)
	}

	topmux.RunServer(4000)
}

func createMainServer() *server.HTTPServer {
    return server.NewHTTPServer()
}

func initTelemetry() {
	telemetry.InitLogger()
}

func createEndpoints() []*server.HTTPServer {
    user := handlers.CombineUserHandlers()
    return []*server.HTTPServer{user}
}