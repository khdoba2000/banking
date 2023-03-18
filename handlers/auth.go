package handlers

import (
	"github.com/gin-gonic/gin"
	"github.com/khdoba/banking/entities"
	"github.com/khdoba/banking/pkg/http"
)

// Login loges in the user with the given credentials
func (h *Handler) Login(c *gin.Context) {
	login := entities.LoginReq{}
	err := c.ShouldBindJSON(&login)
	if err != nil {
		h.handleResponse(c, http.BadRequest, err.Error())
		return
	}

	resp, err := h.authController.Login(
		c.Request.Context(),
		login,
	)
	if err != nil {
		h.handleResponse(c, http.InternalServerError, err.Error())
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

	resp, err := h.authController.Signup(
		c.Request.Context(),
		req,
	)
	if err != nil {
		h.handleResponse(c, http.InternalServerError, err.Error())
		return
	}

	h.handleResponse(c, http.OK, resp)
}
