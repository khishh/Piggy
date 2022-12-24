package repository

import (
	"github.com/khishh/personal-finance-app/graph/model"
	"github.com/khishh/personal-finance-app/pkg/models"
	"gorm.io/gorm"
)

type UserRepository interface {
	CreateUser(userInput *model.UserInput) (*models.User, error)

	GetOneUser(sub string) (*models.User, error)
}

type UserService struct {
	Db *gorm.DB
}

var _ UserRepository = &UserService{}

func NewUserService(db *gorm.DB) *UserService {
	return &UserService{
		Db: db,
	}
}

// have to add a pre-create hook to make sure that sub does not exist in the db

func (u *UserService) CreateUser(userInput *model.UserInput) (*models.User, error) {
	user := &models.User{
		Email:     userInput.Email,
		LastName:  userInput.LastName,
		FirstName: userInput.FirstName,
		Picture:   *userInput.Picture,
		Sub:       userInput.Sub,
	}
	err := u.Db.Create(user).Error
	return user, err
}

func (u *UserService) GetOneUser(sub string) (*models.User, error) {
	user := &models.User{}
	err := u.Db.Where("sub = ? ", sub).First(user).Error
	return user, err
}
