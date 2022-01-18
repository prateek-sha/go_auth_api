package db

import (
	"github.com/prateek-sha/go_auth_api/src/client/cassandra"
	"github.com/prateek-sha/go_auth_api/src/domain/access_token"
	errors "github.com/prateek-sha/go_auth_api/src/utils/error"
)

const (
	queryGetAccessToken        = "SELECT access_token, user_id, client_id, expires FROM access_token WHERE access_token = ? "
	queryCreateAcccessToken    = "INSERT INTO access_token (access_token, user_id, client_id, expires) VALUES (?, ?, ?, ?)"
	queryUpdateExpirationToken = "UPDATE access_token SET expires=? WHERE access_token=?"
)

type DbRepoistory interface {
	GetById(string) (*access_token.AccessToken, *errors.RestError)
	Create(access_token.AccessToken) *errors.RestError
	UpdateExpirationTime(access_token.AccessToken) *errors.RestError
}

type dbRepoistory struct {
}

func New() DbRepoistory {
	return &dbRepoistory{}
}

func (db *dbRepoistory) GetById(id string) (*access_token.AccessToken, *errors.RestError) {
	session, sessionError := cassandra.GetSession()
	if sessionError != nil {
		return nil, errors.NewInternalServerError(sessionError.Error())
	}
	var result access_token.AccessToken
	if err := session.Query(queryGetAccessToken, id).Scan(&result.AccessToken, &result.UserId, &result.ClientId, &result.Expires); err != nil {
		return nil, errors.NewInternalServerError(err.Error())
	}
	defer session.Close()
	return &result, nil
}

func (db *dbRepoistory) Create(at access_token.AccessToken) *errors.RestError {
	session, sessionError := cassandra.GetSession()
	if sessionError != nil {
		return errors.NewInternalServerError(sessionError.Error())
	}
	if err := session.Query(queryCreateAcccessToken, at.AccessToken, at.UserId, at.ClientId, at.Expires).Exec(); err != nil {
		return errors.NewInternalServerError(err.Error())
	}
	defer session.Close()
	return nil
}

func (db *dbRepoistory) UpdateExpirationTime(at access_token.AccessToken) *errors.RestError {
	session, sessionError := cassandra.GetSession()
	if sessionError != nil {
		return errors.NewInternalServerError(sessionError.Error())
	}
	if err := session.Query(queryUpdateExpirationToken, at.Expires, at.AccessToken).Exec(); err != nil {
		return errors.NewInternalServerError(err.Error())
	}
	defer session.Close()
	return nil
}
