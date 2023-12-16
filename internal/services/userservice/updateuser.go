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

// UpdateUser implements apiconnect.UserServiceHandler.
func (s *UserService) UpdateUser(ctx context.Context, req *connect.Request[api.UpdateUserRequest]) (*connect.Response[api.User], error) {
	q := sqlc.New(s.db)
	newUser := &api.User{
		Id:        req.Msg.GetId(),
		FirstName: req.Msg.GetFirstName(),
		LastName:  req.Msg.GetLastName(),
	}
	num, e := strconv.Atoi(newUser.Id)
	if e != nil {
		fmt.Println("conversion not done")
	}

	if newUser.FirstName == "" {
		return nil, connect.NewError(connect.CodeUnavailable, errors.New("missing firstname field from request body"))
	} else if newUser.LastName == "" {
		return nil, connect.NewError(connect.CodeUnavailable, errors.New("missing lastname field from request body"))
	}

	err := q.UpdateUser(ctx, sqlc.UpdateUserParams{
		FirstName: newUser.FirstName,
		LastName:  newUser.LastName,
		ID:        int32(num),
	})
	if err != nil {
		return nil, connect.NewError(connect.CodeUnknown, errors.New("failed to query for user"))
	}

	response := &connect.Response[api.User]{
		Msg: newUser,
	}
	return response, nil
}
