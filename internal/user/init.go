package user

import "github.com/allocamelus/allocamelus/internal/data"

// Init prepares user sql
func Init(p data.Prepare) {
	initCreate(p)
	initValidate(p)
}
