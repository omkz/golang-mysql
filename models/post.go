package models

import (
	"github.com/omkz/golang-mysql/config"
)

type Post struct {
	Id      string
	Title   string
	Content string
}

func PostAll() ([]*Post, error) {

	rows, err := config.DB.Query("select * from posts")

	if err != nil {
		return nil, err
	}

	defer rows.Close()

	posts := []*Post{}

	for rows.Next() {
		post := &Post{}
		err := rows.Scan(&post.Id, &post.Title, &post.Content)
		if err != nil {
			return nil, err
		}
		posts = append(posts, post)
	}
	if err = rows.Err(); err != nil {
		return nil, err
	}

	return posts, nil
}
