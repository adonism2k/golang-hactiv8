package model

import (
	"time"
)

type User struct {
	ID        int       `gorm:"primarykey"`
	Username  string    `gorm:"not null"`
	Age       int       `gorm:"not null"`
	Email     string    `gorm:"not null"`
	Password  string    `gorm:"not null"`
	CreatedAt time.Time `gorm:"autoCreateTime"`
	UpdatedAt time.Time `gorm:"autoUpdateTime"`
}

// func (u *User) BeforeCreate() error {
// 	u.CreatedAt = time.Now()
// 	u.UpdatedAt = time.Now()
// 	return nil
// }

// func (u *User) BeforeUpdate() error {
// 	u.UpdatedAt = time.Now()
// 	return nil
// }

func (User) Register() {
	// RegisterRequest Model godoc
	// @Description RegisterRequest Model
	type Request struct {
		Age      int    `json:"age" example:"18" validate:"required,number"`
		Username string `json:"customer_name" example:"adnsm" validate:"required"`
		Email    string `json:"email" example:"abdianrizky11@gmail.com" validate:"required,email,min=6,max=32"`
		Password string `json:"password" example:"bcrypt hashed password" validate:"required,min=6,max=32"`
	} // @name RegisterRequest

	// RegisterResponse Model godoc
	// @Description RegisterResponse Model
	type Response struct {
		ID       int    `json:"id" example:"1"`
		Age      int    `json:"age" example:"18"`
		Username string `json:"customer_name" example:"adnsm"`
		Email    string `json:"email" example:"abdianrizky11@gmail.com"`
	} // @name RegisterResponse
}

func (User) Edit() {
	// UserEditRequest Model godoc
	// @Description UserEditRequest Model
	type Request struct {
		Username string `json:"customer_name" example:"adnsm" validate:"required"`
		Email    string `json:"email" example:"abdianrizky11@gmail.com" validate:"required,email,min=6,max=32"`
	} // @name UserEditRequest

	// UserEditResponse Model godoc
	// @Description UserEditResponse Model
	type Response struct {
		ID        int       `json:"id" example:"1"`
		Username  string    `json:"customer_name" example:"adnsm"`
		Age       int       `json:"age" example:"18"`
		Email     string    `json:"email" example:"abdianrizky11@gmail.com"`
		UpdatedAt time.Time `json:"updated_at" example:"2022-10-10T11:52:28.431369Z"`
	} // @name UserEditResponse
}
