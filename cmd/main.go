package main

import (
	"github.com/gin-gonic/gin"
	"github.com/khdoba2000/banking/configs"
	"github.com/khdoba2000/banking/constants"
	accountcontroller "github.com/khdoba2000/banking/controllers/account"
	authcontroller "github.com/khdoba2000/banking/controllers/auth"
	transactioncontroller "github.com/khdoba2000/banking/controllers/transaction"
	"github.com/khdoba2000/banking/handlers"
	"github.com/khdoba2000/banking/logger"
	"github.com/khdoba2000/banking/routers"
	"github.com/khdoba2000/banking/storage"
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
	transactioncontroller := transactioncontroller.NewTransactionController(log, strg)

	//handlers init
	h := handlers.New(
		cfg,
		log,
		authcontroller,
		accountcontroller,
		transactioncontroller,
	)

	//routers
	router := routers.New(h, cfg, log)

	router.Start()

}
