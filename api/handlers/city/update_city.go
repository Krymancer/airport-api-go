package cities

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/krymancer/airport-api-go/pkg/common/models"
)

type UpdateCityRequestBody struct {
	Name  string `json:"name"`
	State string `json:"state"`
}

func (h handler) UpdateCity(c *gin.Context) {
	id := c.Param("id")
	body := UpdateCityRequestBody{}

	if err := c.BindJSON(&body); err != nil {
		c.AbortWithError(http.StatusBadRequest, err)
		return
	}

	var city models.City

	if result := h.DB.First(&city, id); result.Error != nil {
		c.AbortWithError(http.StatusNotFound, result.Error)
		return
	}

	city.Name = body.Name
	city.State = body.State

	h.DB.Save(&city)

	c.JSON(http.StatusOK, &city)
}
