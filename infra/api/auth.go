package api

import "net/http"

func SetAuthHeader(req *http.Request, token string) {
	req.Header.Set("Authorization", "Bearer "+token)
}
