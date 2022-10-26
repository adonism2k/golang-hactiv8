package model

import (
	"context"
	"log"
	"strings"
	"time"

	"gorm.io/gorm/clause"
)

type User struct {
	ID        int       `gorm:"primarykey"`
	Username  string    `gorm:"not null,uniqueIndex"`
	Age       int       `gorm:"not null"`
	Email     string    `gorm:"not null,uniqueIndex"`
	Password  string    `gorm:"not null"`
	CreatedAt time.Time `gorm:"autoCreateTime"`
	UpdatedAt time.Time `gorm:"autoUpdateTime"`
}

// UserPhotoResponse Model godoc
// @Description UserPhotoResponse Model
type UserPhoto struct {
	ID       int    `json:"id" example:"1"`
	Age      int    `json:"age" example:"18"`
	Username string `json:"username" example:"adnsm"`
	Email    string `json:"email" example:"abdianrizky11@gmail.com"`
} // @name UserPhotoResponse

func (u User) FindByEmail(email string) User {
	ctx, cancel := context.WithTimeout(context.Background(), dbTimeout)
	defer cancel()

	result := db.WithContext(ctx).First(&u, "email = ?", strings.ToLower(email))
	if result.Error != nil {
		log.Println(result.Error)
		return User{}
	}

	return u
}

func (u *User) Create(user User) User {
	ctx, cancel := context.WithTimeout(context.Background(), dbTimeout)
	defer cancel()

	tx := db.WithContext(ctx).Begin()

	result := tx.Create(&user)
	if result.Error != nil {
		log.Println(result.Error)
		tx.Rollback()
		return User{}
	}

	tx.Commit()

	return user
}

func (u *User) Update(id int, newUser User) User {
	ctx, cancel := context.WithTimeout(context.Background(), dbTimeout)
	defer cancel()

	tx := db.WithContext(ctx).Begin()

	var user User
	result := tx.Model(&user).Clauses(clause.Returning{}).Where("id = ?", id).Updates(newUser)
	if result.Error != nil {
		log.Println(result.Error)
		tx.Rollback()
		return User{}
	}

	tx.Commit()

	return user
}

func (u *User) Delete(user User) error {
	ctx, cancel := context.WithTimeout(context.Background(), dbTimeout)
	defer cancel()

	tx := db.WithContext(ctx).Begin()

	result := tx.Delete(&user)
	if result.Error != nil {
		log.Println(result.Error)
		tx.Rollback()
		return result.Error
	}

	tx.Commit()

	return nil
}
