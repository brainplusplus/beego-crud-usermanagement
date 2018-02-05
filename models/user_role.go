package models


type UserRole struct {
	Id 		  uint `gorm:"primary_key"`
	UserId    string `gorm:"type:varchar(50)"`
	RoleId    string `gorm:"type:varchar(50)"`
}
