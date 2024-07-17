package models

import "time"

type FamilyList struct {
	FlId       int       `gorm:"primaryKey;AUTO_INCREMENT;" json:"fl_id"`
	CstId      int       `gorm:"not null" json:"cst_id"`
	FlRelation string    `gorm:"type:varchar(50);not null;" json:"fl_relation"`
	FlName     string    `gorm:"type:varchar(50);not null;" json:"fl_name"`
	FlDOB      time.Time `gorm:"not null;" json:"fl_dob"`
	Customers  Customers `gorm:"foreignKey:CstId;references:CstId" json:"customer"`
}

type FamilyListResponse struct {
	FlId       int    `json:"fl_id"`
	CstId      int    `json:"cst_id"`
	CstName    string `json:"cst_name"`
	FlRelation string `json:"fl_relation"`
	FlName     string `json:"fl_name"`
	FlDOB      string `json:"fl_dob"`
}
