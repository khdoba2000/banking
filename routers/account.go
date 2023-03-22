package routers

// AccountRouters ...
func (r Router) AccountRouters() {
	authGroup := r.router.Group("/api/accounts")
	authGroup.GET("/get-balance", r.handler.GetAccount)
}
