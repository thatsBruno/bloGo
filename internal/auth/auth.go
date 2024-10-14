package auth

import (
	"errors"
	"net/http"
	"strings"
)

func GetAPIKey(headers http.Header) (string, error) {
	val := headers.Get("Authorization")

	if val == "" {
		return "", errors.New("no Auth information")
	}

	vals := strings.Split(val, " ")
	if len(vals) != 2 {
		return "", errors.New("malformed Auth")
	}

	if vals[0] != "ApiKey" {
		return "", errors.New("malformed Auth")
	}

	return vals[1], nil
}
