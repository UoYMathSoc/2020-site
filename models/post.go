package models

import (
	"os"

	"github.com/UoYMathSoc/2020-site/database"
	"golang.org/x/tools/present"
)

type Post struct {
	*present.Doc
	Key          string
	Related      []*Post
	Newer, Older *Post
}

type PostStore struct {
	querier database.Querier
}

func NewPostStore(q database.Querier) PostStore {
	return PostStore{q}
}

// For now always returns my first post
func (ps *PostStore) Get(key string) (*Post, error) {
	path := "posts/my-first-post.md" // "posts/" + key + fileExt

	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	doc, err := present.Parse(file, path, 0)

	post := Post{
		Doc:     doc,
		Key:     "my-first-post",
		Related: nil,
		Newer:   nil,
		Older:   nil,
	}
	return &post, nil
}
