package routers

func (r Router) AuthRouters() {
	authGroup := r.router.Group("/api/auth")
	authGroup.POST("/login", r.handler.Login)
	authGroup.POST("/sign-up", r.handler.SignUp)
}
