package cities

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/krymancer/airport-api-go/pkg/common/models"
)

// @Summary Retrieve a city by ID
// @Description Get a specific city from the database by its ID.
// @Tags Cities
// @Accept  json
// @Produce  json
// @Param id path int true "City ID"
// @Success 200 {object} models.City "Successfully retrieved the city"
// @Failure 404 {object} error "City not found"
// @Router /cities/{id} [get]
func (h handler) GetCity(c *gin.Context) {
	id := c.Param("id")

	var city models.City

	if result := h.DB.First(&city, id); result.Error != nil {
		c.AbortWithError(http.StatusNotFound, result.Error)
		return
	}

	c.JSON(http.StatusOK, &city)
}
