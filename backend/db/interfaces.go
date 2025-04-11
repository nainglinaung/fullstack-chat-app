package db

type DBLayer interface {
	CreateUser(user *User) error
	GetUserByUsername(username string) (*User, error)
}
