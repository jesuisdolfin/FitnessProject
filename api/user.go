package api

import (
	"net/http"

	db "github.com/jesuisdolfin/FitnessProject/db/sqlc"
	"github.com/gin-gonic/gin"
)


type createUserRequest struct {
	Name string `json:"name" binding:"required"`
	Weight string `json:"weight" binding:"required"`
	Height string `json:"height" binding:"required"`
}
func (server *Server) createUser(ctx *gin.Context) {
	var req createUserRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	arg := db.CreateUserParams{
		Name: req.Name,
		Weight: req.Weight,
		Height: req.Height,
	}

	user, err := server.store.CreateUser(ctx, arg)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	ctx.JSON(http.StatusOK, user)
}