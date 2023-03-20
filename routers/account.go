package routers

// AccountRouters ...
func (r Router) AccountRouters() {
	authGroup := r.router.Group("/api/accounts")
	authGroup.GET("/list", r.handler.ListAccounts)
}
