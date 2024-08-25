package models

import (
	"database/sql"
)

type Post struct {
	ID      int
	Title   string
	Content string
}

func GetAllPosts(db *sql.DB) ([]Post, error) {
	rows, err := db.Query("SELECT id, title, content FROM posts")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var posts []Post
	for rows.Next() {
		var post Post
		if err := rows.Scan(&post.ID, &post.Title, &post.Content); err != nil {
			return nil, err
		}
		posts = append(posts, post)
	}
	return posts, nil
}

func CreatePost(db *sql.DB, title, content string) error {
	_, err := db.Exec("INSERT INTO posts (title, content) VALUES (?, ?)", title, content)
	return err
}

func UpdatePost(db *sql.DB, id, title, content string) error {
	_, err := db.Exec("UPDATE posts SET title = ?, content = ? WHERE id = ?", title, content, id)
	return err
}

func GetPostById(db *sql.DB, id string) (*Post, error) {
	rows, err := db.Query("SELECT id, title, content FROM posts WHERE id = ?", id)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var post Post
	for rows.Next() {
		if err := rows.Scan(&post.ID, &post.Title, &post.Content); err != nil {
			return nil, err
		} else {
			break
		}
	}

	return &post, nil
}
