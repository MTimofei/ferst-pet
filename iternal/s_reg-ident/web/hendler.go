package web

import (
	"fmt"
	"log"
	"net/http"
	"pet/iternal/s_reg-ident/cookiepkg"
	"pet/iternal/s_reg-ident/db/comsql"
	"pet/iternal/s_reg-ident/str/authe"
	"pet/iternal/s_reg-ident/str/regin"
	"pet/iternal/s_reg-ident/str/salt"
	"pet/pkg/myerr"
	"pet/pkg/pars"
)

func (con *Connect) handlerMain(w http.ResponseWriter, r *http.Request) {
	//r.UserAgent()
	//log.Println(r.Header)
	cookieref, err := r.Cookie("RefJWT")
	if err != nil {
		myerr.ServesError(w, con.HashTempl, err)
		return
	}

	a, err := authe.AuthRefJWT(con.KRef, cookieref.Value)
	if err != nil {
		myerr.ServesError(w, con.HashTempl, err)
		return
	}

	if cookiecli, err := r.Cookie("Client"); err != nil {
		err = nil
	} else {
		log.Println("client cookei", cookiecli.Value)
		token, err := con.KAcc.CreateJWTAcc(a.Authdata)
		if err != nil {
			log.Println(err)
			return
		}
		cookieacc := cookiepkg.CreateCookieAcc(token, cookiecli.Value)
		log.Println("acc cookei", cookieacc.Value)
		http.SetCookie(w, cookieacc)
		http.Redirect(w, r, cookiecli.Value, http.StatusSeeOther)
	}

	pars.ParsPage(w, "hi", con.HashTempl, *a)
}

func (con *Connect) handlerRegPage(w http.ResponseWriter, r *http.Request) {
	log.Println("connection reg")

	pars.ParsPage(w, "reg", con.HashTempl, nil)

}

func (con *Connect) handlerRegProcess(w http.ResponseWriter, r *http.Request) {
	log.Println("process reg")

	if len(r.FormValue("name")) == 0 || len(r.FormValue("password")) == 0 || len(r.FormValue("email")) == 0 {
		err := fmt.Errorf("not se value")
		myerr.ServesError(w, con.HashTempl, err)
		return
	}

	err := comsql.CheckUinquenessLogin(con.MySQL, r)
	if err != nil {
		myerr.ServesError(w, con.HashTempl, err)
		return
	}

	err = comsql.CheckUinquenessEmail(con.MySQL, r)
	if err != nil {
		myerr.ServesError(w, con.HashTempl, err)
		return
	}

	rd, err := regin.New(r)
	if err != nil {
		myerr.ServesError(w, con.HashTempl, err)
		return
	}

	salt := salt.GenerateSalt()
	key := salt.GeneraterKey(rd.GetPass())

	err = comsql.SendRegData(con.MySQL, salt, rd, key)
	if err != nil {
		myerr.ServesError(w, con.HashTempl, err)
		return
	}

	pars.ParsPage(w, "regstat", con.HashTempl, http.StatusOK)

}

func (con *Connect) handlerAuthPage(w http.ResponseWriter, r *http.Request) {
	log.Println("connection auth")

	pars.ParsPage(w, "auth", con.HashTempl, nil)

}

func (con *Connect) handlerAuthProcess(w http.ResponseWriter, r *http.Request) {
	log.Println("process auth")

	if r.Method != "POST" {
		err := fmt.Errorf("method not post")
		myerr.ServesError(w, con.HashTempl, err)
		return
	}
	if len(r.FormValue("name")) == 0 || len(r.FormValue("password")) == 0 {
		err := fmt.Errorf("not se value")
		myerr.ServesError(w, con.HashTempl, err)
		return
	}

	resolt, err := comsql.GetAccountData(con.MySQL, r.FormValue("name"))
	if err != nil {
		myerr.ServesError(w, con.HashTempl, err)
		return
	}

	a := authe.New(resolt, r)
	a.Compare()
	token, err := a.CreateJWTRef(con.KRef)
	if err != nil {
		myerr.ServesError(w, con.HashTempl, err)
		return
	}

	http.SetCookie(w, cookiepkg.CreateCookieAouth(token))

	pars.ParsPage(w, "regstat", con.HashTempl, http.StatusOK)
}
