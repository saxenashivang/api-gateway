package main

import (
	"github.com/saxenashivang/api-gateway/bootstrap"
)

// main : entry point
func main() {
	//logger initialize before app starts because in provider we need logger
	// Load conigurations
	// if err := config.Load(); err != nil {
	// 	logger.Fatal(err)
	// }
	bootstrap.GetApp().Run()
}
