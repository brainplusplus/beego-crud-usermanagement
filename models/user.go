package models

import (
	"time"
)

// Add fields `ID`, `CreatedAt`, `UpdatedAt`, `DeletedAt`
type User struct {
	Id        string `gorm:"type:varchar(50);primary_key"`
	Username  string    `gorm:"not null;type:varchar(100)"`
	Password  string    `gorm:"not null;type:varchar(100)"`
	Name      string    `gorm:"not null;type:varchar(50)"`
	Phone     string    `gorm:"type:varchar(40);null"`
	Email     string    `gorm:"type:varchar(100);null"`
	CreatedAt time.Time
	UpdatedAt time.Time
	//DeletedAt *time.Time
	status    int
	Roles     []Role `gorm:"many2many:user_role;"`
}

func (t *User) TableName() string {
	return "user"
}

