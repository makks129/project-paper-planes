package utils

import "net/http"

func GetCookie(name string, r *http.Request) string {
	cookies := r.Cookies()
	for _, cookie := range cookies {
		if cookie.Name == name {
			return cookie.Value
		}
	}
	return ""
}
