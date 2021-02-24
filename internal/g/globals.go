package g

import (
	"github.com/allocamelus/allocamelus/configs"
	"github.com/allocamelus/allocamelus/internal/data"
	"github.com/jdinabox/goutils/fiber/session"
)

// Data global
var Data *data.Data

// Session global
var Session *session.Store

// Config global
var Config *configs.Config
