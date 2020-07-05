package model

import (
	"context"

	_ "github.com/go-sql-driver/mysql"

	db "github.com/yongwoon/echo-blog/db"
	"github.com/yongwoon/echo-blog/db/orm"
)

type (
	// IPost interface
	IPost interface {
	}

	// Post struct
	Post struct {
		ID    uint64 `json:"id"`
		Title string `json:"title"`
		Body  string `json:"body"`
	}
)

// All return all posts
func (p *Post) All() ([]*Post, error) {
	db, _ := db.Connect()

	posts, err := orm.Posts().All(context.Background(), db)
	if err != nil {
		return nil, err
	}

	list := make([]*Post, len(posts))
	for i := 0; i < len(posts); i++ {
		post := posts[i]
		list[i] = &Post{post.ID, post.Title, post.Body}
	}

	return list, nil
}
