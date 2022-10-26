package model

import (
	"context"
	"log"
	"strings"
	"time"
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

// func (u *User) BeforeCreate() error {
// 	u.CreatedAt = time.Now()
// 	u.UpdatedAt = time.Now()
// 	return nil
// }

// func (u *User) BeforeUpdate() error {
// 	u.UpdatedAt = time.Now()
// 	return nil
// }

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

func (u *User) Update(oldUser User, newUser User) User {
	ctx, cancel := context.WithTimeout(context.Background(), dbTimeout)
	defer cancel()

	tx := db.WithContext(ctx).Begin()

	result := tx.Model(&u).First(oldUser).Updates(newUser)
	if result.Error != nil {
		log.Println(result.Error)
		tx.Rollback()
		return User{}
	}

	tx.Commit()

	return newUser
}

func (u *User) Delete(user User) bool {
	ctx, cancel := context.WithTimeout(context.Background(), dbTimeout)
	defer cancel()

	tx := db.WithContext(ctx).Begin()

	result := tx.Delete(&user)
	if result.Error != nil {
		log.Println(result.Error)
		tx.Rollback()
		return false
	}

	tx.Commit()

	return true
}
