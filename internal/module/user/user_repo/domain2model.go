package user_repo

import (
	"dxxproject/agreed/domain"
	"dxxproject/agreed/model"
)

func DomainToModel(dUser *domain.User) *model.User {

	return &model.User{
		UserId:   dUser.UserId,
		Email:    dUser.Email,
		Phone:    dUser.Phone,
		Username: dUser.Username,
		Password: dUser.Password,
	}
}
