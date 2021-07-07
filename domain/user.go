package domain

import "github.com/APIGateway/domain/userType"

type User struct {
	Id       uint64
	UserType userType.UserType
	Username string
}
