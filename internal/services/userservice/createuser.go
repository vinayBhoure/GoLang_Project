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

// CreateUser implements apiconnect.UserServiceHandler.
func (s *UserService) CreateUser(ctx context.Context, req *connect.Request[api.CreateUserRequest]) (*connect.Response[api.User], error) {
	var err error
	q := sqlc.New(s.db)

	if req.Msg.GetUser() == nil {
		return nil, fmt.Errorf("invalid data format, expected: %v", api.CreateUserRequest{})
	} else if req.Msg.GetUser().FirstName == "" {
		return nil, connect.NewError(connect.CodeUnavailable, errors.New("missing firstname field from request body"))
	} else if req.Msg.GetUser().LastName == "" {
		return nil, connect.NewError(connect.CodeUnavailable, errors.New("missing lastname field from request body"))
	}

	user, err := q.CreateUser(ctx, sqlc.CreateUserParams{
		FirstName: req.Msg.GetUser().FirstName,
		LastName:  req.Msg.GetUser().LastName,
	})
	if err != nil {
		fmt.Println("error : ", err)
		return nil, connect.NewError(connect.CodeUnavailable, errors.New("failed to create user"))
	}

	newUser := &api.User{
		Id:        strconv.Itoa(int(user.ID)),
		FirstName: user.FirstName,
		LastName:  user.LastName,
	}

	response := &connect.Response[api.User]{
		Msg: newUser,
	}
	return response, nil
}
