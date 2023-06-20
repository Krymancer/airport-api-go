package tickets

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/krymancer/airport-api-go/pkg/common/models"
)

func (h handler) DeleteTicket(c *gin.Context) {
	id := c.Param("id")

	var ticket models.Ticket

	if result := h.DB.First(&ticket, id); result.Error != nil {
		c.AbortWithError(http.StatusNotFound, result.Error)
		return
	}

	h.DB.Delete(&ticket)

	c.Status(http.StatusOK)
}
