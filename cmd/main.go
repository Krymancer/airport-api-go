package main

import (
	"github.com/gin-gonic/gin"
	"github.com/krymancer/airport-api-go/pkg/common/db"
	"github.com/spf13/viper"
)

func main() {
	viper.SetConfigFile("./pkg/common/config/envs/env.env")
	viper.ReadInConfig()

	port := viper.GetString("PORT")
	dbUrl := viper.GetString("DB_URL")

	r := gin.Default()
	db.InitDatabase(dbUrl)

	r.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"port":  port,
			"dbUrl": dbUrl,
		})
	})

}
