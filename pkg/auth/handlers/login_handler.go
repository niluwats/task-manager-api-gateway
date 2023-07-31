package handlers

import (
	"context"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/niluwats/api-gateway/pkg/auth/pb"
)

// LoginRequest containing string
// @swagger:parameters LoginRequest
type LoginRequest struct {
	Email    string `json:"email" binding:"required" example:"johnsmith@example.com" swaggertype:"integer"`
	Password string `json:"password" binding:"required" example:"John123!" swaggertype:"integer"`
}

// Login godoc
// @Summary Authenticate user
// @Description Generate JWT is user is authenticated
// @Tags users
// @Accept json
// @Produce json
// @Success 200 {object} pb.LoginResponse
// @Failure 400 {object} ErrResponse
// @Failure 404 {object} pb.LoginResponse
// @Failure 401 {object} pb.LoginResponse
// @Router /login [post]
func Login(ctx *gin.Context, c pb.AuthServiceClient) {
	request := LoginRequest{}

	if err := ctx.ShouldBindJSON(&request); err != nil {
		ctx.JSON(http.StatusBadRequest, ErrResponse{Error: err.Error()})
		return
	}

	res, _ := c.Login(context.Background(), &pb.LoginRequest{Email: request.Email, Password: request.Password})

	ctx.JSON(int(res.Status), &res)
}
