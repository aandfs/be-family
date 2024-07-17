package models

import "time"

type Customers struct {
	CstId         int         `gorm:"primaryKey;AUTO_INCREMENT" json:"cst_id"`
	NationalityId int         `gorm:"not null" json:"nationality_id"`
	CstName       string      `gorm:"type:char(50);not null" json:"cst_name"`
	CstDOB        time.Time   `gorm:"type:date;not null" json:"cst_dob"`
	CstPhoneNum   string      `gorm:"type:varchar(20);not null" json:"cst_phone_num"`
	CstEmail      string      `gorm:"type:varchar(50);" json:"cst_email"`
	Nationality   Nationality `gorm:"foreignKey:NationalityId;references:NationalityId;" json:"nationality"`
}

type CustomersResponse struct {
	CstId           int    `json:"cst_id"`
	NationalityId   int    `json:"nationality_id"`
	NationalityName string `json:"nationality_name"`
	CstName         string `json:"cst_name"`
	CstDOB          string `json:"cst_dob"`
	CstPhoneNum     string `json:"cst_phone_num"`
	CstEmail        string `json:"cst_email"`
}
