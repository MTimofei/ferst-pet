package cookiepkg

import (
	"net/http"
	"time"
)

func CreateCookieAouth(token string) (cookie *http.Cookie) {
	cookie = &http.Cookie{
		Name:     "RefJWT",
		Value:    token,
		Path:     "/",
		Expires:  time.Now().Add(24 * time.Hour),
		HttpOnly: true,
		//Secure: true, перевести в https
	}
	return cookie
}

func CreateCookieAcc(token string, urlclient string) (cookie *http.Cookie) {
	cookie = &http.Cookie{
		Name:     "AccJWT",
		Value:    token,
		Path:     urlclient,
		Expires:  time.Now().Add(24 * time.Hour),
		HttpOnly: true,
		//Secure: true, перевести в https
	}
	return cookie
}
