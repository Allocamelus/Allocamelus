package token

import "github.com/allocamelus/allocamelus/internal/data"

// Init prepared statements
func Init(p data.Prepare) {
	initCheck(p)
	initToken(p)
}
