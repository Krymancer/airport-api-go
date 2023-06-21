package tickets

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/krymancer/airport-api-go/models"
)

// DeleteTicket godoc
// @Summary Delete a ticket
// @Description Delete a ticket by its ID
// @Tags Tickets
// @Accept  json
// @Produce  json
// @Param id path int true "Ticket ID"
// @Success 200 {string} string "Successfully deleted a ticket"
// @Failure 404 {object} error "Not Found"
// @Router /tickets/{id} [delete]
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
