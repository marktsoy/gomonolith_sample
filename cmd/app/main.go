package main

import (
	"fmt"

	_ "github.com/lib/pq" // Include postgress driver
	"github.com/marktsoy/gomonolith_sample/internal/app"
)

// init function - called once
// More info: https://tutorialedge.net/golang/the-go-init-function/#:~:text=The%20init%20Function,will%20only%20be%20called%20once.
func init() {
	fmt.Println("Init function. I am called only once. I am called first")
}

func main() {

	app := app.New(app.NewConfig())

	app.Run()
}
