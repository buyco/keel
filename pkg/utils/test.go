package utils

import (
	"bytes"
	log "github.com/sirupsen/logrus"
	"os"
)

func CaptureStdout(f func()) string {
	var buf bytes.Buffer
	log.SetOutput(&buf)
	f()
	log.SetOutput(os.Stdout)
	return buf.String()
}
