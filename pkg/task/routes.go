package task

import (
	"github.com/gin-gonic/gin"
	"github.com/niluwats/api-gateway/pkg/auth"
	"github.com/niluwats/api-gateway/pkg/task/handlers"
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
func RegisterRoutes(r *gin.Engine, authSvc *auth.ServiceClient) {
	url := ginSwagger.URL("http://localhost:8080/swagger/doc.json")
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler, url))

	svcClient := &ServiceClient{Client: InitServiceClient()}

	a := auth.InitAuthMiddleware(authSvc)

	routes := r.Group("/task")
	routes.Use(a.AuthRequired)
	routes.POST("/", svcClient.createProject)
	routes.PUT("/:ID", svcClient.updateProject)
	routes.DELETE("/:ID", svcClient.removeProject)
	routes.GET("/", svcClient.viewAllProjects)
	routes.GET("/:ID", svcClient.viewProject)
}

func (svc *ServiceClient) createProject(ctx *gin.Context) {
	handlers.CreateProject(ctx, svc.Client)
}

func (svc *ServiceClient) updateProject(ctx *gin.Context) {
	handlers.UpdateProject(ctx, svc.Client)
}

func (svc *ServiceClient) removeProject(ctx *gin.Context) {
	handlers.RemoveProject(ctx, svc.Client)
}

func (svc *ServiceClient) viewProject(ctx *gin.Context) {
	handlers.ViewProject(ctx, svc.Client)
}

func (svc *ServiceClient) viewAllProjects(ctx *gin.Context) {
	handlers.ViewAllProjects(ctx, svc.Client)
}
