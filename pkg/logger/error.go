package logger

import (
	"runtime/debug"

	"k8s.io/klog/v2"
)

// Error logs err if not nil
// Returns true on err
func Error(err error) bool {
	if err == nil {
		return false
	}
	// Verbose logging
	if klog.V(2).Enabled() {
		klog.Error(err, " : ", string(debug.Stack()))
		return true
	}
	klog.Error(err)
	return true
}
