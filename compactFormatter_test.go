package main

import (
	"testing"
	log "github.com/Sirupsen/logrus"
	"strings"
	"bytes"
)

func TestCompactFromatter(t *testing.T) {
	// given
	newBuffer := bytes.NewBuffer([]byte(""))
	log.SetOutput(newBuffer)
	log.SetFormatter(&CompactFormatter{})

	entry := log.Entry{}
	entry.Logger = log.StandardLogger()

	// test
	entry.WithFields(log.Fields{"a key":"a value"}).Info("my message")

	// check
	logLine := strings.TrimSpace(string(newBuffer.Bytes()))
	expectedLine := "i                     'my message                              ' 'a key=a value'"

	if strings.HasSuffix(logLine, expectedLine) {
		t.Log("success")
	} else {
		t.Errorf("should ends with \n'%s' but is \n'%s' ", expectedLine, logLine)

	}

}
