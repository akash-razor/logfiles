package models

import "github.com/jinzhu/gorm"

type(
	LogModel struct{
		gorm.Model
		Name string `json:"name"`
		Email string `json:"email"`
		Phone int `json:"phone"`
		Description string `json:"description"`
	}

)
