package util

import (
	"log"
	"time"

	"k8s.io/apimachinery/pkg/util/wait"
	"k8s.io/klog"
)

type KlogWriter struct{}

func (writer KlogWriter) Write(data []byte) (n int, err error) {
	klog.Info(string(data))
	return len(data), nil
}

func InitLogs() {
	log.SetOutput(KlogWriter{})
	log.SetFlags(log.LUTC | log.Ldate | log.Ltime)
	flushFreq := 5 * time.Second
	go wait.Until(klog.Flush, flushFreq, wait.NeverStop)
}

func FlushLogs() {
	klog.Flush()
}
