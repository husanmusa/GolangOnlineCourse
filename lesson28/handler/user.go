package handler

import (
	"net/http"

	"lesson28/models"

	"github.com/gin-gonic/gin"
)

// CreateUser creates a user in the user service
//
//	@Summary		Creates a user
//	@Description	Creates a user
//
// @Security 		ApiKeyAuth
//
//	@Tags			users
//	@Accept			json
//	@Produce		json
//	@Param			user	body		models.User	true "User to create"
//	@Success		201		{object}	interface{}
//	@Failure		404		{object}	error
//	@Failure		422		{object}	error
//	@Failure		500		{object}	error
//	@Router			/user [post]
func (h *Handler) CreateUser(ctx *gin.Context) {
	user := models.User{}
	if err := ctx.BindJSON(&user); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := h.userService.CreateUser(user); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusCreated, gin.H{"message": "User created"})
}
