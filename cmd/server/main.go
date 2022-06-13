package main

import (
	"fmt"
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

	return nil
}

func main() {
	fmt.Println("I'm aliveeee!")
	if err := Run(); err != nil {
		fmt.Println(err)
	}
}
