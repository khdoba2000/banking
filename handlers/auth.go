package handlers

import (
	"context"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/khdoba2000/banking/configs"
	"github.com/khdoba2000/banking/entities"
	"github.com/khdoba2000/banking/pkg/http"
	"github.com/khdoba2000/banking/pkg/jwt"
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

	phoneNumberInToken, err := jwt.ExtractFromClaims("phone_number", c.Request.Header.Get("Authorization"), []byte(h.cfg.JWTSecretKey))
	if err != nil {
		h.handleResponse(c, StatusFromError(err), err.Error())
		return
	}
	if req.PhoneNumber != phoneNumberInToken {
		// if the phone number did not match the one in the token generated in verifyCode phase
		h.handleResponse(c, http.Forbidden, "invalid phone number")
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

// SendCode ...
func (h *Handler) SendCode(c *gin.Context) {
	req := entities.SendCodeReq{}
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

	err = h.authController.SendCode(
		ctx,
		req,
	)
	if err != nil {
		h.handleResponse(c, StatusFromError(err), err.Error())
		return
	}

	h.handleResponse(c, http.OK, "success")
}

// VerifyCode ...
func (h *Handler) VerifyCode(c *gin.Context) {
	req := entities.VerifyCodeReq{}
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

	resp, err := h.authController.VerifyCode(
		ctx,
		req,
	)
	if err != nil {
		h.handleResponse(c, StatusFromError(err), err.Error())
		return
	}

	h.handleResponse(c, http.OK, resp)
}
