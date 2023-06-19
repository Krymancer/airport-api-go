package flights

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/krymancer/airport-api-go/pkg/common/models"
)

func (h handler) GetFlights(c *gin.Context) {

	var flight []models.Flight

	if result := h.DB.Find(&flight); result.Error != nil {
		c.AbortWithError(http.StatusNotFound, result.Error)
		return
	}

	c.JSON(http.StatusOK, &flight)
}
