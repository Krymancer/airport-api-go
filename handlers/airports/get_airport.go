package airports

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/krymancer/airport-api-go/models"
)

// @Summary Get an airport
// @Description Get details of an airport using its ID
// @Tags Airports
// @Accept  json
// @Produce  json
// @Param id path string true "Airport ID"
// @Success 200 {object} models.Airport "Successfully retrieved airport"
// @Failure 404 {object} error "Not Found"
// @Router /airports/{id} [get]
func (h handler) GetAirport(c *gin.Context) {
	id := c.Param("id")

	var airport models.Airport

	if result := h.DB.First(&airport, id); result.Error != nil {
		c.AbortWithError(http.StatusNotFound, result.Error)
		return
	}

	c.JSON(http.StatusOK, &airport)
}
