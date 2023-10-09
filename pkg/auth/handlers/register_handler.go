package handlers

import (
	"context"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/niluwats/api-gateway/pkg/auth/pb"
)

// ErrResponse containing string
// @swagger:parameters ErrResponse
type ErrResponse struct {
	Error string `json:"error"`
}

// RegisterRequest represents body of Register request.
type RegisterRequest struct {
	FirstName string `json:"firstname" binding:"required" example:"john" swaggertype:"string"`
	LastName  string `json:"lastname" binding:"required" example:"smith" swaggertype:"string"`
	Email     string `json:"email" binding:"required,email" example:"johnsmith@example.com" swaggertype:"string"`
	Password  string `json:"password" binding:"required,gte=8,lte=32" example:"John123!" swaggertype:"string"`
}

// Register godoc
// @Summary Register
// @Description Create a new user
// @Tags users
// @Accept json
// @Produce json
// @Param requestBody body RegisterRequest true "RegisterRequest"
// @Success 201 {object} pb.RegisterResponse
// @Failure 400 {object} ErrResponse
// @Failure 500 {object} pb.RegisterResponse
// @Failure 409 {object} pb.RegisterResponse
// @Router /register [post]
func Register(ctx *gin.Context, c pb.AuthServiceClient) {
	request := RegisterRequest{}

	if err := ctx.ShouldBindJSON(&request); err != nil {
		ctx.JSON(http.StatusBadRequest, ErrResponse{Error: err.Error()})
		return
	}

	res, _ := c.Register(context.Background(), &pb.RegisterRequest{
		Firstname: request.FirstName,
		Lastname:  request.LastName,
		Email:     request.Email,
		Password:  request.Password,
	})

	ctx.JSON(int(res.Status), &res)
}
