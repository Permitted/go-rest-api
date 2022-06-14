package main

import (
	"context"
	"fmt"
	"github.com/Permitted/go-rest-api/internal/comment"
	"github.com/Permitted/go-rest-api/internal/database"
)

// Run - is going to responsible for
// the instantiation and start up of
// go application
func Run() error {
	fmt.Println("Starting application.")

	db, err := database.NewDatabase()
	if err != nil {
		fmt.Println("Failed to connect to the database")
		return err
	}
	if err := db.MigrateDB(); err != nil {
		fmt.Println("failed to migrate database")
		return err
	}

	cmtService := comment.NewService(db)

	cmtService.PostComment(
		context.Background(),
		comment.Comment{
			ID:     "71c5d074-b6cf-11ec-b909-0242ac120002",
			Slug:   "manuel-test",
			Author: "root",
			Body:   "testing value",
		})

	fmt.Println(cmtService.GetComment(
		context.Background(),
		"71c5d074-b6cf-11ec-b909-0242ac120002",
	))

	return nil
}

func main() {
	fmt.Println("I'm aliveeee!")
	if err := Run(); err != nil {
		fmt.Println(err)
	}
}
