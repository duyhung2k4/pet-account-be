package service

import (
	"account/model"
	"account/repository"
	"account/utils"
	"errors"
)

type accessService struct {
	repo repository.AccessRepo
}

type AccessService interface {
	CheckUser(username string, password string) (*model.Profile, error)
	GetProfile(userId uint) (*model.Profile, error)
}

func (a *accessService) CheckUser(username string, password string) (*model.Profile, error) {
	profile, errProfile := a.repo.CheckUser(username)

	if errProfile != nil {
		return nil, errProfile
	}

	if !utils.CheckPasswordHash(password, profile.Credential.Password) {
		return nil, errors.New("wrong password")
	}

	return profile, nil
}

func (a *accessService) GetProfile(userId uint) (*model.Profile, error) {
	profile, err := a.repo.GetProfile(userId)
	return profile, err
}

func NewAccessService() AccessService {
	repo := repository.NewAccessRepo()
	return &accessService{
		repo: repo,
	}
}
