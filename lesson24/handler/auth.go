package handler

import (
	"expensity/pkg"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

func (h *Handler) Login(c *gin.Context) {
	username := c.PostForm("username")
	password := c.PostForm("password")

	// Dummy credential check
	// h.userService.GetUser(username, password)
	if (username == "employee" && password == "password") || (username == "senior" && password == "password") {
		tokenString, err := pkg.CreateToken(username)
		if err != nil {
			c.String(http.StatusInternalServerError, "Error creating token")
			return
		}

		fmt.Printf("Token created: %s\n", tokenString)
		c.JSON(http.StatusOK, gin.H{"token": tokenString})
	} else {
		c.String(http.StatusUnauthorized, "Invalid credentials")
	}

}
