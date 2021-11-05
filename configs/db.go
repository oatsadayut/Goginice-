package configs

import (
	"fmt"
	"os"

	"github.com/oatsadayut/goginbasic/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func Connection() {
	dsn := os.Getenv("DATABASE_DSN")
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		fmt.Println("--------- Database Error ----------")
		fmt.Println(err.Error())
		fmt.Println("-----------------------------------")
		panic(err) //ของ Go ใช้ Panic ออกไปได้เลยจะไม่ทำงานต่อเบรกไปเลย
	}

	// Migration
	db.AutoMigrate(&models.User{})

	DB = db

	fmt.Println("--------- Database Success ----------")
	fmt.Println("Connect Database Success")
	fmt.Println("-------------------------------------")
}
