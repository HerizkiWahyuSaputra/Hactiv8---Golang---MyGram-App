package models

import "time"

type User struct {
    ID           uint      `gorm:"primary_key" json:"id"`
    FullName     string    `json:"full_name"`
    Email        string    `json:"email"`
    Password     string    `json:"password"`
    ProfileImage string    `json:"profile_image"`
    CreatedAt    time.Time `json:"created_at"`
    UpdatedAt    time.Time `json:"updated_at"`
}