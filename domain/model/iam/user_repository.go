package iam

type IUserRepository interface {
	FindAll() []*User
	FindOne() *User
	Save()
	Remove()
}
