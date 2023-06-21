package tickets

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/krymancer/airport-api-go/models"
)

// @Summary Get all tickets
// @Description get all tickets
// @Tags Tickets
// @Accept  json
// @Produce  json
// @Success 200 {array} models.Ticket
// @Failure 404 {object} error "No tickets found"
// @Router /tickets [get]
func (h handler) GetTickets(c *gin.Context) {

	var tickets []models.Ticket

	if result := h.DB.Find(&tickets); result.Error != nil {
		c.AbortWithError(http.StatusNotFound, result.Error)
		return
	}

	c.JSON(http.StatusOK, &tickets)
}

func (h handler) GetTicketsByFlight(c *gin.Context) {
	flightNumberParam := c.Param("flight")

	flightNumber, err := strconv.Atoi(flightNumberParam)

	if err != nil {
		c.AbortWithError(http.StatusBadRequest, err)
		return
	}

	var tickets []models.Ticket

	if result := h.DB.Where("flight_number = ?", flightNumber).Find(&tickets); result.Error != nil {
		c.AbortWithError(http.StatusNotFound, result.Error)
		return
	}

	c.JSON(http.StatusOK, &tickets)
}

// GetTicketsByCPF godoc
// @Summary Get tickets by CPF
// @Description get tickets by passenger's CPF
// @Tags Tickets
// @Accept  json
// @Produce  json
// @Param cpf path string true "CPF"
// @Success 200 {array} models.Ticket
// @Failure 404 {object} error "No tickets found"
// @Router /tickets/cpf/{cpf} [get]
func (h handler) GetTicketsByCPF(c *gin.Context) {
	cpf := c.Param("cpf")

	var tickets []models.Ticket

	if result := h.DB.Preload("Passenger").Where("passenger.cpf = ?", cpf).Find(&tickets); result.Error != nil {
		c.AbortWithError(http.StatusNotFound, result.Error)
		return
	}

	c.JSON(http.StatusOK, &tickets)
}
