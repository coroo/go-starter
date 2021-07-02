package entity

import (
	"time"
)

type UserProfile struct {
	ID               int            `gorm:"AUTO_INCREMENT" json:"id"`
	Email 			 string    		`gorm:"unique" json:"email"`
	Password 		 string    		`json:"password"`
	Address 		 string    		`gorm:"type:text" json:"address"`
	CreatedAt        time.Time      `json:"created_at"`
	UpdatedAt        time.Time      `json:"updated_at"`
}

type UserProfileCreate struct {
	ID               int            `json:"-"`
	Email 			 string    		`json:"email"`
	Address 		 string    		`gorm:"type:text" json:"address"`
	Password 		 string    		`json:"password"`
}

type UserProfileDelete struct {
	ID int `json:"id"`
}

type UserProfileLogin struct {
	Email 			 string    		`json:"email"`
	Password 		 string    		`json:"password"`
}
