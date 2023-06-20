package main

import (
	"github.com/gin-gonic/gin"
	"github.com/krymancer/airport-api-go/api/handlers/airports"
	"github.com/krymancer/airport-api-go/api/handlers/cities"
	"github.com/krymancer/airport-api-go/api/handlers/flights"
	"github.com/krymancer/airport-api-go/api/handlers/tickets"
	"github.com/krymancer/airport-api-go/pkg/common/db"
	"github.com/spf13/viper"
)

func main() {
	viper.SetConfigFile("./pkg/common/config/envs/.env")
	viper.ReadInConfig()

	port := viper.GetString("PORT")
	dbUrl := viper.GetString("DB_URL")

	r := gin.Default()
	h := db.InitDatabase(dbUrl)

	airports.RegisterRoutes(r, h)
	cities.RegisterRoutes(r, h)
	flights.RegisterRoutes(r, h)
	tickets.RegisterRoutes(r, h)

	r.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"port":  port,
			"dbUrl": dbUrl,
		})
	})

}
