package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/niluwats/api-gateway/pkg/task/pb"
	"google.golang.org/grpc/status"
)

// ErrResponse containing string
// @swagger:parameters ErrResponse
type ErrResponse struct {
	Error string `json:"error"`
}

// ProjectReq represents body of a new project request and update project request.
type ProjectReq struct {
	Name        string  `json:"name" binding:"required" example:"alpha" swaggertype:"string"`
	Description string  `json:"description" binding:"required" example:"this is the phase 1" swaggertype:"string"`
	Creator     int32   `json:"creator" binding:"required" example:"1" swaggertype:"integer"`
	Assignees   []int32 `json:"assignees" binding:"required" example:"[1,2,3]" swaggertype:"array,integer"`
}

// CreateProject godoc
// @Summary Create Project
// @Description Create a new project
// @Tags projects
// @Accept json
// @Produce json
// @Param requestBody body ProjectReq true "ProjectReq"
// @Success 201 {object} pb.ProjectResponse
// @Failure 400 {object} ErrResponse
// @Failure 500 {object} pb.ProjectResponse
// @Failure 409 {object} pb.ProjectResponse
// @Router /project [post]
func CreateProject(ctx *gin.Context, c pb.TaskServiceClient) {
	request := ProjectReq{}

	if err := ctx.ShouldBindJSON(&request); err != nil {
		ctx.JSON(http.StatusBadRequest, ErrResponse{Error: err.Error()})
		return
	}

	res, err := c.CreateProject(ctx, &pb.NewProjectRequest{
		Name:        request.Name,
		Description: request.Description,
		Creator:     request.Creator,
		Assignees:   request.Assignees,
	})

	if err != nil {
		statusErr, ok := status.FromError(err)
		if ok {
			ctx.JSON(int(statusErr.Code()), statusErr.Message())
		} else {
			ctx.JSON(http.StatusInternalServerError, err.Error())
		}
		return
	}

	ctx.JSON(int(res.CommonResponse.Status), &res)
}

// UpdateProject godoc
// @Summary Update Project
// @Description Update a project
// @Tags projects
// @Accept json
// @Produce json
// @Param requestBody body ProjectReq true "ProjectReq"
// @Success 201 {object} pb.ProjectResponse
// @Failure 400 {object} ErrResponse
// @Failure 500 {object} pb.ProjectResponse
// @Failure 409 {object} pb.ProjectResponse
// @Router /project/{ID} [put]
func UpdateProject(ctx *gin.Context, c pb.TaskServiceClient) {
	Id := ctx.Param("ID")
	request := ProjectReq{}

	if err := ctx.ShouldBindJSON(&request); err != nil {
		ctx.JSON(http.StatusBadRequest, ErrResponse{Error: err.Error()})
		return
	}

	res, _ := c.UpdateProject(ctx, &pb.UpdaterojectRequest{
		Id:          Id,
		Name:        request.Name,
		Description: request.Description,
		Creator:     request.Creator,
		Assignees:   request.Assignees,
	})

	ctx.JSON(int(res.CommonResponse.Status), &res)
}

// RemoveProject godoc
// @Summary Remove Project
// @Description remove a project
// @Tags projects
// @Accept json
// @Produce json
// @Success 201 {object} pb.RemovedProjectResponse
// @Failure 400 {object} ErrResponse
// @Failure 500 {object} pb.RemovedProjectResponse
// @Failure 409 {object} pb.RemovedProjectResponse
// @Router /project/{ID} [delete]
func RemoveProject(ctx *gin.Context, c pb.TaskServiceClient) {
	Id := ctx.Param("ID")

	res, _ := c.Remove(ctx, &pb.ViewProjectRequest{Id: Id})
	ctx.JSON(int(res.CommonResponse.Status), &res)
}

// ViewProject godoc
// @Summary View Project
// @Description view a project by ID
// @Tags projects
// @Accept json
// @Produce json
// @Success 201 {object} pb.ProjectResponse
// @Failure 400 {object} ErrResponse
// @Failure 500 {object} pb.ProjectResponse
// @Failure 409 {object} pb.ProjectResponse
// @Router /project/{ID} [get]
func ViewProject(ctx *gin.Context, c pb.TaskServiceClient) {
	Id := ctx.Param("ID")

	res, _ := c.ViewProject(ctx, &pb.ViewProjectRequest{Id: Id})
	ctx.JSON(int(res.CommonResponse.Status), &res)
}

// ViewAllProjects godoc
// @Summary View Projects
// @Description view all project
// @Tags projects
// @Accept json
// @Produce json
// @Success 201 {object} pb.ProjectsResponse
// @Failure 400 {object} ErrResponse
// @Failure 500 {object} pb.ProjectsResponse
// @Failure 409 {object} pb.ProjectsResponse
// @Router /project [get]
func ViewAllProjects(ctx *gin.Context, c pb.TaskServiceClient) {
	res, _ := c.ViewAllProjects(ctx, &pb.EmptyParams{})
	ctx.JSON(int(res.CommonResponse.Status), &res)
}
