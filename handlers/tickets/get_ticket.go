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
	TicketNumber   int    `json:"number" example:"1234" format:"integer" description:"The unique number of the ticket"`
	FlightNumber   int    `json:"flight_number" example:"5678" format:"integer" description:"The unique number of the flight"`
	Origin         string `json:"origin" example:"New York" format:"string" description:"The origin city of the flight"`
	Destination    string `json:"destination" example:"Los Angeles" format:"string" description:"The destination city of the flight"`
	PassengerCPF   string `json:"passenger_cpf" example:"123.456.789-00" format:"string" description:"The CPF number of the passenger"`
	PassengerName  string `json:"passenger_name" example:"John Doe" format:"string" description:"The name of the passenger"`
	SeatNumber     string `json:"seat_number" example:"A1" format:"string" description:"The number of the seat on the flight"`
	LuggageChecked bool   `json:"luggage_checked" example:"true" description:"Indicates if the passenger has checked luggage"`
}

// @Summary Retrieve a ticket voucher
// @Description Get the voucher details for a specific ticket by number
// @Tags Vouchers
// @Accept  json
// @Produce  json
// @Param number path int true "Ticket number"
// @Success 200 {object} Voucher "Successfully retrieved voucher"
// @Failure 404 {object} error "Not Found"
// @Router /voucher/{number} [get]
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
		TicketNumber:   ticket.Number,
		FlightNumber:   flight.Number,
		Origin:         origin.Name,
		Destination:    destination.Name,
		PassengerCPF:   passenger.CPF,
		PassengerName:  passenger.Name,
		SeatNumber:     ticket.SeatNumber,
		LuggageChecked: len(ticket.Lugagge) > 0,
	}

	c.JSON(http.StatusOK, &voucher)
}

// GetLuaggageLabel godoc
// @Summary Retrieve a ticket's luggage label
// @Description Get the luggage label for a specific ticket by number
// @Tags Lugagges
// @Accept  json
// @Produce  json
// @Param number path int true "Ticket number"
// @Success 200 {array} models.Lugagge "Successfully retrieved luggage label"
// @Failure 404 {object} error "Not Found"
// @Router /luggageLabel/{number} [get]
func (h handler) GetLuaggageLabel(c *gin.Context) {
	number := c.Param("number")

	var ticket models.Ticket

	if result := h.DB.Where("number = ?", number).First(&ticket); result.Error != nil {
		c.AbortWithError(http.StatusNotFound, result.Error)
		return
	}

	c.JSON(http.StatusOK, &ticket.Lugagge)
}
