package fiberutil

import (
	"strconv"

	"github.com/gofiber/fiber/v2"
)

// ParamsInt64 get int64 parameters
// 	returns 0 or defaultValue int64 on empty parameter
func ParamsInt64(c *fiber.Ctx, key string, defaultValue ...int64) int64 {
	intStr := c.Params("id")
	if intStr == "" {
		return defInt64(defaultValue)
	}
	paramInt, err := strconv.Atoi(intStr)
	if err != nil {
		return defInt64(defaultValue)
	}
	return int64(paramInt)
}

func defInt64(defaultValue []int64) int64 {
	if len(defaultValue) > 0 {
		return defaultValue[0]
	}
	return 0
}
