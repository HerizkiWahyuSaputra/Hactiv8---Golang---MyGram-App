package models

type Photo struct {
    ID        uint   `gorm:"primary_key" json:"id"`
    UserID    uint   `json:"user_id"`
    ImagePath string `json:"image_path"`
    CreatedAt string `json:"created_at"`
    UpdatedAt string `json:"updated_at"`
}
