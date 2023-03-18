package handlers

import (
	"context"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/khdoba/banking/configs"
	"github.com/khdoba/banking/entities"
	"github.com/khdoba/banking/pkg/http"
)

// Login loges in the user with the given credentials
func (h *Handler) Login(c *gin.Context) {
	loginReq := entities.LoginReq{}
	err := c.ShouldBindJSON(&loginReq)
	if err != nil {
		h.handleResponse(c, http.BadRequest, err.Error())
		return
	}

	err = loginReq.Validate()
	if err != nil {
		h.handleResponse(c, http.BadRequest, err.Error())
		return
	}

	ctx, cancel := context.WithTimeout(c.Request.Context(), time.Second*time.Duration(configs.Config().CtxTimeoutSeconds))
	defer cancel()

	resp, err := h.authController.Login(
		ctx,
		loginReq,
	)
	if err != nil {
		h.handleResponse(c, StatusFromError(err), err.Error())
		return
	}

	h.handleResponse(c, http.OK, resp)
}

// SignUp ...
func (h *Handler) SignUp(c *gin.Context) {
	req := entities.SignupReq{}
	err := c.ShouldBindJSON(&req)
	if err != nil {
		h.handleResponse(c, http.BadRequest, err.Error())
		return
	}

	err = req.Validate()
	if err != nil {
		h.handleResponse(c, http.BadRequest, err.Error())
		return
	}

	ctx, cancel := context.WithTimeout(c.Request.Context(), time.Second*time.Duration(configs.Config().CtxTimeoutSeconds))
	defer cancel()

	resp, err := h.authController.Signup(
		ctx,
		req,
	)
	if err != nil {
		h.handleResponse(c, StatusFromError(err), err.Error())
		return
	}

	h.handleResponse(c, http.OK, resp)
}
