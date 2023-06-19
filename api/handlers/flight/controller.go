package flights

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

	routes := r.Group("/flights")
	routes.POST("/", h.AddFlight)
	routes.GET("/", h.GetFlights)
	routes.GET("/:id", h.GetFlight)
	routes.PUT("/:id", h.UpdateFlight)
	routes.DELETE("/:id", h.DeleteFlight)
}
