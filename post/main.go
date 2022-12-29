package main

import (
	"poc/post/db"
	"poc/post/router"
)

func main() {
	db.TestInsert()
	db.GetPosts()
	router.Serve()
}
