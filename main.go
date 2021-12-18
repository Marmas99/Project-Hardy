package main

import (
	"io"
	"net/http"

	"github.com/sirupsen/logrus"
)

func main() {
	logger := logrus.New()
	logger.Info("Starting server")
	startServer(logger)
}

func startServer(logger *logrus.Logger) {
	logger.Info("Initiating handlers")
	http.HandleFunc("/hello", helloServer)

	logger.Info("Server listening...")
	http.ListenAndServe(":3030", nil)
}

func helloServer(writer http.ResponseWriter, request *http.Request) {
	io.WriteString(writer, "hello from server")
}
