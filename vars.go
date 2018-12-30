package main

import "os"

var (
	// Port port to listen
	Port = os.Getenv("PORT")

	// DatabaseDSN database dsn
	DatabaseDSN = os.Getenv("DATABASE_DSN")
)

func init() {
	if len(Port) == 0 {
		Port = "3000"
	}
}
