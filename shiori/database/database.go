package database

import (
	"shiori/model"

	"github.com/RadhiFadlillah/go-readability"
)

type Database interface {
	SaveBookmark(article readability.Article, tags ...string) (model.Bookmark, error)
	GetBookmarks(indices ...string) ([]model.Bookmark, error)
}

func checkError(err error) {
	if err != nil {
		panic(err)
	}
}
