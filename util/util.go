package util

import (
	"crypto/rand"
	"net/url"
	"strings"
)

// generateAlphanumericToken generates a code with uppercase/lowercase/numbers with specified length
func generateAlphanumericToken(size int) string {
	const Master = "qwertyuiopasdfghjklzxcvbnmQWERTYUIOPASDFGHJKLZXCVBNM1234567890" // the token chars must be one of these
	var b strings.Builder
	b.Grow(size)
	bytes := make([]byte, 1)
	for i := 0; i < size; i++ {
		for {
			_, _ = rand.Read(bytes)
			bytes[0] &= 63 // 0011 1111
			if int(bytes[0]) < len(Master) {
				b.WriteByte(Master[int(bytes[0])])
				break
			}
		}
	}
	return b.String()
}

// GenerateLinkToken generates the link token
func GenerateLinkToken() string {
	return generateAlphanumericToken(7)
}

// IsUrlValid checks if a string is a url
// From https://stackoverflow.com/a/55551215/4213397
func IsUrlValid(link string) bool {
	u, err := url.Parse(link)
	return err == nil && u.Scheme != "" && u.Host != ""
}