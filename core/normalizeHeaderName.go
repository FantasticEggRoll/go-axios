package core

import (
	"strings"
)

func NormalizeHeaderName(header Header, normalizedName string) {
	for k, v := range header.Header {
		if k != normalizedName && strings.ToLower(k) == strings.ToLower(normalizedName) {
			header.Del(k)
			for _, e := range v {
				header.Add(normalizedName, e)
			}
		}
	}
}
