// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.22.0

package sqlc

import ()

type Tweet struct {
	ID      int32
	Content string
	User    int32
}

type User struct {
	ID        int32
	FirstName string
	LastName  string
}
