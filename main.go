package main

import (
	"os"

	"github.com/gin-gonic/gin"
	"github.com/oatsadayut/goginbasic/configs"
	"github.com/oatsadayut/goginbasic/routes"
)

func main() {
	configs.Connection()              //Connect Database
	r := SetupRouter()                //Router
	r.Run(":" + os.Getenv("GO_PORT")) // listen and serve on 0.0.0.0:3001
}

//สร้าง Function แล้ว return Router ออกไป
func SetupRouter() *gin.Engine {
	router := gin.Default()

	apiV1 := router.Group("/api/v1") //localhost:3000/api/v1 กรณีถ้าอยากมีมากกว่า 1 เวอร์ชั่น ใน โฟนเดอร์ Route ให้มี Sub โฟนเดอร์ V1 , V2
	routes.InitHomeRoute(apiV1)      //Route Group เข้าไป
	routes.InitUserRoute(apiV1)
	return router
}
