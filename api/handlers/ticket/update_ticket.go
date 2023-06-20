package tickets

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/krymancer/airport-api-go/pkg/common/models"
)

type UpdateTicketBodyRequest struct {
	Number     int              `json:"number"`
	SeatNumber string           `json:"seat_number"`
	Lugagge    []models.Lugagge `json:"lugagge"`
	Flight     models.Flight    `json:"flight"`
	Passenger  models.Passenger `json:"passenger"`
}

func (h handler) UpdateTicket(c *gin.Context) {
	id := c.Param("id")
	body := UpdateTicketBodyRequest{}

	if err := c.BindJSON(&body); err != nil {
		c.AbortWithError(http.StatusBadRequest, err)
		return
	}

	var ticket models.Ticket
	var flight models.Flight
	var passenger models.Passenger

	if result := h.DB.First(&ticket, id); result.Error != nil {
		c.AbortWithError(http.StatusNotFound, result.Error)
		return
	}

	if result := h.DB.Where("number = ?", body.Flight.Number).First(&flight); result.Error != nil {
		c.AbortWithError(http.StatusNotFound, result.Error)
		return
	}

	if result := h.DB.Where("cpf = ?", body.Passenger.CPF).First(&passenger); result.Error != nil {
		c.AbortWithError(http.StatusNotFound, result.Error)
		return
	}

	ticket.Number = body.Number
	ticket.SeatNumber = body.SeatNumber
	ticket.Passenger = passenger
	ticket.Flight = flight
	ticket.Lugagge = body.Lugagge

	h.DB.Save(&ticket)

	c.JSON(http.StatusOK, &ticket)
}
