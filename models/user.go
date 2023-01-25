package models

import (
	"time"

	"github.com/satori/go.uuid"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type User struct {
	Id        uuid.UUID `json:"id" gorm:"primary_key"`
	FirstName string    `json:"f_name"`
	LastName  string    `json:"l_name"`
	Email     string    `json:"email" gorm:"unique"`
	Password  []byte    `json:"-"`
	Phone     string    `json:"phone"`
	RoleId    uint      `json:"role_id"`
	Role      Role      `json:"role" gorm:"foreignKey:RoleId"`
	Address   []Address `json:"address" gorm:"foreignKey:UserID"`
	Product   []Product `json:"product" gorm:"foreignKey:UserID"`
	CreatedAt time.Time
	UpdatedAt time.Time
}

func (user *User) SetPassword(password string) {
	hashedPassword, _ := bcrypt.GenerateFromPassword([]byte(password), 14)
	user.Password = hashedPassword
}

func (user *User) ComparePassword(password string) error {
	return bcrypt.CompareHashAndPassword(user.Password, []byte(password))
}

func (user *User) Count(db *gorm.DB) int64 {

	var total int64
	
	db.Model(&User{}).Count(&total)

	return total
}

func (user *User) Take(db *gorm.DB, limit int, offset int) interface{} {
	var products []User

	db.Preload("Role").Offset(offset).Limit(limit).Find(&products)

	return products
}
