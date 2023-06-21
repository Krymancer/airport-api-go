package flights

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/krymancer/airport-api-go/pkg/common/models"
)

// @Summary Get all flights
// @Description Retrieves all flight records from the database.
// @Tags Flights
// @Accept  json
// @Produce  json
// @Success 200 {array} models.Flight "Successfully retrieved all flights"
// @Failure 404 {object} error "No flights found"
// @Router /flights [get]
func (h handler) GetFlights(c *gin.Context) {

	var flight []models.Flight

	if result := h.DB.Find(&flight); result.Error != nil {
		c.AbortWithError(http.StatusNotFound, result.Error)
		return
	}

	c.JSON(http.StatusOK, &flight)
}
