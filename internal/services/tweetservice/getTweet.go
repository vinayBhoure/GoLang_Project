package tweetservice

import (
	"context"
	"strconv"

	"connectrpc.com/connect"
	"github.com/codesmith-dev/twitter/internal/gen/api"
	sqlcTweet "github.com/codesmith-dev/twitter/internal/gen/sqlc2"
)

func (t *TweetService) GetTweet(ctx context.Context, req *connect.Request[api.GetTweetRequest]) (*connect.Response[api.Tweet], error) {

	q := sqlcTweet.New(t.db)

	id := req.Msg.GetId()
	num, err := strconv.Atoi(id)
	if err != nil {
		return nil, err
	}

	tweet, err := q.GetTweet(ctx, int32(num))
	if err != nil {
		return nil, err
	}

	return &connect.Response[api.Tweet]{
		Msg: &api.Tweet{
			Id:      strconv.Itoa(int(tweet.ID)),
			Content: tweet.Content,
			User:    strconv.Itoa(int(tweet.Userid)),
		},
	}, nil
}
