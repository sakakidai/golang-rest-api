package models

import "time"

type User struct {
	ID               uint      `json:"id" gorm:"primaryKey"`
	Email            string    `json:"email" gorm:"unique"`
	Password         string    `json:"password"`
	ConfirmedAt      time.Time `json:"confirmed_at" gorm:"default:null"`
	UnconfirmedEmail string    `json:"unconfirmed_email" gorm:"unique"`
	CreatedAt        time.Time `json:"created_at"`
	UpdatedAt        time.Time `json:"updated_at"`
}

type UserResponse struct {
	ID    uint   `json:"id" gorm:"primaryKey"`
	Email string `json:"email" gorm:"unique"`
}

type TemporaryUserResponse struct {
}
