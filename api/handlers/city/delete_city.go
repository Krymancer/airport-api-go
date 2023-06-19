package cities

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/krymancer/airport-api-go/pkg/common/models"
)

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
