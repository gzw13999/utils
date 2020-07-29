package utils

import (
	"bytes"
	"compress/gzip"
)

func Gzip(b []byte,level int) ([]byte, error) {
	var buf bytes.Buffer
	if gzw, err := gzip.NewWriterLevel(&buf, level); err != nil {
		return nil, err
	} else {
		if _, err := gzw.Write(b); err != nil {
			return nil, err
		} else {
			if err := gzw.Close(); err != nil {
				return nil, err
			} else {
				return buf.Bytes(), nil
			}
		}
	}
}
