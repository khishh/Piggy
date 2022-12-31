package repository

import (
	"github.com/khishh/personal-finance-app/graph/model"
	"github.com/khishh/personal-finance-app/pkg/models"
	"gorm.io/gorm"
)

type UserRepository interface {
	CreateUserOnSignIn(userInput *model.UserInput) (*models.User, error)
	UpdateUserWithAccessToken(sub string, accessToken string) (*models.User, error)

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

func (u *UserService) CreateUserOnSignIn(userInput *model.UserInput) (*models.User, error) {

	// check if user exist. If so return the existing user
	user, err := u.GetOneUser(userInput.Sub)

	if err == nil {
		return user, err
	}

	// create a new user
	user = &models.User{
		Email:     userInput.Email,
		LastName:  userInput.LastName,
		FirstName: userInput.FirstName,
		Picture:   *userInput.Picture,
		Sub:       userInput.Sub,
	}
	err = u.Db.Create(user).Error
	return user, err
}

func (u *UserService) GetOneUser(sub string) (*models.User, error) {
	user := &models.User{}
	err := u.Db.Where("sub = ? ", sub).First(user).Error
	return user, err
}

func (u *UserService) UpdateUserWithAccessToken(sub string, accessToken string) (*models.User, error) {
	user, err := u.GetOneUser(sub)
	if err != nil {
		return nil, err
	}

	u.Db.Model(user).Update("access_token", accessToken)
	return user, nil
}
