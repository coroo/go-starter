package entity

import (
	"time"
	"gorm.io/gorm"
)

type User struct {
	ID               int            `gorm:"AUTO_INCREMENT" json:"id"`
	Uuid 			 string	    	`gorm:"unique" json:"uuid"`
	Email 			 string    		`gorm:"unique" json:"email"`
	Password 		 string    		`json:"password"`
	Name 		 	 string    		`gorm:"type:text" json:"name"`
	CreatedAt        time.Time      `json:"created_at"`
	UpdatedAt        time.Time      `json:"updated_at"`
	DeletedAt        gorm.DeletedAt `json:"deleted_at"`
}

type UserCreate struct {
	ID               int            `json:"-"`
	uuid 			 string	    	`gorm:"type:uuid;default:uuid_generate_v4()"`
	Email 			 string    		`json:"email"`
	Name 		 	 string    		`gorm:"type:text" json:"name"`
	Password 		 string    		`json:"password"`
}

type UserDelete struct {
	ID int `json:"id"`
}

type UserLogin struct {
	Email 			 string    		`json:"email"`
	Password 		 string    		`json:"password"`
}

type Token struct {
	AccessToken 	string `json:"access_token"`
	RefreshToken 	string `json:"refresh_token"`
}

type TokenReqBody struct {
	RefreshToken 	string `json:"refresh_token"`
}
