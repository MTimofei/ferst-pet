package cookie

import (
	"crypto/rsa"
	"log"
	"net/http"
	"pet/integration_verification/real-time/grpcclient/cookie/jwtpkg"
	"time"

	"github.com/golang-jwt/jwt"
)

func createCookieClientServer(urlclient *string) (cookie *http.Cookie) {
	cookie = &http.Cookie{
		Name:     "Client",
		Value:    *urlclient,
		Path:     "http://localhost:8889/",
		Expires:  time.Now().Add(5 * time.Second),
		HttpOnly: true,
		//Secure: true, перевести в https
	}
	return cookie
}

func HandlingCookiesClients(urlserves *string, key *rsa.PublicKey, w http.ResponseWriter, r *http.Request) (token *jwt.Token, err error) {
	cookieacc, err := r.Cookie("AccJWT")
	if err != nil {
		if err.Error() != http.ErrNoCookie.Error() {
			return nil, err
		}
		log.Println(err)
		cookie := createCookieClientServer(urlserves)
		http.SetCookie(w, cookie)
		http.Redirect(w, r, "http://localhost:8889/", http.StatusSeeOther)
		return
	}
	token, err = jwtpkg.VerificationJWTAcc(cookieacc.Value, key)
	if err != nil {
		log.Println(err)
		cookie := createCookieClientServer(urlserves)
		http.SetCookie(w, cookie)
		http.Redirect(w, r, "http://localhost:8889/", http.StatusSeeOther)
		return
	}
	return token, nil
}
