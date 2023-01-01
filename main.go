package main

import (
	"context"
	"database/sql"
	"fmt"
	"html/template"
	"log"
	"os"

	"github.com/go-redis/redis"
	// "github.com/gorilla/mux"
	_ "github.com/jackc/pgx/v5/stdlib"

	"photo.blog/app"
	"photo.blog/db"
	"photo.blog/session"
)

var (
	// Templates
	tpl *template.Template

	// Redis Database for sessions
	ctx context.Context

	// Error
	err error
)

func init() {

	// Init Template
	tpl = template.Must(template.ParseGlob("assets/templates/*"))

	// Init DB
	db.DB, err = sql.Open("pgx", os.Getenv("DATABASE_URL"))
	if err != nil {
		fmt.Fprintf(os.Stderr, "Unable to connect to database: %v\n", err)
		os.Exit(1)
	}

	// Init Redis
	session.RDB = redis.NewClient(&redis.Options{
		Addr:     "localhost:16379",
		Password: "", // no password set
		DB:       0,  // use default DB
	})

}

func main() {

	// Test all periphals
	err = db.DB.Ping()
	if err != nil {
		log.Fatalln("Cannot connect to Database", err)
	}

	if err = session.RDB.Ping().Err(); err != nil {
		log.Fatalln("Cannot connect to Redis Database")
	}

	fmt.Println("Connection secured!")

	webapp := &app.App{}
	webapp.Init(tpl)
	webapp.Run()
}
