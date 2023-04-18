package cookie

import (
	"net/http"
	"time"
)

func CreateCookieClient(urlclient string) (cookie *http.Cookie) {
	cookie = &http.Cookie{
		Name:     "Client",
		Value:    urlclient,
		Path:     "http://localhost:8889/",
		Expires:  time.Now().Add(5 * time.Second),
		HttpOnly: true,
		//Secure: true, перевести в https
	}
	return cookie
}
