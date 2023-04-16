package web

import (
	"net/http"
	"pet/pkg/pars"
)

func hendlerMain(w http.ResponseWriter, r *http.Request) {
	pars.ParsPage(w, "ui/HTML/mainpage.html")
}
