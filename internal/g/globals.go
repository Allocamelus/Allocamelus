package g

import (
	"github.com/allocamelus/allocamelus/internal/config"
	"github.com/allocamelus/allocamelus/internal/data"
	"github.com/allocamelus/allocamelus/pkg/fiberutil/session"
)

// Data global
var Data *data.Data

// Session global
var Session *session.Store

// Config global
var Config *config.Config
