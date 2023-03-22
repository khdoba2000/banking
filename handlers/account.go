package handlers

import (
	"context"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/khdoba2000/banking/configs"
	"github.com/khdoba2000/banking/pkg/http"
	"github.com/khdoba2000/banking/pkg/jwt"
)

// GetAccount gets account of a customer
func (h *Handler) GetAccount(c *gin.Context) {

	ctx, cancel := context.WithTimeout(c.Request.Context(), time.Second*time.Duration(configs.Config().CtxTimeoutSeconds))
	defer cancel()

	customerID, err := jwt.ExtractFromClaims("id", c.Request.Header.Get("Authorization"), []byte(h.cfg.JWTSecretKey))
	if err != nil {
		h.handleResponse(c, StatusFromError(err), err.Error())
		return
	}

	resp, err := h.accountController.GetByOwnerID(
		ctx,
		customerID.(string),
	)
	if err != nil {
		h.handleResponse(c, StatusFromError(err), err.Error())
		return
	}

	h.handleResponse(c, http.OK, resp)
}
