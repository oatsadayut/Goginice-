package routes

import (
	"github.com/gin-gonic/gin"
	usercontroller "github.com/oatsadayut/goginbasic/controllers/user"
)

//rg คือ Router Group ที่โยนเข้ามาจาก main.go function SetupRoter
func InitUserRoute(rg *gin.RouterGroup) {
	routerGroup := rg.Group("/users") //domain:3000/api/v1/users

	routerGroup.GET("/", usercontroller.GetAll)            //นำเข้า UserController เข้ามาแล้วเรียก Function GatAll
	routerGroup.POST("/register", usercontroller.Register) //นำเข้า UserController เข้ามาแล้วเรียก Function Register
	routerGroup.POST("/login", usercontroller.Login)       //นำเข้า UserController เข้ามาแล้วเรียก Function login
	routerGroup.GET("/:id", usercontroller.GetById)
	routerGroup.GET("/search", usercontroller.SearchByFullname) //{{domain_url/api/v1/users/search?fullname=<name>}}

}
