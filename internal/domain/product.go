package domain

import "time"

type Product struct {
	ID          uint      `json:"id" gorm:"PrimaryKey"`
	Name        string    `json:"name" gorm:"index;"`
	Description string    `json:"description"`
	CategoryId  uint      `json:"category_id"`
	ImageUrl    string    `json:"image_url" `
	Price       float64   `json:"price"`
	UserId      int       `json:"user_id"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}
