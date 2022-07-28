package main

import (
	"context"
	"errors"
	"fmt"

	"github.com/golang-migrate/migrate/v4"
	"github.com/legomenon/comments-api/internal/comment"
	"github.com/legomenon/comments-api/internal/db"
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

	cmtService.PostComment(
		context.Background(),
		comment.Comment{
			ID:     "0c7e49f2-9238-4306-a401-4add08aa574d",
			Slug:   "Manual test",
			Body:   "hello world",
			Author: "Legomenon",
		},
	)

	fmt.Println(cmtService.GetComment(
		context.Background(),
		"0c7e49f2-9238-4306-a401-4add08aa574d",
	))
	return nil
}

func main() {
	fmt.Println("go rest api")
	if err := Run(); err != nil {
		fmt.Println(err)
	}
}
