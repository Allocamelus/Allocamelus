package logger

import (
	"flag"
	"strconv"

	"k8s.io/klog/v2"
)

// InitKlog initializes klog
// Leave filePath empty to use stderr
func InitKlog(v int8, filePath string) {
	klog.InitFlags(nil)
	flag.Set("v", strconv.Itoa(int(v)))
	if filePath != "" {
		// By default klog writes to stderr. Setting logtostderr to false makes klog
		// write to a log file.
		flag.Set("logtostderr", "false")
		flag.Set("log_file", filePath)
	}
	flag.Parse()
}
