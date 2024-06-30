package auth

import (
	"errors"
	"net/http"
	"strings"
)

// GetAPIKey extracts on API Key from
// the headers of an HTTP request
// Example: Authorization: ApiKey{insert apikey here}

func GetAPIKey(headers http.Header) (string, error) {
	val := headers.Get("Authorization")
	if val == "" {
		return "", errors.New("no authenitcation info found")
	}

	vals := strings.Split(val, " ") // here the space is acting as delimiter
	if len(vals) != 2 {
		return "", errors.New("malformed auth header")
	}
	if vals[0] != "ApiKey" { //checks if the first part of auth header is ApiKey or not
		return "", errors.New("malformed first part of auth header")
	}
	return vals[1], nil
}
