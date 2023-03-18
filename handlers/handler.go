package handlers

import (
	"github.com/gin-gonic/gin"
	"github.com/khdoba/banking/configs"
	authcontroller "github.com/khdoba/banking/controllers/auth"
	"github.com/khdoba/banking/logger"
	"github.com/khdoba/banking/pkg/http"
)

type Handler struct {
	cfg            *configs.Configuration
	log            logger.LoggerI
	authController authcontroller.AuthController
}

// New creates a new Handler
func New(cfg *configs.Configuration, log logger.LoggerI, authController authcontroller.AuthController) Handler {
	return Handler{
		cfg:            cfg,
		log:            log,
		authController: authController,
	}
}

// handleResponse
func (h *Handler) handleResponse(c *gin.Context, status http.Status, data interface{}) {
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

	c.JSON(status.Code, http.Response{
		Status:      status.Status,
		Description: status.Description,
		Data:        data,
	})
}

// func (h *Handler) getOffsetParam(c *gin.Context) (offset int, err error) {
// 	offsetStr := c.DefaultQuery("offset", h.cfg.DefaultOffset)
// 	return strconv.Atoi(offsetStr)
// }

// func (h *Handler) getLimitParam(c *gin.Context) (offset int, err error) {
// 	offsetStr := c.DefaultQuery("limit", h.cfg.DefaultLimit)
// 	return strconv.Atoi(offsetStr)
// }
