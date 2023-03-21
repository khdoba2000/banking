package routers

// TransactionRouters ...
func (r Router) TransactionRouters() {
	authGroup := r.router.Group("/api/transactions")
	// authGroup.POST("/create", r.handler.CreateTransaction)
	authGroup.POST("/expense", r.handler.CreateExpense)
	authGroup.POST("/income", r.handler.CreateIncome)
	authGroup.POST("/transfer", r.handler.CreateTransfer)
}
