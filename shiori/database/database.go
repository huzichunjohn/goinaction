package database

import (
	"database/sql"
	"shiori/model"

	"github.com/RadhiFadlillah/go-readability"
)

type Database interface {
	SaveBookmark(article readability.Article, tags ...string) (model.Bookmark, error)
	GetBookmarks(indices ...string) ([]model.Bookmark, error)
	DeleteBookmarks(indices ...string) ([]int, []int, error)
	SearchBookmarks(keyword string, tags ...string) ([]model.Bookmark, error)
}

func checkError(err error) {
	if err != nil && err != sql.ErrNoRows {
		panic(err)
	}
}
