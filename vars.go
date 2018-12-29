package main

import "os"

var (
	// Port port to listen
	Port = os.Getenv("PORT")
)

func init() {
	if len(Port) == 0 {
		Port = "3000"
	}
}
