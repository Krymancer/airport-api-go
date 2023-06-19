package airports

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

    routes := r.Group("/airports")
    routes.POST("/", h.AddAirport)
    routes.GET("/", h.GetAirports)
    routes.GET("/:id", h.GetAirport)
    routes.PUT("/:id", h.UpdateAirport)
    routes.DELETE("/:id", h.DeleteAirport)
}