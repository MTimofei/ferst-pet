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
	cookieref, err := r.Cookie("RefJWT")
	if err != nil {
		myerr.ServesError(w, con.KeshTempl, err)
		return
	}

	userverific, err := authe.AuthRefJWT(con.KeyRef, cookieref.Value)
	if err != nil {
		myerr.ServesError(w, con.KeshTempl, err)
		return
	}

	if cookiecli, err := r.Cookie("Client"); err != nil {
		err = nil
	} else {
		log.Println("client cookei", cookiecli.Value)
		token, err := con.KeyAcc.CreateJWTAcc(userverific.Authdata)
		if err != nil {
			log.Println(err)
			return
		}
		cookieacc := cookiepkg.CreateCookieAcc(token, cookiecli.Value)
		log.Println("acc cookei", cookieacc.Value)
		http.SetCookie(w, cookieacc)
		http.Redirect(w, r, cookiecli.Value, http.StatusSeeOther)
	}

	pars.ParsPage(w, "hi", con.KeshTempl, userverific.Authdata.Logname)
}

func (con *Connect) handlerRegPage(w http.ResponseWriter, r *http.Request) {
	log.Println("connection reg")
	pars.ParsPage(w, "reg", con.KeshTempl, nil)

}

func (con *Connect) handlerRegProcess(w http.ResponseWriter, r *http.Request) {
	log.Println("process reg")

	if len(r.FormValue("name")) == 0 || len(r.FormValue("password")) == 0 || len(r.FormValue("email")) == 0 {
		err := fmt.Errorf("not se value")
		myerr.ServesError(w, con.KeshTempl, err)
		return
	}

	err := comsql.CheckUinquenessLogin(con.MySQL, r)
	if err != nil {
		myerr.ServesError(w, con.KeshTempl, err)
		return
	}

	err = comsql.CheckUinquenessEmail(con.MySQL, r)
	if err != nil {
		myerr.ServesError(w, con.KeshTempl, err)
		return
	}

	rd, err := regin.New(r)
	if err != nil {
		myerr.ServesError(w, con.KeshTempl, err)
		return
	}

	salt := salt.GenerateSalt()
	hashpasswopd := salt.GeneraterHashPassword(rd.GetPass())

	err = comsql.SendRegData(con.MySQL, salt, rd, hashpasswopd)
	if err != nil {
		myerr.ServesError(w, con.KeshTempl, err)
		return
	}

	pars.ParsPage(w, "regstat", con.KeshTempl, http.StatusOK)

}

func (con *Connect) handlerAuthPage(w http.ResponseWriter, r *http.Request) {
	log.Println("connection auth")

	pars.ParsPage(w, "auth", con.KeshTempl, nil)

}

func (con *Connect) handlerAuthProcess(w http.ResponseWriter, r *http.Request) {
	log.Println("process auth")

	if r.Method != "POST" {
		err := fmt.Errorf("method not post")
		myerr.ServesError(w, con.KeshTempl, err)
		return
	}
	if len(r.FormValue("name")) == 0 || len(r.FormValue("password")) == 0 {
		err := fmt.Errorf("not se value")
		myerr.ServesError(w, con.KeshTempl, err)
		return
	}

	resolt, err := comsql.GetAccountData(con.MySQL, r.FormValue("name"))
	if err != nil {
		myerr.ServesError(w, con.KeshTempl, err)
		return
	}

	userverific := authe.New(resolt, r)
	userverific.Compare()
	token, err := userverific.CreateJWTRefresh(con.KeyRef)
	if err != nil {
		myerr.ServesError(w, con.KeshTempl, err)
		return
	}

	http.SetCookie(w, cookiepkg.CreateCookieAouth(token))

	pars.ParsPage(w, "regstat", con.KeshTempl, http.StatusOK)
}
