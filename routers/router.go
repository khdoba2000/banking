package routers

import (
	"github.com/gin-gonic/gin"
	"github.com/khdoba/banking/configs"
	"github.com/khdoba/banking/handlers"
	"github.com/khdoba/banking/logger"
	"github.com/khdoba/banking/middlewares"
	"go.uber.org/zap"
)

type Router struct {
	handler handlers.Handler
	config  *configs.Configuration
	router  *gin.Engine
	logger  logger.LoggerI
}

// New creates a new router
func New(h handlers.Handler, cfg *configs.Configuration, logger logger.LoggerI) Router {
	r := gin.New()

	return Router{
		handler: h,
		router:  r,
		logger:  logger,
		config:  cfg,
	}
}

func (r Router) Start() {

	r.router.Use(gin.Logger(), gin.Recovery())
	r.router.Use(middlewares.CustomCORSMiddleware())

	casbinJWTRoleAuthorizer, err := middlewares.NewCasbinJWTRoleAuthorizer(r.config, r.logger)
	if err != nil {
		r.logger.Fatal("Could not initialize Cabin JWT Role Authorizer", zap.Error(err))
	}
	r.router.Use(casbinJWTRoleAuthorizer.Middleware())

	r.AuthRouters()
	r.AccountRouters()
	// r.CustomerRouters()

	r.logger.Info("HTTP: Server being started...", logger.String("port", r.config.HTTPPort))

	r.router.Run(r.config.HTTPPort)
}
