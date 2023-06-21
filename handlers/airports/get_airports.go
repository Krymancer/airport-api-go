package airports

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/krymancer/airport-api-go/pkg/common/models"
)

// @Summary Get all airports
// @Description Get a list of all airports
// @Tags Airports
// @Accept  json
// @Produce  json
// @Success 200 {array} models.Airport "Successfully retrieved all airports"
// @Failure 404 {object} error "Not Found"
// @Router /airports [get]
func (h handler) GetAirports(c *gin.Context) {

	var airports []models.Airport

	if result := h.DB.Find(&airports); result.Error != nil {
		c.AbortWithError(http.StatusNotFound, result.Error)
		return
	}

	c.JSON(http.StatusOK, &airports)
}
