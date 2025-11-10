package logger

import (
	"k8s.io/klog/v2"
)

func Info(err error) {
	klog.Info(err)
}
