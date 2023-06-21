package main

import (
	"github.com/gin-gonic/gin"
	"github.com/krymancer/airport-api-go/handlers/airports"
	"github.com/krymancer/airport-api-go/handlers/cities"
	"github.com/krymancer/airport-api-go/handlers/flights"
	"github.com/krymancer/airport-api-go/handlers/tickets"
	"github.com/krymancer/airport-api-go/internal/database"
	"github.com/spf13/viper"

	docs "github.com/krymancer/airport-api-go/docs"
	swaggerfiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

// @title           Airport API
// @version         1.0
// @description     This implements the Airport API Reference.
// @termsOfService  http://swagger.io/terms/
// @host      localhost:3000
// @BasePath  /
func main() {
	viper.SetConfigFile("./config/envs/.env")
	viper.ReadInConfig()

	port := viper.GetString("PORT")
	databaseUrl := viper.GetString("DB_URL")

	docs.SwaggerInfo.BasePath = "/"

	r := gin.Default()
	h := database.InitDatabase(databaseUrl)

	airports.RegisterRoutes(r, h)
	cities.RegisterRoutes(r, h)
	flights.RegisterRoutes(r, h)
	tickets.RegisterRoutes(r, h)

	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))

	r.Run(port)
}
