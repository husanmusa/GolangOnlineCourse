package handler

import (
	_ "farmish/docs"
	"farmish/pkg/handler/middleware"
	"farmish/pkg/pubsub"
	"farmish/pkg/service"
	pb "farmish/proto"
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

type Handler struct {
	client   pb.AuthServiceClient
	services *service.Service
	ps       *pubsub.PubSub
}

// @title           Farmish API
// @version         1.0
// @description     This is a sample server celler server.
// @securityDefinitions.apikey	ApiKeyAuth
// @in							header
// @name						Authorization
// @host      localhost:8080
// @BasePath  /api/v1
func NewHandler(services *service.Service, ps *pubsub.PubSub, client pb.AuthServiceClient) *Handler {
	return &Handler{
		client,
		services,
		ps,
	}
}

func (h *Handler) InitRoutes() *gin.Engine {
	router := gin.Default()
	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	auth := router.Group("/auth")
	{
		auth.POST("/sign-up", h.SignUp)
		auth.POST("/sign-in", h.SignIn)
	}

	v1 := router.Group("/api/v1")

	v1.Use(middleware.AuthMiddleware)
	{
		farm := v1.Group("/farm")
		{
			farm.POST("/create", h.CreateFarm)
			farm.PUT("/update", h.UpdateFarm)
			farm.GET("/:id", h.GetFarm)
			farm.DELETE("/:id", h.DeleteFarm)
		}

		animal := v1.Group("/animal")
		{
			animal.POST("/create", h.CreateAnimal)
			animal.PUT("/update", h.UpdateAnimal)
			animal.DELETE("/delete/:id", h.DeleteAnimal)
			animal.GET("/:id", h.GetAnimalById)
			animal.POST("/feed", h.Feeding)
			animal.POST("/treat", h.Treatment)
			animal.POST("/healthy", h.ToggleHealthyAnimal)
			animal.POST("/hungry", h.ToggleHungryAnimal)
		}

		warehouse := v1.Group("/warehouse")
		{
			warehouse.POST("/create/stock", h.CreateStock)
			warehouse.POST("/supply/feed", h.SupplyFeedStock)
			warehouse.PUT("/supply/medicine", h.SupplyMedicineStock)
			warehouse.GET("/stock")
			warehouse.DELETE("/delete/stock/:id")
		}

		v1.Group("/dashboard").GET("", h.Dashboard)
	}
	return router
}
