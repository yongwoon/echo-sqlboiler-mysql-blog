package model

import (
	"context"

	_ "github.com/go-sql-driver/mysql"

	"github.com/volatiletech/sqlboiler/v4/boil"
	"github.com/volatiletech/sqlboiler/v4/queries/qm"
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

// Find return all posts
func (p *Post) Find(id uint64) (*Post, error) {
	db, _ := db.Connect()

	res, err := orm.FindPost(context.Background(), db, id)
	if err != nil {
		return nil, err
	}

	post := &Post{ID: res.ID, Title: res.Title, Body: res.Body}

	return post, nil
}

// Create create post
// TODO transaction
func (p *Post) Create(title, body string) (*Post, error) {
	db, _ := db.Connect()

	latestPost, err := orm.Posts(qm.Select("MAX(id) as id")).One(context.Background(), db)
	if err != nil {
		return nil, err
	}

	ormPost := orm.Post{
		ID:    latestPost.ID + 1,
		Title: title,
		Body:  body,
	}

	err = ormPost.Insert(context.Background(), db, boil.Infer())
	if err != nil {
		return nil, err
	}

	post := &Post{ID: ormPost.ID, Title: ormPost.Title, Body: ormPost.Body}

	return post, nil
}
