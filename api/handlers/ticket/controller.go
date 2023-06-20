package tickets

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

	routes := r.Group("/tickets")
	routes.POST("/", h.AddTicket)
	routes.GET("/", h.GetTickets)
	routes.GET("/flight/:flight", h.GetTicketsByFlight)
	routes.GET("/cpf/:cpf", h.GetTicketsByCPF)
	routes.GET("/:id", h.GetTicket)
	routes.PUT("/:id", h.UpdateTicket)
	routes.DELETE("/:id", h.DeleteTicket)
}
