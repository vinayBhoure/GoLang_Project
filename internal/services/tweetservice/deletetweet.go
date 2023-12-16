package tweetservice

import (
	"context"
	"errors"
	"fmt"
	"strconv"

	"connectrpc.com/connect"
	"github.com/codesmith-dev/twitter/internal/gen/api"
	sqlcTweet "github.com/codesmith-dev/twitter/internal/gen/sqlc2"
	"google.golang.org/protobuf/types/known/emptypb"
)

func (t *TweetService) DeleteTweet(ctx context.Context, req *connect.Request[api.DeleteTweetRequest]) (*connect.Response[emptypb.Empty], error) {
	// panic("unimplementedt ")
	q := sqlcTweet.New(t.db)

	id := req.Msg.GetId()

	num, err := strconv.Atoi(id)
	if err != nil {
		return nil, nil
	}

	err = q.DeleteTweet(ctx, int32(num))
	if err != nil {
		return nil, connect.NewError(connect.CodeCanceled, errors.New("tweet not found"))
	}
	fmt.Println("Deleted tweet")
	return &connect.Response[emptypb.Empty]{}, nil
}
