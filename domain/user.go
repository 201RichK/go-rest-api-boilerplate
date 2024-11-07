package domain

import (
	"github.com/rs/xid"
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	ID        uint   `json:"id" gorm:"primaryKey"`
	XId       string `json:"x_id" gorm:"NOT NULL; UNIQUE_INDEX"`
	FirstName string `json:"firstname" gorm:"size:255"`
	LastName  string `json:"lastname" gorm:"size:255"`
	Email     string `json:"email" gorm:"NOT NULL; UNIQUE_INDEX"`
	Phone     string `json:"phone" gorm:"NOT NULL; UNIQUE_INDEX"`
	Password  string `json:"password" gorm:"NOT NULL"`
	Role      string `json:"role" gorm:"NOT_NULL;size:255;DEFAULT:'standard'"`
	Active    bool   `json:"article" gorm:"NOT NULL; DEFAULT: true"`
	Salt      string `json:"-" gorm:"NOT NULL"`
}

func (u *User) BeforeCreate(tx *gorm.DB) (err error) {
	u.XId = xid.New().String()
	return
}
