package user

import (
	"github.com/allocamelus/allocamelus/internal/data"
	"github.com/allocamelus/allocamelus/internal/user/token"
)

// Init prepares user sql
func Init(p data.Prepare) {
	token.Init(p)
	initCreate(p)
	initValidate(p)
}
