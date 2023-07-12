package httpserver

import (
	"context"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/niluwats/api-gateway/pkg/auth/dto"
	"github.com/niluwats/api-gateway/pkg/auth/pb"
)

func Register(ctx *gin.Context, c pb.AuthServiceClient) {
	request := dto.RegisterRequest{}

	if err := ctx.BindJSON(&request); err != nil {
		ctx.JSON(http.StatusBadRequest, dto.Response{Status: false, Message: err.Error()})
		return
	}

	res, err := c.Register(context.Background(), &pb.RegisterRequest{
		Email:    request.Email,
		Password: request.Password,
	})

	if err != nil {
		ctx.JSON(http.StatusBadGateway, dto.Response{Status: false, Message: err.Error()})
		return
	}

	ctx.JSON(http.StatusCreated, &res)
}
