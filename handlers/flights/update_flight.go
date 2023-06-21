package flights

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/krymancer/airport-api-go/pkg/common/models"
)

type UpdateFlightBodyRequest struct {
	OriginID      uint                 `json:"origin_id" binding:"required" example:"1" format:"integer" description:"ID of the origin airport"`
	DestinationID uint                 `json:"destination_id" binding:"required" example:"2" format:"integer" description:"ID of the destination airport"`
	FlightClasses []models.FlightClass `json:"flight_classes" binding:"required" description:"Array of Flight Classes"`
	Number        string               `json:"number" binding:"required" example:"FA123" format:"string" description:"The flight number"`
}

// @Summary Update a flight
// @Description Updates a flight record with the provided data.
// @Tags Flights
// @Accept  json
// @Produce  json
// @Param   id     path    int     true        "Flight ID"
// @Param   body   body    UpdateFlightBodyRequest  true "Body object"
// @Success 200 {object} models.Flight "Successfully updated flight"
// @Failure 400 {object} error "Bad Request"
// @Failure 404 {object} error "Flight Not Found"
// @Failure 409 {object} error "Conflict with existing flight number"
// @Router /flights/{id} [put]
func (h handler) UpdateFlight(c *gin.Context) {
	id := c.Param("id")
	body := UpdateFlightBodyRequest{}

	if err := c.BindJSON(&body); err != nil {
		c.AbortWithError(http.StatusBadRequest, err)
		return
	}

	var flight models.Flight
	var origin models.Airport
	var destination models.Airport

	if result := h.DB.Where("number = ?", body.Number).First(&flight); result.Error == nil {
		c.AbortWithStatus(http.StatusConflict)
		return
	}

	if result := h.DB.First(&origin, body.OriginID); result.Error != nil {
		jsonData := []byte(`{"error": "Origin airport not found"}`)
		c.AbortWithStatusJSON(http.StatusNotFound, jsonData)
		return
	}

	flight.OriginID = origin.ID

	if result := h.DB.First(&destination, body.DestinationID); result.Error != nil {
		jsonData := []byte(`{"error": "Destination airport not found"}`)
		c.AbortWithStatusJSON(http.StatusNotFound, jsonData)
		return
	}

	flight.DestinationID = destination.ID

	number, err := strconv.Atoi(body.Number)

	if err != nil {
		c.AbortWithError(http.StatusBadRequest, err)
		return
	}

	flight.Number = number

	if result := h.DB.First(&flight, id); result.Error != nil {
		c.AbortWithError(http.StatusNotFound, result.Error)
		return
	}

	flight.Number = number
	flight.FlightClasses = body.FlightClasses

	h.DB.Save(&flight)

	c.JSON(http.StatusOK, &flight)
}

type UpdateFlightClassBodyRequest struct {
	Price float64 `json:"price" example:"100.50" format:"float" description:"Price per seat for the flight class" binding:"required"`
}

// @Summary Update a flight class price
// @Description Updates the price of a flight class with the provided data.
// @Tags Flights
// @Accept  json
// @Produce  json
// @Param   number     path    string     true        "Flight Number"
// @Param   class      path    string     true        "Flight Class"
// @Param   body       body    UpdateFlightClassBodyRequest  true "Body object"
// @Success 200 {object} models.Flight "Successfully updated flight class price"
// @Failure 400 {object} error "Bad Request"
// @Failure 404 {object} error "Flight or Flight Class Not Found"
// @Router /flights/{number}/{class} [put]
func (h handler) UpdateClassPrice(c *gin.Context) {
	number := c.Param("number")
	class := c.Param("class")

	body := UpdateFlightClassBodyRequest{}

	if err := c.BindJSON(&body); err != nil {
		c.AbortWithError(http.StatusBadRequest, err)
		return
	}

	var flight models.Flight
	var flightClass models.FlightClass

	if result := h.DB.First(&flight, number); result.Error != nil {
		c.AbortWithError(http.StatusNotFound, result.Error)
		return
	}

	if result := h.DB.First(&flightClass, class); result.Error != nil {
		c.AbortWithError(http.StatusNotFound, result.Error)
		return
	}

	flightClass.PricePerSeat = body.Price

	h.DB.Save(&flight)

	c.JSON(http.StatusOK, &flight)
}
