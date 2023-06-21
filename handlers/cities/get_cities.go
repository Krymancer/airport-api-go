package cities

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/krymancer/airport-api-go/pkg/common/models"
)

// @Summary Retrieve all cities
// @Description Get the list of all cities from the database.
// @Tags Cities
// @Accept  json
// @Produce  json
// @Success 200 {array} models.City "Successfully retrieved all cities"
// @Failure 404 {object} error "No cities found"
// @Router /cities [get]
func (h handler) GetCities(c *gin.Context) {

	var cities []models.City

	if result := h.DB.Find(&cities); result.Error != nil {
		c.AbortWithError(http.StatusNotFound, result.Error)
		return
	}

	c.JSON(http.StatusOK, &cities)
}
