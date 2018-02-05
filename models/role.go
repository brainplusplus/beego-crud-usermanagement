package models

import (

)

type Role struct {
	Id 	  string `gorm:"type:varchar(50);primary_key"`
	Name  string `gorm:"type:varchar;size(50)"`
	Notes string `gorm:"type:text;null"`
}

func (t *Role) TableName() string {
	return "role"
}