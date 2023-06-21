package tickets

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/krymancer/airport-api-go/models"
)

type UpdateTicketBodyRequest struct {
	Number      int              `json:"number" example:"1" format:"integer" description:"The unique number of the ticket"`
	SeatNumber  string           `json:"seat_number" example:"A1" format:"string" description:"The number of the seat on the flight"`
	Lugagge     []models.Lugagge `json:"lugagge" description:"Array of luggage items associated with the ticket"`
	FlightID    uint             `json:"flight" example:"123" format:"integer" description:"The ID of the flight associated with the ticket"`
	PassengerID uint             `json:"passenger" example:"456" format:"integer" description:"The ID of the passenger to whom the ticket is issued"`
}

// UpdateTicket godoc
// @Summary Update a ticket
// @Description update ticket by ID
// @Tags Tickets
// @Accept  json
// @Produce  json
// @Param id path int true "Ticket ID"
// @Param body body UpdateTicketBodyRequest true "Ticket details to be updated"
// @Success 200 {object} models.Ticket
// @Failure 400 {object} error "Bad Request"
// @Failure 404 {object} error "Not Found"
// @Router /tickets/{id} [put]
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

	if result := h.DB.Where("number = ?", body.FlightID).First(&flight); result.Error != nil {
		c.AbortWithError(http.StatusNotFound, result.Error)
		return
	}

	if result := h.DB.Where("cpf = ?", body.PassengerID).First(&passenger); result.Error != nil {
		c.AbortWithError(http.StatusNotFound, result.Error)
		return
	}

	ticket.Number = body.Number
	ticket.SeatNumber = body.SeatNumber
	ticket.PassengerID = passenger.ID
	ticket.FlightID = flight.ID
	ticket.Lugagge = body.Lugagge

	h.DB.Save(&ticket)

	c.JSON(http.StatusOK, &ticket)
}
