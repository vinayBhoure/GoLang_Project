package tweetservice

import (
	"context"
	"errors"
	"fmt"
	"strconv"

	"connectrpc.com/connect"
	"github.com/codesmith-dev/twitter/internal/gen/api"
	sqlcTweet "github.com/codesmith-dev/twitter/internal/gen/sqlc2"
)

func (t *TweetService) CreateTweet(ctx context.Context, req *connect.Request[api.CreateTweetRequest]) (*connect.Response[api.Tweet], error) {
	q := sqlcTweet.New(t.db)

	userid := req.Msg.GetTweet().User
	num, err := strconv.Atoi(userid)
	if err != nil {
		return nil, nil
	}

	tweet, err := q.CreateTweet(ctx, sqlcTweet.CreateTweetParams{
		Content: req.Msg.GetTweet().Content,
		Userid:  int64(num),
	})
	if err != nil {
		fmt.Println(err)
		return nil, connect.NewError(connect.CodeAborted, errors.New("failed to perform tweet query"))
	}

	newTweet := &api.Tweet{
		Id:      strconv.Itoa(int(tweet.ID)),
		Content: tweet.Content,
		User:    strconv.Itoa(int(tweet.Userid)),
	}
	response := &connect.Response[api.Tweet]{
		Msg: newTweet,
	}
	return response, nil
}
