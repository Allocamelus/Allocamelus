package event

import "github.com/allocamelus/allocamelus/internal/data"

// Init prepared statements
func Init(p data.Prepare) {
	initEvents(p)
}
