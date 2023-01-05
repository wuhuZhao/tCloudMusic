package model

type User struct {
	ID       uint
	Username string
	Password string
	Email    string
}

func (User) TableName() string {
	return "user"
}
