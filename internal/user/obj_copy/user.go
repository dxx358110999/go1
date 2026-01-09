package obj_copy

import (
	"dxxproject/agreed/domain"
	"dxxproject/agreed/dto"
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
func SignupDtoToDomain(dto *dto.UserSignup) *domain.User {
	user := &domain.User{
		UserId:   0,
		Email:    dto.Email,
		Username: dto.Username,
		Password: dto.Password,
		Phone:    dto.Phone,
	}
	return user
}

func LoginDtoToDomain(dto *dto.UserLogin) *model.User {
	user := &model.User{
		Username: dto.Username,
		Password: dto.Password,
	}
	return user
}
