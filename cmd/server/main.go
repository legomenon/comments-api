package main

import (
	"context"
	"fmt"

	"github.com/legomenon/comments-api/internal/db"
)

func Run() error {
	fmt.Println("starting application")

	db, err := db.NewDatabase()
	if err != nil {
		fmt.Println("failed to connect to the database: ", err)
		return err
	}

	if err = db.Ping(context.Background()); err != nil {
		return err
	}
	fmt.Println("Connected to database")
	return nil
}

func main() {
	fmt.Println("go rest api")
	if err := Run(); err != nil {
		fmt.Println(err)
	}
}
