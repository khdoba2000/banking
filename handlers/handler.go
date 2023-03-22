package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/khdoba2000/banking/configs"
	accountcontroller "github.com/khdoba2000/banking/controllers/account"
	authcontroller "github.com/khdoba2000/banking/controllers/auth"
	transactioncontroller "github.com/khdoba2000/banking/controllers/transaction"
	"github.com/khdoba2000/banking/logger"
	e "github.com/khdoba2000/banking/pkg/errors"

	httppkg "github.com/khdoba2000/banking/pkg/http"
)

type Handler struct {
	cfg                   *configs.Configuration
	log                   logger.LoggerI
	authController        authcontroller.AuthController
	accountController     accountcontroller.AccountController
	transactionController transactioncontroller.TransactionController
}

// New creates a new Handler
func New(
	cfg *configs.Configuration,
	log logger.LoggerI,
	authController authcontroller.AuthController,
	accountController accountcontroller.AccountController,
	transactionController transactioncontroller.TransactionController,
) Handler {
	return Handler{
		cfg:                   cfg,
		log:                   log,
		authController:        authController,
		accountController:     accountController,
		transactionController: transactionController,
	}
}

// handleResponse
func (h *Handler) handleResponse(c *gin.Context, status httppkg.Status, data interface{}) {
	switch code := status.Code; {
	case code < 300:
		h.log.Info(
			"---Response--->",
			logger.Int("code", status.Code),
			logger.String("status", status.Status),
			logger.Any("description", status.Description),
			// logger.Any("data", data),
		)
	case code < 400:
		h.log.Warn(
			"!!!Response--->",
			logger.Int("code", status.Code),
			logger.String("status", status.Status),
			logger.Any("description", status.Description),
			logger.Any("data", data),
		)
	default:
		h.log.Error(
			"!!!Response--->",
			logger.Int("code", status.Code),
			logger.String("status", status.Status),
			logger.Any("description", status.Description),
			logger.Any("data", data),
		)
	}

	c.JSON(status.Code, httppkg.Response{
		Status:      status.Status,
		Description: status.Description,
		Data:        data,
	})
}

// StatusFromError ...
func StatusFromError(err error) httppkg.Status {
	if err == nil {
		return httppkg.OK
	}

	code, ok := e.ExtractStatusCode(err)
	if !ok || code == http.StatusInternalServerError {
		return httppkg.Status{
			Code:        http.StatusInternalServerError,
			Status:      "INTERNAL_SERVER_ERROR",
			Description: err.Error(),
		}
	} else if code == http.StatusNotFound {
		return httppkg.Status{
			Code:        http.StatusNotFound,
			Status:      "NOT_FOUND",
			Description: err.Error(),
		}
	} else if code == http.StatusBadRequest {
		return httppkg.Status{
			Code:        http.StatusBadRequest,
			Status:      "BAD_REQUEST",
			Description: err.Error(),
		}
	} else if code == http.StatusForbidden {
		return httppkg.Status{
			Code:        http.StatusForbidden,
			Status:      "FORBIDDEN",
			Description: err.Error(),
		}
	} else if code == http.StatusUnauthorized {
		return httppkg.Status{
			Code:        http.StatusUnauthorized,
			Status:      "FORBIDDEN",
			Description: err.Error(),
		}
	} else {
		return httppkg.Status{
			Code:        http.StatusInternalServerError,
			Status:      "INTERNAL_SERVER_ERROR",
			Description: err.Error(),
		}
	}

}
