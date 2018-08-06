package main

import (
	"log"
	lib "taxi-taxes/routes-service/lib"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	// получить дистанцию и продолжительность маршрута
	// points lat1,lon1|lat2,lon2...|latn,lonn
	r.GET("/info", func(c *gin.Context) {
		pointsStr := c.Query("points")

		data, err := lib.GetData(pointsStr)
		if err != nil {
			log.Println(err)
			c.JSON(
				400,
				gin.H{
					"error": err.Error(),
				},
			)
		} else {
			c.JSON(200, data)
		}
	})
	r.Run() // listen and serve on 0.0.0.0:8080
}
