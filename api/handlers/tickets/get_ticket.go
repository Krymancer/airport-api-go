package tickets

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/krymancer/airport-api-go/pkg/common/models"
)

func (h handler) GetTicket(c *gin.Context) {
	id := c.Param("id")

	var ticket models.Ticket

	if result := h.DB.First(&ticket, id); result.Error != nil {
		c.AbortWithError(http.StatusNotFound, result.Error)
		return
	}

	c.JSON(http.StatusOK, &ticket)
}

type Voucher struct {
	TicketNumber  int    `json:"number"`
	FlightNumber  int    `json:"flight_number"`
	Origin        string `json:"origin"`
	Destination   string `json:"destination"`
	PassengerCPF  string `json:"passenger_cpf"`
	PassengerName string `json:"passenger_name"`
	SeatNumber    string `json:"seat_number"`
	LuggageCheked bool   `json:"luggage_checked"`
}

func (h handler) GetVoucher(c *gin.Context) {
	number := c.Param("number")

	var ticket models.Ticket
	var flight models.Flight
	var origin models.City
	var destination models.City
	var passenger models.Passenger

	if result := h.DB.Where("number = ?", number).First(&ticket); result.Error != nil {
		c.AbortWithError(http.StatusNotFound, result.Error)
		return
	}

	if result := h.DB.First(&flight, ticket.FlightID); result.Error != nil {
		c.AbortWithError(http.StatusNotFound, result.Error)
		return
	}

	if result := h.DB.First(&origin, flight.OriginID); result.Error != nil {
		c.AbortWithError(http.StatusNotFound, result.Error)
		return
	}

	if result := h.DB.First(&destination, flight.DestinationID); result.Error != nil {
		c.AbortWithError(http.StatusNotFound, result.Error)
		return
	}

	if result := h.DB.First(&passenger, ticket.PassengerID); result.Error != nil {
		c.AbortWithError(http.StatusNotFound, result.Error)
		return
	}

	voucher := Voucher{
		TicketNumber:  ticket.Number,
		FlightNumber:  flight.Number,
		Origin:        origin.Name,
		Destination:   destination.Name,
		PassengerCPF:  passenger.CPF,
		PassengerName: passenger.Name,
		SeatNumber:    ticket.SeatNumber,
		LuggageCheked: len(ticket.Lugagge) > 0,
	}

	c.JSON(http.StatusOK, &voucher)
}

func (h handler) GetLuaggageLabel(c *gin.Context) {
	number := c.Param("number")

	var ticket models.Ticket

	if result := h.DB.Where("number = ?", number).First(&ticket); result.Error != nil {
		c.AbortWithError(http.StatusNotFound, result.Error)
		return
	}

	c.JSON(http.StatusOK, &ticket.Lugagge)
}
