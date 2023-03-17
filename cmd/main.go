package main

import (
	"github.com/khdoba/banking/configs"
	"github.com/khdoba/banking/controllers"
	"github.com/khdoba/banking/logger"
	"github.com/khdoba/banking/storage"
)

func main() {

	cfg := configs.Config()

	strg := storage.New(cfg)

	logger := logger.NewLogger(cfg.AppName, cfg.LogLevel)

	_ = controllers.New(logger, strg)

}
