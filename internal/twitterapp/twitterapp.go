package twitterapp

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"

	"github.com/codesmith-dev/twitter/internal/gen/api/apiconnect"
	"github.com/codesmith-dev/twitter/internal/services"
	_ "github.com/jackc/pgx/v5/stdlib"
	_ "github.com/lib/pq"
)

var db *sql.DB

const (
	host     = "localhost"
	port     = ":5432"
	username = "postgres"
	password = "postgre123"
	dbname   = "twitter"
)

// Run runs the twitter app.
func Run() {
	var err error
	connStr := "postgres://postgres:postgre123@localhost:5432/twitter?sslmode=disable"
	db, err = sql.Open("pgx", connStr)
	if err != nil {
		fmt.Println("Error occuered while connecting to data base. Err :", err)
	}
	fmt.Println("DataBase connected")
	defer db.Close()

	http.Handle(apiconnect.NewUserServiceHandler(
		// services.NewUserServiceHandler(nil),
		services.NewUserServiceHandler(db),
	))
	http.Handle(apiconnect.NewTweetServiceHandler(
		services.NewTweetServiceHandler(db),
	))
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatalln(err)
	}
}
