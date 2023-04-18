package web

import (
	"log"
	"net/http"
	"pet/iternal/s_main/cookie"
	"pet/pkg/myerr"
	"pet/pkg/pars"
)

func (con *Connect) hendlerMain(w http.ResponseWriter, r *http.Request) {
	cookieacc, err := r.Cookie("AccJWT")
	if err != nil {
		if err.Error() != http.ErrNoCookie.Error() {
			myerr.ServesError(w, con.PageHash, err)
			return
		}
		cookie := cookie.CreateCookieClient("http://localhost:8888/main")
		http.SetCookie(w, cookie)
		http.Redirect(w, r, "http://localhost:8889/", http.StatusSeeOther)
		return
	}
	log.Println(*cookieacc)

	err = pars.ParsPage(w, "hi", con.PageHash, nil)
	if err != nil {
		myerr.ServesError(w, con.PageHash, err)
	}
}
