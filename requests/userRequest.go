package requests

import "time"

type UserRequestList struct {
	Search 		string
	RoleId		string
	OrderBy		string
	Limit 		int64
	Offset 		int64
	Page		int64
	StartDate 	time.Time
	EndDate		time.Time
}

type UserRequest struct {
	Id        string
	Username  string
	Password  string
	Name      string
	Phone     string
	Email     string
	Status    int
	Roles 	  []string
}