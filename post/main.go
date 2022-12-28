package main

import (
	"poc/post/db"
	"poc/post/router"
)

func main() {
	router.Serve()
	db.Migrate()
}
