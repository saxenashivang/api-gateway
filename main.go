package main

import (
	"github.com/saxenashivang/api-gateway/bootstrap"
)

// main : entry point
func main() {
	// logger := lib.GetLogger()
	bootstrap.GetApp().Run()
}
