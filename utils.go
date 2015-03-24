// Package main provides ...
package main

import "time"

// convert ISO8601 to RFC 2822
// RFC 2822 references RFC 1132 and obsoletes RFC 822
func Iso8601ToRfc1123(in string) string {
	iso8601 := "2006-01-02T15:04:05.000Z"
	t, err := time.Parse(iso8601, in)
	if err != nil {
		panic(err)
	}
	return t.Format(time.RFC1123Z)
}
