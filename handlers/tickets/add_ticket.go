package tickets

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/krymancer/airport-api-go/pkg/common/models"
)

type AddTicketBodyRequest struct {
	Number      int              `json:"number" example:"1" format:"integer" description:"The unique number of the ticket"`
	SeatNumber  string           `json:"seat_number" example:"A1" format:"string" description:"The number of the seat on the flight"`
	Lugagge     []models.Lugagge `json:"lugagge" description:"Array of luggage items associated with the ticket"`
	FlightID    uint             `json:"flight" example:"123" format:"integer" description:"The ID of the flight associated with the ticket"`
	PassengerID uint             `json:"passenger" example:"456" format:"integer" description:"The ID of the passenger to whom the ticket is issued"`
}

// AddTicket godoc
// @Summary Add a new ticket
// @Description Add a new ticket with the provided details
// @Tags Tickets
// @Accept  json
// @Produce  json
// @Param AddTicketBodyRequest body AddTicketBodyRequest true "Add a new ticket"
// @Success 201 {object} models.Ticket "Successfully created a new ticket"
// @Failure 400 {object} error "Bad Request"
// @Failure 404 {object} error "Not Found"
// @Failure 409 {object} error "Conflict"
// @Router /tickets [post]
func (h handler) AddTicket(c *gin.Context) {
	body := AddTicketBodyRequest{}

	if err := c.BindJSON(&body); err != nil {
		c.AbortWithError(http.StatusBadRequest, err)
		return
	}

	var ticket models.Ticket
	var flight models.Flight
	var passenger models.Passenger

	if result := h.DB.Where("number = ?", body.Number).First(&ticket); result.Error != nil {
		c.AbortWithError(http.StatusConflict, result.Error)
		return
	}

	if result := h.DB.First(&flight, body.FlightID); result.Error != nil {
		c.AbortWithError(http.StatusNotFound, result.Error)
		return
	}

	if result := h.DB.First(&passenger, body.PassengerID); result.Error != nil {
		c.AbortWithError(http.StatusNotFound, result.Error)
		return
	}

	ticket.Number = body.Number
	ticket.SeatNumber = body.SeatNumber
	ticket.PassengerID = passenger.ID
	ticket.FlightID = flight.ID
	ticket.Lugagge = body.Lugagge

	if result := h.DB.Create(&ticket); result.Error != nil {
		c.AbortWithError(http.StatusBadRequest, result.Error)
		return
	}

	c.JSON(http.StatusCreated, &ticket)
}

type AddLugggageBodyRequest struct {
	Number int `json:"number"`
}

func (h handler) AddLuggage(c *gin.Context) {
	number := c.Param("number")

	body := AddLugggageBodyRequest{}

	if err := c.BindJSON(&body); err != nil {
		c.AbortWithError(http.StatusBadRequest, err)
		return
	}

	var ticket models.Ticket

	if result := h.DB.Where("number = ?", number).First(&ticket); result.Error != nil {
		c.AbortWithError(http.StatusNotFound, result.Error)
		return
	}

	var luggage models.Lugagge

	if result := h.DB.Where("number = ?", body.Number).First(&luggage); result.Error != nil {
		c.AbortWithError(http.StatusNotFound, result.Error)
		return
	}

	ticket.Lugagge = append(ticket.Lugagge, luggage)

	if result := h.DB.Save(&ticket); result.Error != nil {
		c.AbortWithError(http.StatusBadRequest, result.Error)
		return
	}

	c.JSON(http.StatusCreated, &luggage.Number)
}
