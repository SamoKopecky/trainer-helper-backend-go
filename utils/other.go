package utils

import (
	"log"
	"net/url"
	"unicode"
)

func UpperFirstChar(s string) string {
	r := []rune(s)
	r[0] = unicode.ToUpper(r[0])
	return string(r)
}

func AddQueryParam(rawUrl string, params map[string]string) string {
	u, err := url.Parse(rawUrl)
	if err != nil {
		log.Fatal(err)
	}

	q := u.Query()
	for k, v := range params {
		q.Set(k, v)
	}
	u.RawQuery = q.Encode()
	return u.String()
}
