package models

type User struct {
	User_ID   int64  `json:"user_id"`
	Username  string `json:"username"`
	Password  string `json:"password"`
	CreatedOn string `json:"created_on"`
}
