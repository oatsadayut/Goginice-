package routes

import "github.com/gin-gonic/gin"

//rg คือ Router Group ที่โยนเข้ามาจาก main.go function SetupRoter
func InitHomeRoute(rg *gin.RouterGroup) {
	routerGroup := rg.Group("/")
	routerGroup.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"API Version": "1.0.2",
		})
	})
}
