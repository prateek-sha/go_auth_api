package db

import (
	"github.com/prateek-sha/go_auth_api/src/domain/access_token"
	errors "github.com/prateek-sha/go_auth_api/src/utils/error"
)

type DbRepoistory interface {
	GetUserById(string) (*access_token.AccessToken, *errors.RestError)
}

type dbRepoistory struct {
}

func New() DbRepoistory {
	return &dbRepoistory{}
}

func (db *dbRepoistory) GetUserById(string) (*access_token.AccessToken, *errors.RestError) {
	return nil, nil
}
