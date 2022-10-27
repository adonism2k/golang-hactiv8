package model

import (
	"context"
	"log"
	"strings"
	"time"

	"gorm.io/gorm/clause"
)

// User Model godoc
// @Description User Model
type User struct {
	ID        int       `gorm:"primarykey" json:"id" example:"1"`
	Username  string    `gorm:"not null;uniqueIndex" json:"username" example:"admin"`
	Email     string    `gorm:"not null;uniqueIndex" json:"email" example:"admin@localhost"`
	Age       int       `gorm:"not null" json:"-" swaggerignore:"true"`
	Password  string    `gorm:"not null" json:"-" swaggerignore:"true"`
	CreatedAt time.Time `gorm:"autoCreateTime" json:"-" swaggerignore:"true"`
	UpdatedAt time.Time `gorm:"autoUpdateTime" json:"-" swaggerignore:"true"`
} // @name User

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

func (u User) FindByUsername(username string) User {
	ctx, cancel := context.WithTimeout(context.Background(), dbTimeout)
	defer cancel()

	result := db.WithContext(ctx).First(&u, "username = ?", strings.ToLower(username))
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
