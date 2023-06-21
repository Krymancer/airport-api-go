package flights

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/krymancer/airport-api-go/models"
)

// @Summary Delete a flight
// @Description Deletes a flight record from the database based on the provided ID.
// @Tags Flights
// @Accept  json
// @Produce  json
// @Param id path int true "Flight ID to delete"
// @Success 200 {object} models.Flight "Successfully deleted flight"
// @Router /flights/{id} [delete]
func (h handler) DeleteFlight(c *gin.Context) {
	id := c.Param("id")

	var flight models.Flight

	if result := h.DB.First(&flight, id); result.Error != nil {
		c.AbortWithError(http.StatusNotFound, result.Error)
		return
	}

	h.DB.Delete(&flight)

	c.Status(http.StatusOK)
}
