package flights

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/krymancer/airport-api-go/pkg/common/models"
)

// @Summary Get a flight
// @Description Retrieves a flight record from the database based on the provided ID.
// @Tags Flights
// @Accept  json
// @Produce  json
// @Param id path int true "Flight ID to get"
// @Success 200 {object} models.Flight "Successfully retrieved flight"
// @Failure 404 {object} error "Flight not found"
// @Router /flights/{id} [get]
func (h handler) GetFlight(c *gin.Context) {
	id := c.Param("id")

	var flight models.Flight

	if result := h.DB.First(&flight, id); result.Error != nil {
		c.AbortWithError(http.StatusNotFound, result.Error)
		return
	}

	c.JSON(http.StatusOK, &flight)
}
