package cities

import (
	"github.com/gin-gonic/gin"

	"gorm.io/gorm"
)

type handler struct {
	DB *gorm.DB
}

func RegisterRoutes(r *gin.Engine, db *gorm.DB) {
	h := &handler{
		DB: db,
	}

	routes := r.Group("/cities")
	routes.POST("/", h.AddCity)
	routes.GET("/", h.GetCities)
	routes.GET("/:id", h.GetCity)
	routes.PUT("/:id", h.UpdateCity)
	routes.DELETE("/:id", h.DeleteCity)
}
