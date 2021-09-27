package model

type User struct {
	Base
	UserName string `json:"user_name"`
	Password string `json:"password"`
}

func (User) TableName() string {
	return "user"
}
