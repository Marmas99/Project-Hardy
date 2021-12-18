package main

import (
	"project/hardy/internal/api"

	"github.com/sirupsen/logrus"
)

func main() {
	logger := logrus.New()
	logger.Info("Starting server")
	api.StartServer(logger)
}
