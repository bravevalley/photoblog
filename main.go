package main

import (
	"context"
	"database/sql"
	"fmt"
	"html/template"
	"log"
	"net/http"
	"os"

	"github.com/go-redis/redis"
	_ "github.com/jackc/pgx/v5/stdlib"
)

var (
	// Templates
	tpl *template.Template

	// MySQL database
	DB *sql.DB

	// Redis Database for sessions
	ctx context.Context
	rdb *redis.Client

	// Error
	err error
)

func init() {

	// Init Template
	tpl = template.Must(template.ParseGlob("templates/*"))

	// Init DB
	DB, err = sql.Open("pgx", os.Getenv("DATABASE_URL"))
	if err != nil {
		fmt.Fprintf(os.Stderr, "Unable to connect to database: %v\n", err)
		os.Exit(1)
	}

	// Init Redis
	rdb = redis.NewClient(&redis.Options{
		Addr:     "localhost:16379",
		Password: "", // no password set
		DB:       0,  // use default DB
	})

}

func main() {

	// Test all periphals
	err = DB.Ping()
	if err != nil {
		log.Fatalln("Cannot connect to Database")
	}

	if err = rdb.Ping().Err(); err != nil {
		log.Fatalln("Cannot connect to Redis Database")
	}

	fmt.Println("Connection secured!")

	http.ListenAndServe(":8080", nil)

}