package user

import (
	"ecommerce/domain"
	userHandler "ecommerce/rest/handlers/user" //alias
)

type Service interface {
	userHandler.Service //embadding
}

type UserRepo interface {
	Create(user domain.User) (*domain.User, error)
	Find(email string, pass string) (*domain.User, error)
}