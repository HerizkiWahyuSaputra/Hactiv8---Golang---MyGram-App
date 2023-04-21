package models

type Comment struct {
    ID        uint   `gorm:"primary_key" json:"id"`
    UserID    uint   `json:"user_id"`
    PhotoID   uint   `json:"photo_id"`
    Comment   string `json:"comment"`
    CreatedAt string `json:"created_at"`
    UpdatedAt string `json:"updated_at"`
}
