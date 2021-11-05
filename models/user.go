package models

import (
	"time"

	"github.com/matthewhartstonge/argon2"
	"gorm.io/gorm"
)

// `
type User struct {
	ID        uint   `json:"id" gorm:"primaryKey;autoIncrememnt;not null"`
	Fullname  string `gorm:"type:varchar(255);not null"`
	Email     string `json:"-" gorm:"type:varchar(255);not null;unique"`
	Password  string `gorm:"type:varchar(255);not null"`
	IsAdmit   bool   `gorm:"type:bool;default:false"`
	IsActive  bool   `json:"-" gorm:"type:bool;default:false"`
	CreatedAt time.Time
	UpdatedAt time.Time
}

func (user *User) BeforeCreate(db *gorm.DB) error {
	user.Password = hashPassword(user.Password)
	return nil
}

func hashPassword(password string) string {
	arg := argon2.DefaultConfig()
	encoded, _ := arg.HashEncoded([]byte(password))
	return string(encoded)
}

// type Tabler interface {
// 	TableName() string
// }

// // TableName overrides the table name used by User to `profiles`
// func (User) TableName() string {
// 	return "profiles"
// }
