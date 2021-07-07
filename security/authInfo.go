package security

import "github.com/APIGateway/domain/userType"

type AuthInfo struct {
	Token    string
	UserType userType.UserType
	UserId   uint64
}
