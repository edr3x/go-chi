package schemas

import (
	"time"
)

type Role string

const (
	Admin    Role = "admin"
	Customer Role = "customer"
)

type User struct {
	ID        string `json:"id" gorm:"type:uuid;default:gen_random_uuid();primaryKey"`
	Name      string `json:"name" gorm:"not null"`
	Age       int    `json:"age" gorm:"not null"`
	Email     string `json:"username" gorm:"unique;not null;index:,option:CONCURRENTLY"`
	Role      Role   `json:"role" gorm:"default:customer"`
	Password  string `json:"password" gorm:"not null"`
	CreatedAt time.Time
	UpdatedAt time.Time
}
