package main

import (
	"poc/post/db"
	"poc/post/router"
)

func main() {
	db.GetPosts()
	router.Serve()
}
