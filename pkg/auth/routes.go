package auth

import (
	"github.com/gin-gonic/gin"
	"github.com/niluwats/api-gateway/pkg/auth/handlers"

	_ "github.com/niluwats/api-gateway/api/docs"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

// @title           Task Manager API Gateway
// @swagger 		"2.0"
// @version         1.0
// @description     This is a REST server created with Gin.
// @termsOfService  http://swagger.io/terms/

// @contact.name   API Support
// @contact.url    http://www.swagger.io/support
// @contact.email  support@swagger.io

// @license.name  Apache 2.0
// @license.url   http://www.apache.org/licenses/LICENSE-2.0.html

// @host      localhost:8080
// @BasePath  /api/v1

// @securityDefinitions.apiKey  BearerAuth
// @in header
// @name Authorization

// @externalDocs.description  OpenAPI
// @externalDocs.url          https://swagger.io/resources/open-api/
func RegisterRoutes(r *gin.Engine) *ServiceClient {
	url := ginSwagger.URL("http://localhost:8080/swagger/doc.json")
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler, url))

	svcClient := &ServiceClient{
		Client: InitServiceClient(),
	}

	a := InitAuthMiddleware(svcClient)

	routes := r.Group("/api/v1")

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
