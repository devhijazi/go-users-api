package helpers

import (
	"github.com/devhijazi/go-users-api/pkg/errors"
	"gorm.io/gorm"
)

type PaginationOptions struct {
	Search       string
	Page         int
	ItemsPerPage int
	Fag          func(db *gorm.DB) *gorm.DB
}

type PaginationReturnData struct {
	Count        int                      `form:"count" json:"count"`
	Pages        int                      `form:"pages" json:"pages"`
	InPage       int                      `form:"in_page" json:"in_page"`
	ItemsInPage  int                      `form:"items_in_page" json:"items_in_page"`
	ItemsPerPage int                      `form:"items_per_page" json:"items_per_page"`
	Items        []map[string]interface{} `form:"items" json:"items"`
}

func PaginationHelper(db *gorm.DB, table string, options PaginationOptions) (*PaginationReturnData, *errors.Error) {
	search := options.Search
	page := options.Page
	itemsPerPage := options.ItemsPerPage

	if page == 0 {
		page = 1
	}

	switch {
	case itemsPerPage > 100:
		itemsPerPage = 100
	case itemsPerPage <= 0:
		itemsPerPage = 10
	}

	var count int64

	db.Table(table).Count(&count)

	pages := int(int(count) / itemsPerPage)

	var documents []map[string]interface{}

	offset := (page - 1) * itemsPerPage

	documentsQuery := db.Table(table).Offset(offset).Limit(itemsPerPage)

	if search != "" && options.Fag != nil {
		documentsQuery.Scopes(options.Fag)
	}

	documentsQuery.Order("created_at desc").Find(&documents)

	documentsCount := len(documents)

	if documentsCount == 0 {
		documents = []map[string]interface{}{}
	}

	return &PaginationReturnData{
		Count:        int(count),
		Pages:        pages,
		InPage:       page,
		ItemsInPage:  documentsCount,
		ItemsPerPage: itemsPerPage,
		Items:        documents,
	}, nil
}
