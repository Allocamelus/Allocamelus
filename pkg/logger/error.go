package logger

import (
	"runtime/debug"

	"k8s.io/klog/v2"
)

// Error logs err if not nil
// Returns true on err
func Error(err error) bool {
	if err != nil {
		klog.Error(err, " : ", string(debug.Stack()))
		return true
	}
	return false
}
