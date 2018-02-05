package requests

type LoginRequest struct {
	Username  string	`json:"Username"`
	Password  string	`json:"Password"`
}