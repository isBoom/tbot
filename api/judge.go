package api

import (
	"strings"
)

func Judge(str string,key []string) string {
	for _, k := range key {
		if strings.Index(str,k) != -1 {
			return k
		}
	}
	return ""
}
