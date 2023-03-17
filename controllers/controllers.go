package controllers

import (
	authcontroller "github.com/khdoba/banking/controllers/auth"
	"github.com/khdoba/banking/logger"
	"github.com/khdoba/banking/storage"
)

// Controller
type Controller struct {
	Auth authcontroller.AuthController
}

// New creates a new Controller
func New(log logger.LoggerI, storage storage.Storage) Controller {
	return Controller{
		Auth: authcontroller.NewAuthController(log, storage),
	}
}
