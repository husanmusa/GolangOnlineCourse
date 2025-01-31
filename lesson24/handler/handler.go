package handler

import (
	"expensity/service"
	"fmt"
	"net/http"

	"expensity/pkg"

	"github.com/gin-gonic/gin"
)

type Handler struct {
	userService *service.UserService
}

func NewHandler(userService *service.UserService) *Handler {
	return &Handler{userService: userService}
}

func Run(h *Handler) *gin.Engine {
	r := gin.Default()

	//auth := r.Group("/auth")

	r.POST("/login", h.Login)

	//api := r.Group("/api")
	r.Use(authenticateMiddleware)

	r.POST("/user", h.CreateUser)

	return r
}

func authenticateMiddleware(c *gin.Context) {
	// Retrieve the token from the cookie
	tokenString := c.GetHeader("Authorization")

	// Verify the token
	token, err := pkg.VerifyToken(tokenString)
	if err != nil {
		fmt.Printf("Token verification failed: %v\\n", err)
		c.Redirect(http.StatusSeeOther, "/login")
		c.Abort()
		return
	}

	// Print information about the verified token
	fmt.Printf("Token verified successfully. Claims: %+v\\n", token.Claims)

	// Continue with the next middleware or route handler
	c.Next()
}
