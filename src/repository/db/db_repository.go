package db

import (
	"github.com/hdomin/bookstore_oauth-api/src/clients/cassandra"
	accesstoken "github.com/hdomin/bookstore_oauth-api/src/domain/access_token"
	"github.com/hdomin/bookstore_oauth-api/src/utils/errors"
)

func NewRepository() DbRepository {
	return &dbRepository{}
}

type DbRepository interface {
	GetById(string) (*accesstoken.AccessToken, *errors.RestErr)
}

type dbRepository struct {
}

func (r *dbRepository) GetById(id string) (*accesstoken.AccessToken, *errors.RestErr) {
	//TODO: Implement Database access token
	_, err := cassandra.GetSession()
	if err != nil {
		panic(err)
	}
	return nil, errors.NewInternalServerError("data base connection not implemented!")
}
