package web

import (
	"log"
	"net/http"
	"pet/iternal/s_main/cookie"
	"pet/iternal/s_main/jwt"
	"pet/iternal/s_main/str"
	"pet/pkg/myerr"
	"pet/pkg/pars"
)

func (con *Connect) hendlerMain(w http.ResponseWriter, r *http.Request) {
	cookieacc, err := r.Cookie("AccJWT")
	log.Println(cookieacc.Value)
	if err != nil {
		if err.Error() != http.ErrNoCookie.Error() {
			myerr.ServesError(w, con.PageHash, err)
			return
		}
		log.Println(err)
		cookie := cookie.CreateCookieClient("http://localhost:8888/main")
		http.SetCookie(w, cookie)
		http.Redirect(w, r, "http://localhost:8889/", http.StatusSeeOther)
		return
	}
	token, err := jwt.VerificationJWTAcc(cookieacc.Value, con.PublicKey)
	if err != nil {
		log.Println(err)
		cookie := cookie.CreateCookieClient("http://localhost:8888/main")
		http.SetCookie(w, cookie)
		http.Redirect(w, r, "http://localhost:8889/", http.StatusSeeOther)
		return
		// myerr.ServesError(w, con.PageHash, err)
		// return
	}

	err = pars.ParsPage(w, "hi", con.PageHash, str.NewAccountFromJWT(token))
	if err != nil {
		myerr.ServesError(w, con.PageHash, err)
		return
	}
}
