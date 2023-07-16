package handlers

import (
	"context"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/niluwats/api-gateway/pkg/auth/pb"
)

type ErrResponse struct {
	Error string `json:"error"`
}

type RegisterRequest struct {
	Email    string `json:"email" binding:"required,email"`
	Password string `json:"password" binding:"required,gte=8,lte=32"`
}

func Register(ctx *gin.Context, c pb.AuthServiceClient) {
	request := RegisterRequest{}

	if err := ctx.ShouldBindJSON(&request); err != nil {
		ctx.JSON(http.StatusBadRequest, ErrResponse{Error: err.Error()})
		return
	}

	res, err := c.Register(context.Background(), &pb.RegisterRequest{
		Email:    request.Email,
		Password: request.Password,
	})

	if err != nil {
		ctx.JSON(http.StatusBadGateway, &res)
		return
	}

	ctx.JSON(http.StatusCreated, &res)
}
