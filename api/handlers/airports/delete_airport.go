package airports

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/krymancer/airport-api-go/pkg/common/models"
)

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
