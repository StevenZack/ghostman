package util

import (
	"strings"
)

// Header Header
type Header struct {
	Key, Value string
}

// ParseHeader ParseHeader
func ParseHeader(s string) []Header {
	ss := strings.Split(s, ";")
	header := []Header{}
	for _, h := range ss {
		hs := strings.Split(h, ":")
		if len(hs) < 2 {
			continue
		}
		k := hs[0]
		v := hs[1]
		header = append(header, Header{k, v})
	}
	return header
}
