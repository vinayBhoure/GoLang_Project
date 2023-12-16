package userservice

import (
	"context"
	"errors"
	"fmt"
	"strconv"

	"connectrpc.com/connect"
	"github.com/codesmith-dev/twitter/internal/gen/api"
	"github.com/codesmith-dev/twitter/internal/gen/sqlc"
	"google.golang.org/protobuf/types/known/emptypb"
)

// DeleteUser implements apiconnect.UserServiceHandler.
func (s *UserService) DeleteUser(ctx context.Context, req *connect.Request[api.DeleteUserRequest]) (*connect.Response[emptypb.Empty], error) {
	q := sqlc.New(s.db)
	UserId := req.Msg.GetId()

	if UserId == "" {
		return nil, connect.NewError(connect.CodeCanceled, errors.New("user id cannot be empty"))
	}

	num, e := strconv.Atoi(UserId)
	if e != nil {
		fmt.Println("Error: ", e)
	}

	err := q.DeleteUser(ctx, int32(num))
	if err != nil {
		return nil, connect.NewError(connect.CodeCanceled, errors.New("user not deleted"))
	}

	return &connect.Response[emptypb.Empty]{}, nil
}
