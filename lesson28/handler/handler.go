package handler

import (
	"fmt"
	"lesson28/service"
	"net/http"

	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	"lesson28/pkg"

	"github.com/gin-gonic/gin"

	_ "lesson28/docs"
)

type Handler struct {
	userService *service.UserService
}

func NewHandler(userService *service.UserService) *Handler {
	return &Handler{userService: userService}
}

// Run ...
// @title			LIST API
// @securityDefinitions.apikey ApiKeyAuth
// @in header
// @name Authorization
// @version		1.0
// @description	Testing Swagger APIs.
// @contact.name	API Support
// @contact.url	http://www.swagger.io/support
// @contact.email	support@swagger.io
// @host			localhost:8080
func Run(h *Handler) *gin.Engine {
	r := gin.Default()

	url := ginSwagger.URL("http://localhost:8080/swagger/doc.json")
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler, url))

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
