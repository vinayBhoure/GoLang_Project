package twitterapp_test

import (
	"context"
	"fmt"
	"net/http"
	"testing"

	"connectrpc.com/connect"
	"github.com/codesmith-dev/twitter/internal/gen/api"
	"github.com/codesmith-dev/twitter/internal/gen/api/apiconnect"
)

func TestTwitterApp(t *testing.T) {
	// twitterapp.Run()

	testUser := &api.User{
		FirstName: "Test",
		LastName:  "User",
	}

	// User Client
	userClient := apiconnect.NewUserServiceClient(http.DefaultClient, "http://127.0.0.1:8080")

	//Creating User
	createUser, err := userClient.CreateUser(context.Background(), connect.NewRequest(&api.CreateUserRequest{
		User: testUser,
	}))
	if err != nil {
		t.Errorf("CreateUser failed: %v", err)
		t.FailNow()
	}
	id := createUser.Msg.Id
	notEmpty(t, id)
	fmt.Println("User ID: ", id)

	//Getting User
	getUser, err := userClient.GetUser(context.Background(), connect.NewRequest(&api.GetUserRequest{
		Id: id,
	}))
	if err != nil {
		t.Errorf("GetUser failed: %v", err)
		t.FailNow()
	}
	equalStr(t, getUser.Msg.FirstName, testUser.FirstName)
	equalStr(t, getUser.Msg.LastName, testUser.LastName)

	fmt.Println("User: ", getUser.Msg.FirstName, getUser.Msg.LastName)

	//Updating User
	fname := "Updated"
	lname := "User"
	fmt.Println("id before update user", id)
	updateUser, err := userClient.UpdateUser(context.Background(), connect.NewRequest(&api.UpdateUserRequest{
		Id:        id,
		FirstName: &fname,
		LastName:  &lname,
	}))
	if err != nil {
		t.Errorf("UpdateUserTest failed: %v", err)
		t.FailNow()
	}
	equalStr(t, updateUser.Msg.FirstName, "Updated")
	equalStr(t, updateUser.Msg.LastName, "User")

	// List Users
	pgsize := 10
	pgtoken := "1"
	userList, err := userClient.ListUsers(context.Background(), connect.NewRequest(&api.ListUserRequest{
		PageSize:  int32(pgsize),
		PageToken: pgtoken,
	}))
	if err != nil {
		t.Errorf("ListUsers failed: %v", err)
		t.FailNow()
	}

	nextPage := userList.Msg.NextPageToken
	ok(t, err)
	if len(userList.Msg.Users) != pgsize {
		t.FailNow()
	}

	if nextPage != "page does not exist" {
		t.FailNow()
	}

	//Deleting User
	_, err = userClient.DeleteUser(context.Background(), connect.NewRequest(&api.DeleteUserRequest{
		Id: id,
	}))
	ok(t, err)
	_, err = userClient.GetUser(context.Background(), connect.NewRequest(&api.GetUserRequest{
		Id: id,
	}))
	if err == nil {
		t.Errorf("DeleteUser failed: %v", err)
		t.FailNow()
	}

	tweetClient := apiconnect.NewTweetServiceClient(http.DefaultClient, "http://127.0.0.1:8080")

	//Creating Tweet
	testTweet := &api.Tweet{
		Content: "Test Tweet",
		User:    "1",
	}

	createTweet, err := tweetClient.CreateTweet(context.Background(), connect.NewRequest(&api.CreateTweetRequest{
		Tweet: testTweet,
	}))
	if err != nil {
		t.Errorf("CreateTweet failed: %v", err)
		t.FailNow()
	}
	tid := createTweet.Msg.Id

	notEmpty(t, tid)

	//Getting Tweet
	getTweet, err := tweetClient.GetTweet(context.Background(), connect.NewRequest(&api.GetTweetRequest{
		Id: tid,
	}))
	ok(t, err)
	equalStr(t, getTweet.Msg.Content, testTweet.Content)
	equalStr(t, getTweet.Msg.User, testTweet.User)

	con := "Updated Tweet"
	//Updating Tweet
	updateTweet, err := tweetClient.UpdateTweet(context.Background(), connect.NewRequest(&api.UpdateTweetRequest{
		Id:      tid,
		Content: &con,
	}))
	ok(t, err)
	equalStr(t, updateTweet.Msg.Content, "con")
	equalStr(t, updateTweet.Msg.User, tid)

	//Deleting Tweet
	_, err = tweetClient.DeleteTweet(context.Background(), connect.NewRequest(&api.DeleteTweetRequest{
		Id: tid,
	}))
	ok(t, err)

	_, err = tweetClient.GetTweet(context.Background(), connect.NewRequest(&api.GetTweetRequest{
		Id: tid,
	}))
	if err == nil {
		t.Errorf("DeleteUser failed: %v", err)
		t.FailNow()
	}

}

func ok(t *testing.T, err error) {
	if err != nil {
		t.FailNow()
	}
}

func notEmpty(t *testing.T, v string) {
	if v == "" {
		t.FailNow()
	}
}

func equalStr(t *testing.T, f, s string) {
	if f != s {
		t.FailNow()
	}
}

func str(t *testing.T, v string) string {
	if v == "" {
		t.FailNow()
	}
	return v
}
