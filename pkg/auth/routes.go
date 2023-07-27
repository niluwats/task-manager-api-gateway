package auth

import (
	"github.com/gin-gonic/gin"
	"github.com/niluwats/api-gateway/pkg/auth/handlers"
)

func RegisterRoutes(r *gin.Engine) *ServiceClient {
	svcClient := &ServiceClient{
		Client: InitServiceClient(),
	}

	a := InitAuthMiddleware(svcClient)

	routes := r.Group("/auth")

	routes.GET("/", func(ctx *gin.Context) {
		ctx.String(200, "Task Manager - API gateway")
	})

	routes.POST("/register", svcClient.Register)
	routes.POST("/login", svcClient.Login)

	routes.Use(a.AuthRequired)
	routes.GET("/user/:ID", svcClient.ViewUser)

	return svcClient
}

func (svc *ServiceClient) Register(ctx *gin.Context) {
	handlers.Register(ctx, svc.Client)
}

func (svc *ServiceClient) Login(ctx *gin.Context) {
	handlers.Login(ctx, svc.Client)
}

func (svc *ServiceClient) ViewUser(ctx *gin.Context) {
	handlers.ViewUser(ctx, svc.Client)
}
