package main

import (
	"github.com/gin-gonic/gin"
	"github.com/khdoba/banking/configs"
	"github.com/khdoba/banking/constants"
	accountcontroller "github.com/khdoba/banking/controllers/account"
	authcontroller "github.com/khdoba/banking/controllers/auth"
	"github.com/khdoba/banking/handlers"
	"github.com/khdoba/banking/logger"
	"github.com/khdoba/banking/routers"
	"github.com/khdoba/banking/storage"
)

func main() {

	//configuration settings
	cfg := configs.Config()

	// take environment from config then set gin mode according to it
	switch cfg.Environment {
	case constants.DebugMode:
		gin.SetMode(gin.DebugMode)
	case constants.TestMode:
		gin.SetMode(gin.TestMode)
	default:
		gin.SetMode(gin.ReleaseMode)
	}
	//logger
	log := logger.NewLogger(cfg.AppName, cfg.LogLevel)
	defer logger.Cleanup(log)

	//storage init
	strg := storage.New(cfg)

	//controllers init
	authcontroller := authcontroller.NewAuthController(log, strg)
	accountcontroller := accountcontroller.NewAccountController(log, strg)

	//handlers init
	h := handlers.New(cfg, log, authcontroller, accountcontroller)

	//routers
	router := routers.New(h, cfg, log)

	router.Start()

}
