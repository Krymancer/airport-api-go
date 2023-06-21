package cities

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/krymancer/airport-api-go/models"
)

// @Summary Delete a city
// @Description Delete a city from the database.
// @Tags Cities
// @Accept  json
// @Produce  json
// @Param   id path int true "City ID"
// @Success 200 {object} error "City successfully deleted"
// @Failure 404 {object} error "City not found"
// @Router /cities/{id} [delete]
func (h handler) DeleteCity(c *gin.Context) {
	id := c.Param("id")

	var city models.City

	if result := h.DB.First(&city, id); result.Error != nil {
		c.AbortWithError(http.StatusNotFound, result.Error)
		return
	}

	h.DB.Delete(&city)

	c.Status(http.StatusOK)
}
