package logger

import (
	"runtime/debug"

	"k8s.io/klog/v2"
)

// Fatal logs err if not nil
// Returns true on err
func Fatal(err error) {
	if err != nil {
		klog.Fatal(err, " : ", string(debug.Stack()))
	}
}
