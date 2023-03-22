package handlers

import (
	"context"

	"github.com/gin-gonic/gin"
	"github.com/khdoba2000/banking/constants"
	"github.com/khdoba2000/banking/entities"
	"github.com/khdoba2000/banking/pkg/http"
	"github.com/khdoba2000/banking/pkg/jwt"
)

// CreateExpense
func (h *Handler) CreateExpense(c *gin.Context) {
	expense := entities.Expense{}
	err := c.ShouldBindJSON(&expense)
	if err != nil {
		h.handleResponse(c, http.BadRequest, err.Error())
		return
	}

	err = expense.Validate()
	if err != nil {
		h.handleResponse(c, http.InvalidArgument, err.Error())
		return
	}

	ctx, cancel := context.WithTimeout(c.Request.Context(), constants.ContextTimeoutDuration)
	defer cancel()

	customerID, err := jwt.ExtractFromClaims("id", c.Request.Header.Get("Authorization"), []byte(h.cfg.JWTSecretKey))
	if err != nil {
		h.handleResponse(c, StatusFromError(err), err.Error())
		return
	}

	customerIDString, ok := customerID.(string)
	if !ok {
		h.handleResponse(c, http.InvalidArgument, "customerID of this user is not stringable")
		return
	}

	resp, err := h.transactionController.Create(
		ctx,
		entities.CreateTransactionReq{
			CustomerID:  customerIDString,
			Transaction: &expense,
		},
	)
	if err != nil {
		h.handleResponse(c, StatusFromError(err), err.Error())
		return
	}

	h.handleResponse(c, http.OK, resp)
}

// CreateIncome
func (h *Handler) CreateIncome(c *gin.Context) {
	income := entities.Income{}
	err := c.ShouldBindJSON(&income)
	if err != nil {
		h.handleResponse(c, http.BadRequest, err.Error())
		return
	}

	err = income.Validate()
	if err != nil {
		h.handleResponse(c, http.InvalidArgument, err.Error())
		return
	}

	ctx, cancel := context.WithTimeout(c.Request.Context(), constants.ContextTimeoutDuration)
	defer cancel()

	customerID, err := jwt.ExtractFromClaims("id", c.Request.Header.Get("Authorization"), []byte(h.cfg.JWTSecretKey))
	if err != nil {
		h.handleResponse(c, StatusFromError(err), err.Error())
		return
	}

	customerIDString, ok := customerID.(string)
	if !ok {
		h.handleResponse(c, http.BadRequest, "customerID of this user is not stringable")
		return
	}

	resp, err := h.transactionController.Create(
		ctx,
		entities.CreateTransactionReq{
			CustomerID:  customerIDString,
			Transaction: &income,
		},
	)
	if err != nil {
		h.handleResponse(c, StatusFromError(err), err.Error())
		return
	}

	h.handleResponse(c, http.OK, resp)
}

// CreateTransfer
func (h *Handler) CreateTransfer(c *gin.Context) {
	transfer := entities.Transfer{}
	err := c.ShouldBindJSON(&transfer)
	if err != nil {
		h.handleResponse(c, http.BadRequest, err.Error())
		return
	}

	err = transfer.Validate()
	if err != nil {
		h.handleResponse(c, http.InvalidArgument, err.Error())
		return
	}

	ctx, cancel := context.WithTimeout(c.Request.Context(), constants.ContextTimeoutDuration)
	defer cancel()

	customerID, err := jwt.ExtractFromClaims("id", c.Request.Header.Get("Authorization"), []byte(h.cfg.JWTSecretKey))
	if err != nil {
		h.handleResponse(c, StatusFromError(err), err.Error())
		return
	}

	customerIDString, ok := customerID.(string)
	if !ok {
		h.handleResponse(c, http.BadRequest, "customerID of this user is not stringable")
		return
	}

	resp, err := h.transactionController.Create(
		ctx,
		entities.CreateTransactionReq{
			CustomerID:  customerIDString,
			Transaction: &transfer,
		},
	)
	if err != nil {
		h.handleResponse(c, StatusFromError(err), err.Error())
		return
	}

	h.handleResponse(c, http.OK, resp)
}
