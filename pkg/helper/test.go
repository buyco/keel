package helper

import (
	"bytes"
	"os"

	log "github.com/sirupsen/logrus"
)

// CaptureStdout is used to get what is logged
func CaptureStdout(f func()) string {
	var buf bytes.Buffer
	log.SetOutput(&buf)
	f()
	log.SetOutput(os.Stdout)
	return buf.String()
}
