package main

import (
	"database/sql"

	_ "github.com/lib/pq"
)

type Post struct {
	Id      int    `json:"id"`
	Title   string `json:"title"`
	Content string `json:"content"`
	Author  string `json:"author"`
}

var Db *sql.DB

// connect to the Db
func init() {
	var err error
	Db, err = sql.Open("postgres", "user=admin dbname=hello_web password=admin123 sslmode=disable")
	if err != nil {
		panic(err)
	}
}

// Create a new post
func (post *Post) Create() (err error) {
	statement := "INSERT INTO posts (title,content, author) VALUES ($1, $2, $3) RETURNING id"
	stmt, err := Db.Prepare(statement)
	if err != nil {
		return
	}
	defer stmt.Close()
	err = stmt.QueryRow(post.Title, post.Content, post.Author).Scan(&post.Id)
	return
}

// Update a post
func (post *Post) Update() (err error) {
	_, err = Db.Exec("UPDATE posts SET content = $2, author = $3, title = $4 WHERE id = $1", post.Id, post.Content, post.Author, post.Title)
	return
}

// Delete a post
func (post *Post) Delete() (err error) {
	_, err = Db.Exec("DELETE FROM posts WHERE id = $1", post.Id)
	return
}

// get all posts
func GetAllPosts() (posts []Post, err error) {
	rows, err := Db.Query("SELECT id, title, content, author FROM posts")
	if err != nil {
		return
	}
	for rows.Next() {
		post := Post{}
		err = rows.Scan(&post.Id, &post.Title, &post.Content, &post.Author)
		if err != nil {
			return
		}
		posts = append(posts, post)
	}
	rows.Close()
	return
}

// Get a single post
func GetPost(id int) (post Post, err error) {
	post = Post{}
	err = Db.QueryRow("SELECT id, title, content, author FROM posts WHERE id = $1", id).Scan(&post.Id, &post.Title, &post.Content, &post.Author)
	return
}
