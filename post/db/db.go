package db

import (
	"log"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"

	"poc/post/models"
)

func Migrate() {
	db := dbConn()

	err := db.AutoMigrate(&models.Post{})
	if err != nil {
		log.Fatal(err)
	}
}

func GetPostByID(id int) (models.Post, error) {
	dbConn := dbConn()
	var post models.Post

	dbConn.First(&post, id)

	return post, nil
}

func GetPosts() ([]models.Post, error) {
	dbConn := dbConn()
	var posts []models.Post

	dbConn.Find(&posts)

	return posts, nil
}

func CreatePost(body models.RequestCreatePost) error {
	dbConn := dbConn()

	post := models.Post{Author: "John Mayor", Title: "Hello, world", Content: "Yoyo"}

	dbConn.Select("Author", "Title", "Content").Create(&post)

	return nil
}

func dbConn() (db *gorm.DB) {
	dsn := "host=post_db user=postgres password=postgres dbname=postgres port=5432 sslmode=disable TimeZone=Asia/Shanghai"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal(err)
	}

	return db
}
