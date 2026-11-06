// package user

// import (
// 	"ecommerce/domain"
// 	userHandler "ecommerce/rest/handlers/user" //alias
// )

// type Service interface {
// 	userHandler.Service //embadding
// }

// type UserRepo interface {
// 	Create(user domain.User) (*domain.User, error)
// 	Find(email string, pass string) (*domain.User, error)
// }


package user

import (
	"ecommerce/domain"
	userHandler "ecommerce/rest/handlers/user"
)

type Service interface {
	userHandler.Service
}

type UserRepo interface {
	Create(user domain.User) (*domain.User, error)
	Find(email string, pass string) (*domain.User, error)
	GetProfile(id int) (*domain.Profile, error)
	UpdateProfile(id int, bio, city *string) error
}