package web

import (
	"fmt"
	"log"
	"net/http"
	"pet/iternal/s_reg-ident/cookies"
	"pet/iternal/s_reg-ident/db/comsql"
	"pet/iternal/s_reg-ident/str/authe"
	"pet/iternal/s_reg-ident/str/regin"
	"pet/iternal/s_reg-ident/str/salt"
	"pet/pkg/myerr"
	"pet/pkg/pars"
)

func (con *Connect) handlerMain(w http.ResponseWriter, r *http.Request) {
	r.UserAgent()
	cookie, err := r.Cookie("RefJWT")
	if err != nil {
		myerr.ServesError(w, err)
		return
	}
	log.Printf("USER %s", r.UserAgent())
	log.Printf("JWT %v", cookie.Value)
	pars.ParsPage(w, "ui/HTML/regstat2.html")
}

func handlerRegPage(w http.ResponseWriter, r *http.Request) {
	log.Println("connection reg")

	pars.ParsPage(w, "ui/HTML/reg3.html")

}

func (con *Connect) handlerRegProcess(w http.ResponseWriter, r *http.Request) {
	log.Println("process reg")

	if len(r.FormValue("name")) == 0 || len(r.FormValue("password")) == 0 || len(r.FormValue("email")) == 0 {
		err := fmt.Errorf("not se value")
		myerr.ServesError(w, err)
		return
	}

	err := comsql.CheckUinquenessLogin(con.MySQL, r)
	if err != nil {
		myerr.ServesError(w, err)
		return
	}

	err = comsql.CheckUinquenessEmail(con.MySQL, r)
	if err != nil {
		myerr.ServesError(w, err)
		return
	}

	rd, err := regin.New(r)
	if err != nil {
		myerr.ServesError(w, err)
		return
	}

	salt := salt.GenerateSalt()
	key := salt.GeneraterKey(rd.GetPass())

	err = comsql.SendRegData(con.MySQL, salt, rd, key)
	if err != nil {
		myerr.ServesError(w, err)
		return
	}

	pars.ParsPage(w, "ui/HTML/regstat2.html")

}

func handlerAuthPage(w http.ResponseWriter, r *http.Request) {
	log.Println("connection auth")

	pars.ParsPage(w, "ui/HTML/auth.html")

}

func (con *Connect) handlerAuthProcess(w http.ResponseWriter, r *http.Request) {
	log.Println("process auth")

	if r.Method != "POST" {
		err := fmt.Errorf("method not post")
		myerr.ServesError(w, err)
		return
	}
	if len(r.FormValue("name")) == 0 || len(r.FormValue("password")) == 0 {
		err := fmt.Errorf("not se value")
		myerr.ServesError(w, err)
		return
	}

	resolt, err := comsql.GetAccountData(con.MySQL, r.FormValue("name"))
	if err != nil {
		myerr.ServesError(w, err)
		return
	}

	a := authe.New(resolt, r)
	a.Compare()
	token, err := a.CreateJWT(con.K)
	if err != nil {
		myerr.ServesError(w, err)
		return
	}

	http.SetCookie(w, cookies.CreateCookieAouth(token))

	pars.ParsPage(w, "ui/HTML/regstat2.html")
}
