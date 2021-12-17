package user

import (
	"github.com/allocamelus/allocamelus/internal/data"
	"github.com/allocamelus/allocamelus/internal/user/event"
	"github.com/allocamelus/allocamelus/internal/user/token"
)

// Init prepares user sql
func Init(p data.Prepare) {
	event.Init(p)
	token.Init(p)
	initID(p)
	initUser(p)
	initValidate(p)
}
