package cache

import (
	"strconv"
	"strings"

	"github.com/boybird/hello/pkg/e"
)

// Tag tag
type Tag struct {
	ID       int
	Name     string
	State    int
	PageNum  int
	PageSize int
}

// GetTagsKey get tags
func (t Tag) GetTagsKey() string {
	keys := []string{
		e.CACHE_TAG,
		"LIST",
	}

	if t.State >= 0 {
		keys = append(keys, strconv.Itoa(t.State))
	}
	if t.PageNum >= 0 {
		keys = append(keys, strconv.Itoa(t.PageNum))
	}

	if t.PageSize >= 0 {
		keys = append(keys, strconv.Itoa(t.PageSize))
	}

	return strings.Join(keys, "_")

}
