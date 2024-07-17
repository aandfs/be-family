package models

import "github.com/go-playground/validator/v10"

type Nationality struct {
	NationalityId   int    `gorm:"primaryKey;AUTO_INCREMENT;" json:"nationality_id"`
	NationalityName string `gorm:"type:varchar(50);not null" json:"nationality_name" validate:"required"`
	NationalityCode string `gorm:"type:char(2);not null;" json:"nationality_code" validate:"required"`
}

type NationalityResponse struct {
	NationalityId   int    `json:"nationality_id"`
	NationalityName string `json:"nationality_name"`
	NationalityCode string `json:"nationality_code"`
}

var validate = validator.New()

func (n *Nationality) Validate() error {
	return validate.Struct(n)
}
