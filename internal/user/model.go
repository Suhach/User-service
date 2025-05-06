package user

import "gorm.io/gorm"

type User struct {
	gorm.Model
	Email string `gorm:"type:varchar(255);unique"`
	Pass  string `gorm:"type:varchar(255);"`
	Tasks []Task `gorm:"foreignKey:UserID"`
}
