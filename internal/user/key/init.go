package key

import "github.com/allocamelus/allocamelus/internal/data"

// Init prepared statements
func Init(p data.Prepare) {
	initKeys(p)
}
