package access_token

import errors "github.com/prateek-sha/go_auth_api/src/utils/error"

type Repositry interface {
	GetById(string) (*AccessToken, *errors.RestError)
}

type Service interface {
	GetUserById(string) (*AccessToken, *errors.RestError)
}

type service struct {
	repositry Repositry
}

func NewService(repo Repositry) Service {
	return &service{
		repositry: repo,
	}
}

func (s *service) GetUserById(accessTokenId string) (*AccessToken, *errors.RestError) {
	return s.repositry.GetById(accessTokenId)
}
