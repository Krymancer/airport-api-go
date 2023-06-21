package cities

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/krymancer/airport-api-go/pkg/common/models"
)

type UpdateCityRequestBody struct {
	Name  string `json:"name" example:"New York" format:"string" description:"Name of the city"`
	State string `json:"state" example:"NY" format:"string" description:"Abbreviation of the state"`
}

// @Summary Update a city's details
// @Description Update the details of a specific city in the database by its ID.
// @Tags Cities
// @Accept  json
// @Produce  json
// @Param id path int true "City ID"
// @Param name body string false "Name of the city"
// @Param state body string false "State of the city"
// @Success 200 {object} models.City "Successfully updated the city"
// @Failure 400 {object} error "Invalid request body"
// @Failure 404 {object} error "City not found"
// @Router /cities/{id} [put]
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
