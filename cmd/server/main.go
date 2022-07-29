package main

import (
	"errors"
	"fmt"

	"github.com/golang-migrate/migrate/v4"
	"github.com/legomenon/comments-api/internal/comment"
	"github.com/legomenon/comments-api/internal/db"
	transportHttp "github.com/legomenon/comments-api/internal/transport/http"
)

func Run() error {
	fmt.Println("starting application")

	db, err := db.NewDatabase()
	if err != nil {
		fmt.Println("failed to connect to the database: ", err)
		return err
	}

	if err = db.MigrateDB(); err != nil {
		if !errors.Is(err, migrate.ErrNoChange) {
			fmt.Println("failed to migrate database")
			return err
		}
	}
	fmt.Println("Connected to database")

	cmtService := comment.NewService(db)

	httpHandler := transportHttp.NewHandler(cmtService)
	if err = httpHandler.Serve(); err != nil {
		return err
	}

	return nil
}

func main() {
	fmt.Println("go rest api")
	if err := Run(); err != nil {
		fmt.Println(err)
	}
}
