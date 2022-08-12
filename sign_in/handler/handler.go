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
	toys := router.Group("/")
	{
		toys.POST("toys", h.createToysHandler)
		toys.GET("toys", h.getAllToysHandler)
		toys.GET("toys/:id", h.getListByIdToysHandler)
		toys.DELETE("toys/:id", h.deleteToysHandler)
		toys.PUT("toys/:id", h.updateToysHandler)
	}
	category := router.Group("/")
	{
		category.POST("category", h.createCategoryHandler)
		category.GET("category", h.getAllCategoryHanlder)
		category.GET("category/:id", h.getListByIdToysHandler)
		category.DELETE("category/:id", h.deleteCategoryHandler)
		category.PUT("category/:id", h.updateCategoryHandler)
	}


	return router
}

