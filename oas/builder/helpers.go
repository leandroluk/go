// oas/builder/helpers.go
package builder

import (
	"fmt"
	"strconv"
	"strings"
)

func statusCodeKey(code int) string {
	if code < 100 || code > 599 {
		panic(fmt.Sprintf("invalid HTTP status code: %d", code))
	}
	return strconv.Itoa(code)
}

func statusRangeKey(class int) string {
	if class < 1 || class > 5 {
		panic(fmt.Sprintf("invalid HTTP status class: %d", class))
	}
	return strconv.Itoa(class) + "XX"
}

func defaultResponseKey() string {
	return "default"
}

func normalizePath(value string) string {
	if value == "" {
		return "/"
	}
	if strings.HasPrefix(value, "/") {
		return value
	}
	return "/" + value
}
