package user_hdl

import (
	"dxxproject/agreed/domain"
)

func LoginDtoToDomain(ul *UserLogin) *domain.User {
	user := &domain.User{
		Username: ul.Username,
		Password: ul.Password,
	}
	return user
}

func SignupDtoToDomain(us *UserSignup) *domain.User {
	user := &domain.User{
		UserId:   0,
		Email:    us.Email,
		Username: us.Username,
		Password: us.Password,
		Phone:    us.Phone,
	}
	return user
}
