package services

import (
	"sewerage/internal/domain/models"
	"sewerage/internal/domain/repositories"
)

type UserService struct {
	userRepo repositories.UserRepository
}

func NewUserService(repo repositories.UserRepository) *UserService {
	return &UserService{userRepo: repo}
}

// func (service *UserService) GetByEmail(email string) (*models.User, error) {
//     userEntity, err := service.userRepo.GetByEmail(email)
//     if err != nil {
//         return nil, err
//     }
//     return &models.User{Email: userEntity.Email}, nil
// }

func (service *UserService) CreateUser(user *models.User) (models.User, error) {
    createdUser := models.User{Email: "test@tee.com",}
    return createdUser, nil
}

// func (service *UserService) CreateUser(user *models.User) (error) {
// 	existing, _ := service.userRepo.GetByEmail(user.User)
// 	if existing != nil {
// 		return nil, errors.New("email already exists")
// 	}

// 	return service.userRepo.Create(&user)
// }