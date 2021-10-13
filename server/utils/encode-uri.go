package utils

import (
	"net/url"
	"strings"
)

// EncodeURIComponent
//
// via https://blog.csdn.net/zengming00/article/details/81977362
func EncodeURIComponent(str string) string {
	r := url.QueryEscape(str)
	r = strings.Replace(r, "+", "%20", -1)
	r = strings.Replace(r, "*", "%2A", -1)
	return r
}
