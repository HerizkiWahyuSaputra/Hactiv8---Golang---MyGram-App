package models

type SocialMedia struct {
    ID         uint   `gorm:"primary_key" json:"id"`
    Name       string `json:"name"`
    Link       string `json:"link"`
    Icon       string `json:"icon"`
    CreatedAt  string `json:"created_at"`
    UpdatedAt  string `json:"updated_at"`
}
