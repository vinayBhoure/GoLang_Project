package userservice

import (
	"context"
	"errors"
	"fmt"
	"math"
	"strconv"

	"connectrpc.com/connect"
	"github.com/codesmith-dev/twitter/internal/gen/api"
	"github.com/codesmith-dev/twitter/internal/gen/sqlc"
)

type PageInfo struct {
	pgsize  int32
	pgtoken string
}

// ListUsers implements apiconnect.UserServiceHandler.
func (s *UserService) ListUsers(ctx context.Context, req *connect.Request[api.ListUserRequest]) (*connect.Response[api.ListUserResponse], error) {
	q := sqlc.New(s.db)
	pgInfo := PageInfo{
		pgsize:  req.Msg.GetPageSize(),
		pgtoken: req.Msg.GetPageToken(),
	}

	var cnt int32
	e := s.db.QueryRow(`SELECT COUNT(*) FROM users`).Scan(&cnt)
	if e != nil {
		return nil, connect.NewError(connect.CodeInternal, errors.New("failed to query for user"))
	}

	var totalPages int32
	if pgInfo.pgsize < cnt {
		totalPages = int32(math.Ceil(float64(cnt / pgInfo.pgsize)))
	} else {
		totalPages = 1
	}

	users, err := q.ListUser(ctx)
	if err != nil {
		return nil, connect.NewError(connect.CodeAborted, errors.New("failed to query for user"))
	}

	i, err := strconv.Atoi(pgInfo.pgtoken)
	if err != nil {
		fmt.Println("Error:", err)
	}

	low := ((int32(i - 1)) * pgInfo.pgsize)
	high := (int32(i) * pgInfo.pgsize)

	var apiusers []*api.User
	for _, u := range users {

		a := &api.User{
			Id:        strconv.Itoa(int(u.ID)),
			FirstName: u.FirstName,
			LastName:  u.LastName,
		}
		apiusers = append(apiusers, a)
	}

	var newSlice []*api.User
	var nextpgToken string

	if totalPages == int32(i) {

		newSlice = apiusers[low:]
		nextpgToken = "no more pages"
	} else if totalPages > int32(i) {

		newSlice = apiusers[low:high]
		nextpgToken = strconv.Itoa(i + 1)
	} else if int32(i) > totalPages {

		return nil, connect.NewError(connect.CodeUnknown, errors.New("page doesn't exists"))
	}

	output := &api.ListUserResponse{
		Users:         newSlice,
		NextPageToken: nextpgToken,
	}

	response := &connect.Response[api.ListUserResponse]{
		Msg: output,
	}

	return response, nil
}
