package handlers

import (
	"context"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/khdoba/banking/configs"
	"github.com/khdoba/banking/pkg/http"
	"github.com/khdoba/banking/pkg/jwt"
)

// ListAccounts lists accounts of a given customer
func (h *Handler) ListAccounts(c *gin.Context) {

	ctx, cancel := context.WithTimeout(c.Request.Context(), time.Second*time.Duration(configs.Config().CtxTimeoutSeconds))
	defer cancel()

	customerID, err := jwt.ExtractFromClaims("id", c.Request.Header.Get("Authorization"), []byte(h.cfg.JWTSecretKey))
	if err != nil {
		h.handleResponse(c, StatusFromError(err), err.Error())
		return
	}

	resp, err := h.accountController.ListByOwnerID(
		ctx,
		customerID.(string),
	)
	if err != nil {
		h.handleResponse(c, StatusFromError(err), err.Error())
		return
	}

	h.handleResponse(c, http.OK, resp)
}
