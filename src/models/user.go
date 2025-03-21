package models

type User struct {
	Username    string `json:"user_username"`
	FirstName   string `json:"user_firstname"`
	LastName    string `json:"user_lastname"`
	Job         string `json:"user_job"`
	Email       string `json:"user_email"`
	PhoneNumber string `json:"user_phone_number"`
	Password    string `json:"user_password"`
}
