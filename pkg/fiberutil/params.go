package fiberutil

import (
	"strconv"

	"github.com/gofiber/fiber/v2"
)

// ParamsInt64 get int64 parameters
//
//	return 0 or defaultValue int64 on empty parameter
func ParamsInt64(c *fiber.Ctx, key string, defaultValue ...int64) int64 {
	intStr := c.Params(key)
	if intStr == "" {
		return defaultV(defaultValue)
	}
	paramInt, err := strconv.Atoi(intStr)
	if err != nil {
		return defaultV(defaultValue)
	}
	return int64(paramInt)
}

// ParamsBool get Bool parameters
func ParamsBool(c *fiber.Ctx, key string, defaultValue ...bool) bool {
	boolStr := c.Params(key)
	if boolStr == "" {
		return defaultV(defaultValue)
	}
	switch boolStr {
	case "true", "t", "yes", "y":
		return true
	case "false", "f", "no", "n":
		return false
	default:
		return defaultV(defaultValue)
	}
}

// defaultV return default value of type V if v[] is empty
func defaultV[V any](v []V) V {
	if len(v) > 0 {
		return v[0]
	}
	var defVal V
	return defVal
}
