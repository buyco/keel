package utils

import (
	"bytes"
	log "github.com/sirupsen/logrus"
	"os"
)

// CaptureStdout is used to get what is logged
func CaptureStdout(f func()) string {
	var buf bytes.Buffer
	log.SetOutput(&buf)
	f()
	log.SetOutput(os.Stdout)
	return buf.String()
}
