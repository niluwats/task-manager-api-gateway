package handlers

import (
	"context"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/niluwats/api-gateway/pkg/auth/pb"
)

func ViewUser(ctx *gin.Context, c pb.AuthServiceClient) {
	ID, _ := strconv.Atoi(ctx.Param("ID"))

	res, err := c.ViewUser(context.Background(), &pb.ViewUserRequest{UserID: int64(ID)})
	if err != nil {
		ctx.JSON(http.StatusBadGateway, &res)
		return
	}

	ctx.JSON(http.StatusOK, &res)
}
