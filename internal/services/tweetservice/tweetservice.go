package tweetservice

import (
	"database/sql"

	"github.com/codesmith-dev/twitter/internal/gen/api/apiconnect"
)

var _ apiconnect.TweetServiceHandler = &TweetService{}

type TweetService struct {
	db *sql.DB
}

func New(db *sql.DB) *TweetService {
	return &TweetService{
		db: db,
	}
}
