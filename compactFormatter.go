package main

import (
	log "github.com/Sirupsen/logrus"
	"bytes"
	"fmt"
	"sort"
	"time"
)

// log formatter
type CompactFormatter struct {
}

func (f *CompactFormatter) Format(entry *log.Entry) ([]byte, error) {
	var keys = make([]string, 0, len(entry.Data))

	b := &bytes.Buffer{}

	b.WriteString(entry.Time.Format(time.RFC3339))
	b.WriteByte(' ')
	b.WriteByte(entry.Level.String()[0])
	b.WriteByte(' ')

	if cid, ok := entry.Data["correlationId"]; ok && len(cid.(string)) > 0 {
		fmt.Fprintf(b, "%7s", (cid.(string))[:7])
	} else {
		b.WriteString("       ")
	}
	b.WriteByte(' ')

	if app, ok := entry.Data["application"]; ok {
		fmt.Fprintf(b, "%5s", app)
	} else {
		b.WriteString("     ")
	}
	b.WriteByte(' ')

	if pltf, ok := entry.Data["platform"]; ok {
		fmt.Fprintf(b, "%5s", pltf)
	} else {
		b.WriteString("     ")
	}
	b.WriteByte(' ')

	length := len(entry.Message)
	if length > 40 {
		length = 40
	}
	fmt.Fprintf(b, "'%-40s' ", entry.Message[:length])

	for k := range entry.Data {
		if k == "application" || k == "correlationId" || k == "platform" || k == "timestamp" {
			continue
		} else {
			keys = append(keys, k)
		}
	}
	sort.Strings(keys)
	for _, k := range keys {
		switch value := entry.Data[k].(type) {
		case string:
			fmt.Fprintf(b, "'%s=%s' ", k, entry.Data[k].(string))
		case error:
			fmt.Fprintf(b, "%q ", value)
		default:
			fmt.Fprintf(b, "'%s=%s' ", k, value)
		}

	}

	b.WriteByte('\n')
	return b.Bytes(), nil
}
