package access_token

import (
	"strings"

	errors "github.com/prateek-sha/go_auth_api/src/utils/error"
)

type Repositry interface {
	GetById(string) (*AccessToken, *errors.RestError)
	Create(AccessToken) *errors.RestError
	UpdateExpirationTime(AccessToken) *errors.RestError
}

type Service interface {
	GetById(string) (*AccessToken, *errors.RestError)
	Create(AccessToken) *errors.RestError
	UpdateExpirationTime(AccessToken) *errors.RestError
}

type service struct {
	repositry Repositry
}

func NewService(repo Repositry) Service {
	return &service{
		repositry: repo,
	}
}

func (s *service) GetById(accessTokenId string) (*AccessToken, *errors.RestError) {
	accessTokenId = strings.TrimSpace(accessTokenId)
	if len(accessTokenId) == 0 {
		return nil, errors.NewBadRequestError("Cannot null")
	}
	accessToken, err := s.repositry.GetById(accessTokenId)
	if err != nil {
		return nil, errors.NewBadRequestError("failed")
	}
	return accessToken, nil
}

func (s *service) Create(at AccessToken) *errors.RestError {
	if err := at.Validate(); err != nil {
		return err
	}
	return s.repositry.Create(at)
}

func (s *service) UpdateExpirationTime(at AccessToken) *errors.RestError {
	if err := at.Validate(); err != nil {
		return err
	}
	return s.repositry.UpdateExpirationTime(at)
}
