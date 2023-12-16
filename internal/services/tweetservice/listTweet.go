package tweetservice

import (
	"context"
	"fmt"
	"math"
	"strconv"

	"connectrpc.com/connect"
	"github.com/codesmith-dev/twitter/internal/gen/api"
	sqlcTweet "github.com/codesmith-dev/twitter/internal/gen/sqlc2"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type PageInfo struct {
	userid  string
	pgsize  int32
	pgtoken string
}

func (t *TweetService) ListTweets(ctx context.Context, req *connect.Request[api.ListTweetRequest]) (*connect.Response[api.ListTweetResponse], error) {
	// panic("unimplemented")
	q := sqlcTweet.New(t.db)

	pgInfo := PageInfo{
		userid:  req.Msg.GetUser(),
		pgsize:  req.Msg.GetPageSize(),
		pgtoken: req.Msg.GetPageToken(),
	}

	fmt.Println(pgInfo)

	var cnt int32
	e := t.db.QueryRow(`select COUNT(*) from tweets WHERE userid = $1`, pgInfo.userid).Scan(&cnt)
	if e != nil {
		return nil, status.Errorf(codes.Internal, "failed to query for user: %v", e)
	}

	if cnt == 0 {
		return nil, status.Errorf(codes.NotFound, "user does not exist")
	}

	num, err := strconv.Atoi(pgInfo.userid)
	if err != nil {
		return nil, err
	}

	var totalPages int32
	if pgInfo.pgsize < cnt {
		totalPages = int32(math.Ceil(float64(cnt / pgInfo.pgsize)))
	} else {
		totalPages = 1
	}

	tweets, err := q.ListTweet(ctx, int64(num))
	if err != nil {
		return nil, err
	}

	i, err := strconv.Atoi(pgInfo.pgtoken)
	if err != nil {
		fmt.Println("Error:", err)
	}

	low := ((int32(i - 1)) * pgInfo.pgsize)
	high := (int32(i) * pgInfo.pgsize)

	var apitweets []*api.Tweet
	for _, u := range tweets {

		a := &api.Tweet{
			Id:      strconv.Itoa(int(u.ID)),
			Content: u.Content,
			User:    pgInfo.userid,
		}
		apitweets = append(apitweets, a)
	}

	var newSlice []*api.Tweet
	var nextpgToken string

	if totalPages == int32(i) {
		// we have to convert sqlc.user to api.user
		newSlice = apitweets[low:]
		nextpgToken = "no more pages"
	} else if totalPages > int32(i) {

		newSlice = apitweets[low:high]
		nextpgToken = strconv.Itoa(i + 1)
	} else if int32(i) > totalPages {

		nextpgToken = "page does not exist"
	}

	output := &api.ListTweetResponse{
		Tweets:        newSlice,
		NextPageToken: nextpgToken,
	}

	response := &connect.Response[api.ListTweetResponse]{
		Msg: output,
	}

	return response, nil
}
