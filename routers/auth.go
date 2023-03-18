package routers

func (r Router) AuthRouters() {
	authGroup := r.router.Group("/api/auth")
	authGroup.POST("/send-code", r.handler.SendCode)
	authGroup.POST("/verify-code", r.handler.VerifyCode)
	authGroup.POST("/sign-up", r.handler.SignUp)
	authGroup.POST("/login", r.handler.Login)
}
