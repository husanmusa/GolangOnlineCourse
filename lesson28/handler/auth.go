package handler

import (
	"fmt"
	"lesson28/pkg"
	"net/http"

	"github.com/gin-gonic/gin"
)

// Login ...
//
//	@Summary		Login a user
//	@Description	Login a user
//	@Tags			users
//	@Produce		json
//	@Accept			application/x-www-form-urlencoded
//	@Param			username	formData		string	true "username"
//	@Param			password	formData		string	true "password"
//	@Success		201		{object}	interface{}
//	@Failure		404		{object}	error
//	@Failure		422		{object}	error
//	@Failure		500		{object}	error
//	@Router			/login [post]
func (h *Handler) Login(c *gin.Context) {
	username := c.PostForm("username")
	password := c.PostForm("password")

	// Dummy credential check
	// h.userService.GetUser(username, password)
	if (username == "employee" && password == "password") || (username == "senior" && password == "password") {
		tokenString, err := pkg.CreateToken(username)
		if err != nil {
			fmt.Println(err)
			c.String(http.StatusInternalServerError, "Error creating token")
			return
		}

		fmt.Printf("Token created: %s\n", tokenString)
		c.JSON(http.StatusOK, gin.H{"token": tokenString})
	} else {
		c.String(http.StatusUnauthorized, "Invalid credentials")
	}

}
