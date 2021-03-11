package post

import "github.com/allocamelus/allocamelus/internal/data"

// Init prepares sql
func Init(p data.Prepare) {
	initPost(p)
}
