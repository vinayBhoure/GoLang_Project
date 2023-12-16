package userservice

import (
	"database/sql"

	"github.com/codesmith-dev/twitter/internal/gen/api/apiconnect"
)

var _ apiconnect.UserServiceHandler = &UserService{}

type UserService struct {
	db *sql.DB
}

func New(db *sql.DB) *UserService {
	return &UserService{
		db: db,
	}
}
