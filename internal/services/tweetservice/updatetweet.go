package tweetservice

import (
	"context"
	"strconv"

	"connectrpc.com/connect"
	"github.com/codesmith-dev/twitter/internal/gen/api"
	sqlcTweet "github.com/codesmith-dev/twitter/internal/gen/sqlc2"
)

func (t *TweetService) UpdateTweet(ctx context.Context, req *connect.Request[api.UpdateTweetRequest]) (*connect.Response[api.Tweet], error) {
	// panic("unimplemented")

	q := sqlcTweet.New(t.db)
	id := req.Msg.GetId()

	num, err := strconv.Atoi(id)
	if err != nil {
		return nil, err
	}
	e := q.UpdateTweet(ctx, sqlcTweet.UpdateTweetParams{
		Content: req.Msg.GetContent(),
		ID:      int32(num),
	})
	if e != nil {
		return nil, e
	}

	tweet, err := q.GetTweet(ctx, int32(num))
	if err != nil {
		return nil, err
	}

	apitweet := api.Tweet{
		Id:      strconv.Itoa(int(tweet.ID)),
		Content: tweet.Content,
		User:    strconv.Itoa(int(tweet.Userid)),
	}
	response := &connect.Response[api.Tweet]{
		Msg: &apitweet,
	}
	return response, nil
}
