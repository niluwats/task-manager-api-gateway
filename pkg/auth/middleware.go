package auth

import (
	"context"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/niluwats/api-gateway/pkg/auth/pb"
)

type AuthMiddlewareConfig struct {
	service *ServiceClient
}

func InitAuthMiddleware(service *ServiceClient) AuthMiddlewareConfig {
	return AuthMiddlewareConfig{service}
}

func (c *AuthMiddlewareConfig) AuthRequired(ctx *gin.Context) {
	authorization := ctx.Request.Header.Get("authorization")

	if authorization == "" {
		ctx.AbortWithStatus(http.StatusUnauthorized)
		return
	}

	token := strings.Split(authorization, "Bearer ")

	if len(token) < 2 {
		ctx.AbortWithStatus(http.StatusUnauthorized)
		return
	}

	res, err := c.service.Client.ValidateToken(context.Background(), &pb.ValidateTokenRequest{Token: token[1]})

	if err != nil || res.Status != http.StatusOK {
		ctx.AbortWithStatus(http.StatusUnauthorized)
		return
	}

	ctx.Set("userID", res.UserID)
	ctx.Next()
}
