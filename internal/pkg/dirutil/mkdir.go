package dirutil

import "os"

// MakeDir wrapper for os.MkdirAll
func MakeDir(path string) error {
	return os.MkdirAll(path, os.ModeSticky|os.ModePerm)
}
