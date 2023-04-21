package models

type Auth struct {
    ID       uint   `gorm:"primary_key" json:"id"`
    Username string `json:"username"`
    Password string `json:"password"`
}
