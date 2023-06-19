package flights

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/krymancer/airport-api-go/pkg/common/models"
)

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
