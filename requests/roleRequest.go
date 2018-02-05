package requests

type RoleRequestList struct {
	Search 		string
	OrderBy		string
	Limit 		int64
	Offset 		int64
	Page		int64
}

type RoleRequest struct {
	Id        string
	Name      string
	Notes     string
}

