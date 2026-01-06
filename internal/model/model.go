package model

type User struct {
	UserId   int64 `gorm:"column:user_id"`
	Email    *string
	Phone    *string
	Username string
	Password string
}

func (User) TableName() string {
	return "user"
}
