package model

import (
	"time"
)

type User struct {
	ID        int       `gorm:"primarykey" json:"id" example:"1"`
	Username  string    `gorm:"type:varchar(100);not null" json:"customer_name" example:"John Doe"`
	Age       int       `gorm:"not null" json:"age" example:"2022-10-10T11:52:28.431369Z"`
	Email     string    `gorm:"not null" json:"email" example:"2022-10-10T11:52:28.431369Z"`
	Password  string    `gorm:"not null" json:"password" example:"2022-10-10T11:52:28.431369Z"`
	CreatedAt time.Time `gorm:"autoCreateTime" swaggerignore:"true" json:"-"`
	UpdatedAt time.Time `gorm:"autoUpdateTime" swaggerignore:"true" json:"-"`
}

func (u *User) Register() {
	// RegisterRequest Model godoc
	// @Description RegisterRequest Model
	type Request struct {
		Age      int    `json:"age" example:"18"`                          // Age
		Username string `json:"customer_name" example:"adnsm"`             // Customer Name
		Email    string `json:"email" example:"abdianrizky11@gmail.com"`   // Email
		Password string `json:"password" example:"bcrypt hashed password"` // Password
	} // @name RegisterRequest

	// RegisterResponse Model godoc
	// @Description RegisterResponse Model
	type Response struct {
		ID       int    `json:"id" example:"1"`                          // User ID
		Age      int    `json:"age" example:"18"`                        // Age
		Username string `json:"customer_name" example:"adnsm"`           // Customer Name
		Email    string `json:"email" example:"abdianrizky11@gmail.com"` // Email
	} // @name RegisterResponse
}

func (u *User) Login() {
	// LoginRequest Model godoc
	// @Description LoginRequest Model
	type Request struct {
		Email    string `json:"email" example:"`                           // Email
		Password string `json:"password" example:"bcrypt hashed password"` // Password
	} // @name LoginRequest

	// LoginResponse Model godoc
	// @Description LoginResponse Model
	type Response struct {
		Token string `json:"token" example:"eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJzdWIiOiIxMjM0NTY3ODkwIiwibmFtZSI6IkpvaG4gRG9lIiwiaWF0IjoxNTE2MjM5MDIyfQ.SflKxwRJSMeKKF2QT4fwpMeJf36POk6yJV_adQssw5c"` // Token
	} // @name LoginResponse
}

func (u *User) Edit() {
	// EditRequest Model godoc
	// @Description EditRequest Model
	type Request struct {
		Username string `json:"customer_name" example:"adnsm"`           // Customer Name
		Email    string `json:"email" example:"abdianrizky11@gmail.com"` // Email
	}

	// EditResponse Model godoc
	// @Description EditResponse Model
	type Response struct {
		ID        int       `json:"id" example:"1"`                                   // User ID
		Username  string    `json:"customer_name" example:"adnsm"`                    // Customer Name
		Age       int       `json:"age" example:"18"`                                 // Age
		Email     string    `json:"email" example:"abdianrizky11@gmail.com"`          // Email
		UpdatedAt time.Time `json:"updated_at" example:"2022-10-10T11:52:28.431369Z"` // Updated At
	} // @name EditResponse
}
