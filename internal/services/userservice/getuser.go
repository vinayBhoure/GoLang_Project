package userservice

import (
	"context"
	"errors"
	"fmt"
	"strconv"

	"connectrpc.com/connect"
	"github.com/codesmith-dev/twitter/internal/gen/api"
	"github.com/codesmith-dev/twitter/internal/gen/sqlc"
)

// GetUser implements apiconnect.UserServiceHandler.
func (s *UserService) GetUser(ctx context.Context, req *connect.Request[api.GetUserRequest]) (*connect.Response[api.User], error) {
	q := sqlc.New(s.db)

	getUserId := req.Msg.GetId()
	if getUserId == "" {
		return nil, connect.NewError(connect.CodeAborted, errors.New("user id cannot be empty"))
	}
	fmt.Println("Got the requested ID: ", getUserId)

	num, e := strconv.Atoi(getUserId)
	if e != nil {
		fmt.Println("conversion not done")
	}

	user, err := q.GetUser(ctx, int32(num))
	if err != nil {
		return nil, connect.NewError(connect.Code(404), errors.New("can't retrieve data"))
	}

	newUser := &api.User{
		Id:        getUserId,
		FirstName: user.FirstName,
		LastName:  user.LastName,
	}

	response := &connect.Response[api.User]{
		Msg: newUser,
	}

	return response, nil
}
