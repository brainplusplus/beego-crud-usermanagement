package models

import (
	"time"
)

type Model struct {
	Id        string `gorm:"type:varchar(50);primary_key"`
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt *time.Time
	status    int
}




