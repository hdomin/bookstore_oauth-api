package db

import (
	"github.com/gocql/gocql"
	"github.com/hdomin/bookstore_oauth-api/src/clients/cassandra"
	accesstoken "github.com/hdomin/bookstore_oauth-api/src/domain/access_token"
	"github.com/hdomin/bookstore_oauth-api/src/utils/errors"
)

const (
	queryGetAccessToken    = "SELECT access_token, user_id, client_id, expires FROM access_tokens WHERE access_token = ?"
	queryCreateAccessToken = "INSERT INTO access_tokens (access_token, user_id, client_id, expires) values ( ?, ?, ?, ?)"
	queryUpdateExpires     = "UPDATE access_tokens SET expires=? WHERE access_token = ?"
)

func NewRepository() DbRepository {
	return &dbRepository{}
}

type DbRepository interface {
	GetById(string) (*accesstoken.AccessToken, *errors.RestErr)
	Create(accesstoken.AccessToken) *errors.RestErr
	UpdateExpirationTime(accesstoken.AccessToken) *errors.RestErr
}

type dbRepository struct {
}

func (r *dbRepository) GetById(id string) (*accesstoken.AccessToken, *errors.RestErr) {
	session := cassandra.GetSession()

	var result accesstoken.AccessToken
	if err := session.Query(queryGetAccessToken, id).Scan(&result.AccessToken, &result.UserId, &result.ClientId, &result.Expires); err != nil {
		if err == gocql.ErrNotFound {
			return nil, errors.NewNotFoundError("no access token found")
		}
		return nil, errors.NewInternalServerError(err.Error())
	}

	return &result, nil
}

func (r *dbRepository) Create(at accesstoken.AccessToken) *errors.RestErr {
	session := cassandra.GetSession()

	if err := session.Query(queryCreateAccessToken, at.AccessToken, at.UserId, at.ClientId, at.Expires).Exec(); err != nil {
		return errors.NewInternalServerError(err.Error())
	}

	return nil
}

func (r *dbRepository) UpdateExpirationTime(at accesstoken.AccessToken) *errors.RestErr {
	session := cassandra.GetSession()

	if err := session.Query(queryUpdateExpires, at.Expires, at.AccessToken).Exec(); err != nil {
		return errors.NewInternalServerError(err.Error())
	}

	return nil
}
