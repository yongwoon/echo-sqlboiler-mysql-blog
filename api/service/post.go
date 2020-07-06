package service

import "github.com/yongwoon/echo-blog/model"

// PostList all post lists
func PostList() ([]*model.Post, error) {
	p := &model.Post{}
	return p.All()
}

// PostList all post lists
func PostById(id uint64) (*model.Post, error) {
	p := &model.Post{}
	return p.Find(id)
}
