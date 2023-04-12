package cookies

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
