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

// // CreateTransaction
// func (h *Handler) CreateTransaction(c *gin.Context) {
// 	req := entities.Transaction{}
// 	err := c.ShouldBindJSON(&req)
// 	if err != nil {
// 		h.handleResponse(c, http.BadRequest, err.Error())
// 		return
// 	}

// 	// err = req.Validate()
// 	// if err != nil {
// 	// 	h.handleResponse(c, http.BadRequest, err.Error())
// 	// 	return
// 	// }

// 	ctx, cancel := context.WithTimeout(c.Request.Context(), time.Second*time.Duration(configs.Config().CtxTimeoutSeconds))
// 	defer cancel()

// 	// customerID, err := jwt.ExtractFromClaims("id", c.Request.Header.Get("Authorization"), []byte(h.cfg.JWTSecretKey))
// 	// if err != nil {
// 	// 	h.handleResponse(c, StatusFromError(err), err.Error())
// 	// 	return
// 	// }

// 	// customerIDString := customerID.(string)
// 	switch req.TypeID {
// 	case constants.TopupTransactionID:
// 		// req.AccountToID = customerIDString
// 	case constants.WithdrawTransactionID:
// 		// req.AccountFromID = customerIDString
// 	}

// 	resp, err := h.transactionController.Create(
// 		ctx,
// 		&req,
// 	)
// 	if err != nil {
// 		h.handleResponse(c, StatusFromError(err), err.Error())
// 		return
// 	}

// 	h.handleResponse(c, http.OK, resp)
// }

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
		h.handleResponse(c, http.BadRequest, err.Error())
		return
	}

	ctx, cancel := context.WithTimeout(c.Request.Context(), time.Second*time.Duration(configs.Config().CtxTimeoutSeconds))
	defer cancel()

	customerID, err := jwt.ExtractFromClaims("id", c.Request.Header.Get("Authorization"), []byte(h.cfg.JWTSecretKey))
	if err != nil {
		h.handleResponse(c, StatusFromError(err), err.Error())
		return
	}

	customerIDString := customerID.(string)

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
		h.handleResponse(c, http.BadRequest, err.Error())
		return
	}

	ctx, cancel := context.WithTimeout(c.Request.Context(), time.Second*time.Duration(configs.Config().CtxTimeoutSeconds))
	defer cancel()

	customerID, err := jwt.ExtractFromClaims("id", c.Request.Header.Get("Authorization"), []byte(h.cfg.JWTSecretKey))
	if err != nil {
		h.handleResponse(c, StatusFromError(err), err.Error())
		return
	}

	customerIDString := customerID.(string)

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
		h.handleResponse(c, http.BadRequest, err.Error())
		return
	}

	ctx, cancel := context.WithTimeout(c.Request.Context(), time.Second*time.Duration(configs.Config().CtxTimeoutSeconds))
	defer cancel()

	customerID, err := jwt.ExtractFromClaims("id", c.Request.Header.Get("Authorization"), []byte(h.cfg.JWTSecretKey))
	if err != nil {
		h.handleResponse(c, StatusFromError(err), err.Error())
		return
	}

	customerIDString := customerID.(string)

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
