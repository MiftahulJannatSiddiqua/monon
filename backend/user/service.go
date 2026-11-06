package user

import "ecommerce/domain"

type service struct {
	userRepo UserRepo
}

func NewService(userRepo UserRepo) Service {
	return &service{
		userRepo: userRepo,
	}
}

func (svc *service) Create(user domain.User) (*domain.User, error) {
	usr, err := svc.userRepo.Create(user)
	if err != nil {
		return nil, err
	}
	if usr == nil {
		return nil, nil
	}

	return usr, nil
}

func (svc *service) Find(email string, pass string) (*domain.User, error) {
	usr, err := svc.userRepo.Find(email, pass)
	if err != nil {
		return nil, err
	}
	if usr == nil {
		return nil, nil
	}

	return usr, nil
}


//==================

func (svc *service) GetProfile(id int) (*domain.Profile, error) {
	return svc.userRepo.GetProfile(id)
}

func (svc *service) UpdateProfile(id int, bio, city *string) error {
	return svc.userRepo.UpdateProfile(id, bio, city)
}