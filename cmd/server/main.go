package main

import "fmt"

// Run - is going to responsible for
// the instantiation and start up of
// go application
func Run() error {
	fmt.Println("Starting application.")
	return nil
}

func main() {
	fmt.Println("I'm aliveeee!")
	if err := Run(); err != nil {
		fmt.Println(err)
	}
}
