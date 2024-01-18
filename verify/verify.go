package verify

import (
	"net/http"
)

func VerifyLink(link string) bool {
	res, err := http.Get(link)
	if err != nil {
		return false
	}

	if res.StatusCode != 200 {
		return false
	}

	return true

}
