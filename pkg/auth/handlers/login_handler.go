package handlers

import (
	"context"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/niluwats/api-gateway/pkg/auth/pb"
)

type LoginRequest struct {
	Email    string `json:"email" binding:"required"`
	Password string `json:"password" binding:"required"`
}

func Login(ctx *gin.Context, c pb.AuthServiceClient) {
	request := LoginRequest{}

	if err := ctx.ShouldBindJSON(&request); err != nil {
		ctx.JSON(http.StatusBadRequest, ErrResponse{Error: err.Error()})
		return
	}

	res, _ := c.Login(context.Background(), &pb.LoginRequest{Email: request.Email, Password: request.Password})

	ctx.JSON(int(res.Status), &res)
}
