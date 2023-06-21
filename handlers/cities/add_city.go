package cities

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/krymancer/airport-api-go/pkg/common/models"
)

type AddCityBodyRequest struct {
	Name  string `json:"name" example:"New York" format:"string" description:"Name of the city"`
	State string `json:"state" example:"NY" format:"string" description:"Abbreviation of the state"`
}

// @Summary Add a new city
// @Description Add a new city to the database.
// @Tags Cities
// @Accept  json
// @Produce  json
// @Param   body body AddCityBodyRequest true "City details"
// @Success 201 {object} models.City "Created city object"
// @Failure 400 {object} error "Bad request, JSON decoding failed"
// @Failure 404 {object} error "Internal server error when saving the city into the database"
// @Router /cities [post]
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
