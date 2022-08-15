package handler

import (
	"sign_in/service"

	"github.com/gin-gonic/gin"
)

type Handler struct {
	service *service.Service
}

func NewHandler(service *service.Service) *Handler {
	return &Handler{service: service}
}

func (h *Handler) SetupRouter() *gin.Engine {
	router := gin.New()

	auth := router.Group("/auth")
	{
		auth.POST("/sign-up", h.signUp)
		auth.POST("/sign-in", h.signIn)
	}
	api := router.Group("/api", h.userIdentity)
	{
		toys := api.Group("/toys")
		{
			toys.POST("/", h.createToysHandler)
			toys.GET("/", h.getAllToysHandler)
			toys.GET("/:id", h.getListByIdToysHandler)
			toys.DELETE("/:id", h.deleteToysHandler)
			toys.PUT("/:id", h.updateToysHandler)


			category := toys.Group("/:id/category")
			{
				category.POST("/", h.createCategoryHandler)
				category.GET("/", h.getAllCategoryHanlder)

			}

		}
		category := api.Group("/category")
		{

			category.GET("/:id", h.getCategoryByIdHandler)
			category.DELETE("/:id", h.deleteCategoryHandler)
			category.PUT("/:id", h.updateCategoryHandler)
		}

	}

	return router
}
