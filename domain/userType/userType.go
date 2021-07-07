package userType

type UserType string

const (
	Administrator UserType = "ADMINISTRATOR"
	Regular       UserType = "REGULAR"
	Agent         UserType = "AGENT"
)

func (userType UserType) String() string {
	return string(userType)
}
