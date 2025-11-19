package models

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type User struct {
	ID       uuid.UUID `json:"id"`
	Name     string    `form:"name" json:"name"`
	Email    string    `form:"email" json:"email"`
	Password string    `form:"password" json:"password"`
	Role     string    `json:"role"`
}

func (User) TableName() string {
	return "user"
}

func (u *User) BeforeCreate(tx *gorm.DB) (err error) {
	u.ID = uuid.New()
	return
}
