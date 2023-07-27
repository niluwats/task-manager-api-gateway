package handlers

import (
	"context"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/niluwats/api-gateway/pkg/auth/pb"
)

func ViewUser(ctx *gin.Context, c pb.AuthServiceClient) {
	ID, _ := strconv.Atoi(ctx.Param("ID"))

	res, _ := c.ViewUser(context.Background(), &pb.ViewUserRequest{UserID: int64(ID)})

	ctx.JSON(int(res.Status), &res)
}
