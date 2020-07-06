package service

import "github.com/yongwoon/echo-sqlboiler-mysql-blog/model"

// PostList all post lists
func PostList() ([]*model.Post, error) {
	p := &model.Post{}
	return p.All()
}

// PostByID get post by id
func PostByID(id uint64) (*model.Post, error) {
	p := &model.Post{}
	return p.Find(id)
}

// PostCreate Create post
func PostCreate(title, body string) (*model.Post, error) {
	p := &model.Post{}
	return p.Create(title, body)
}
