package cities

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/krymancer/airport-api-go/pkg/common/models"
)

type AddCityBodyRequest struct {
	Name  string `json:"name"`
	State string `json:"state"`
}

func (h handler) AddCity(c *gin.Context) {
	body := AddCityBodyRequest{}

	if err := c.BindJSON(&body); err != nil {
		c.AbortWithError(http.StatusBadRequest, err)
		return
	}

	var city models.City

	city.Name = body.Name
	city.State = body.State

	if result := h.DB.Create(&city); result.Error != nil {
		c.AbortWithError(http.StatusNotFound, result.Error)
		return
	}

	c.JSON(http.StatusCreated, &city)
}
