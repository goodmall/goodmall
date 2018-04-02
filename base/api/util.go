package api

import (
	"strconv"

	"net/url"

	"github.com/goodmall/goodmall/base/util"
)

const (
	DEFAULT_PAGE_SIZE int = 20 // 100
	MAX_PAGE_SIZE     int = 1000
)

func GetPaginatedListFromRequest(c *url.URL, count int) *util.PaginatedList {
	page := parseInt(c.Query().Get("page"), 1)
	perPage := parseInt(c.Query().Get("per_page"), DEFAULT_PAGE_SIZE)
	if perPage <= 0 {
		perPage = DEFAULT_PAGE_SIZE
	}
	if perPage > MAX_PAGE_SIZE {
		perPage = MAX_PAGE_SIZE
	}
	return util.NewPaginatedList(page, perPage, count)
}

func parseInt(value string, defaultValue int) int {
	if value == "" {
		return defaultValue
	}
	if result, err := strconv.Atoi(value); err == nil {
		return result
	}
	return defaultValue
}
