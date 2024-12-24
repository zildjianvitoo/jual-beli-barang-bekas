package domain

import "time"

const (
	SELLER = "seller"
	BUYER  = "buyer"
)

type User struct {
	ID        uint      `json:"id" gorm:"primaryKey"`
	FirstName string    `json:"first_name"`
	LastName  string    `json:"last_name"`
	Email     string    `json:"email" gorm:"index;unique;not null"`
	Password  string    `json:"password"`
	Phone     string    `json:"phone"`
	Expiry    time.Time `json:"expiry"`
	Code      int       `json:"code"`
	Verified  bool      `json:"is_verified" gorm:"default:false"`
	UserType  string    `json:"user_type" gorm:"default:buyer"`
	CreateAt  time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}
