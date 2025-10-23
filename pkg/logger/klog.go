package logger

import (
	"flag"
	"strconv"

	"k8s.io/klog/v2"
)

// InitKlog initializes klog
// Leave filePath empty to use stderr
func InitKlog(v int8, isDir bool, path string) {
	klog.InitFlags(nil)
	flag.Set("v", strconv.Itoa(int(v)))
	if path != "" {
		flag.Set("logtostderr", "false")
		flag.Set("alsologtostderr", "true")
		if isDir {
			flag.Set("log_dir", path)
		} else {
			flag.Set("log_file", path)
		}
	} else {
		flag.Set("logtostderr", "true")
	}
	flag.Parse()
}
