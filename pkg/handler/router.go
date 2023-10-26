package handler

import (
	"github.com/labstack/echo/v4"
)

func (h *Handler) GetRouter() *echo.Echo {
	router := echo.New()

	// Маршруты для изображений APOD
	router.POST("/api/data", h.InsertData)
	router.GET("/api/data", h.GetPeopleByAge)
	router.PUT("/api/data/:id", h.UpdateInfo)
	router.DELETE("/api/data/:id", h.DeleteImageByID)

	return router
}
