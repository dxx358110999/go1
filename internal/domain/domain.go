package domain

type User struct {
	UserId   int64
	Email    *string
	Username string
	Password string
	Phone    *string
}
