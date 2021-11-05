package usercontroller //ตั้งชื่อใหม่

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/oatsadayut/goginbasic/configs"
	"github.com/oatsadayut/goginbasic/models"
)

func GetAll(c *gin.Context) {
	var users []models.User
	configs.DB.Order("id DESC").Find(&users)
	// configs.DB.Raw("setlect * from users").Scan(&users)

	c.JSON(200, gin.H{
		"data": users,
	})
}

func Register(c *gin.Context) {
	var input InputRegister
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	user := models.User{
		Fullname: input.Fullname,
		Email:    input.Email,
		Password: input.Password,
	}

	userExist := configs.DB.Where("email = ?", input.Email).First(&user)
	if userExist.RowsAffected == 1 {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "มีอีเมลนี้ในระบบแล้ว",
		})
		return
	}

	result := configs.DB.Debug().Create(&user)
	if result.Error != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": result.Error,
		})
		return
	}
	c.JSON(200, gin.H{
		"data": input,
	})
}

func Login(c *gin.Context) {
	c.JSON(200, gin.H{
		"data": "login",
	})
}

func GetById(c *gin.Context) {
	id := c.Param("id")

	var user models.User
	result := configs.DB.First(&user, id)

	if result.RowsAffected < 1 {
		c.JSON(http.StatusNotFound, gin.H{
			"error": "ไม่พบ User นี้",
		})
		return
	}
	if result.Error != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": result.Error,
		})
		return
	}

	c.JSON(200, gin.H{
		"data": user,
	})
}

func SearchByFullname(c *gin.Context) {
	fullname := c.Query("fullname")

	var users []models.User

	result := configs.DB.Where("fullname LIKE ?", "%"+fullname+"%").Scopes().Find(&users)
	if result.RowsAffected < 1 {
		return
	}

	c.JSON(200, gin.H{
		"data": fullname,
	})
}
