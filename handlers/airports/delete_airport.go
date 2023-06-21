package airports

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/krymancer/airport-api-go/pkg/common/models"
)

// @Summary Delete an airport
// @Description Delete an airport from the system using its ID
// @Tags Airports
// @Accept  json
// @Produce  json
// @Param id path string true "Airport ID"
// @Success 200 {object} int "OK"
// @Failure 404 {object} error "Not Found"
// @Router /airports/{id} [delete]
func (h handler) DeleteAirport(c *gin.Context) {
	id := c.Param("id")

	var airport models.Airport

	if result := h.DB.First(&airport, id); result.Error != nil {
		c.AbortWithError(http.StatusNotFound, result.Error)
		return
	}

	h.DB.Delete(&airport)

	c.Status(http.StatusOK)
}
