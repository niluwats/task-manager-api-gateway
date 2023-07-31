package handlers

import (
	"context"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/niluwats/api-gateway/pkg/auth/pb"
)

// ViewUser godoc
// @Summary View user
// @Description Get user by ID
// @Tags users
// @Accept json
// @Produce json
// @Success 200 {object} pb.ViewUserResponse
// @Failure 400 {object} ErrResponse
// @Failure 401 {object} ErrResponse
// @Failure 404 {object} pb.ViewUserResponse
// @Failure 500 {object} pb.ViewUserResponse
// @Router /user/{id} [get]
func ViewUser(ctx *gin.Context, c pb.AuthServiceClient) {
	ID, _ := strconv.Atoi(ctx.Param("ID"))

	res, _ := c.ViewUser(context.Background(), &pb.ViewUserRequest{UserID: int64(ID)})

	ctx.JSON(int(res.Status), &res)
}
