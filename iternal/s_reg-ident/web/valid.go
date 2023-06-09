package web

import (
	"errors"
	"net/http"
	"pet/pkg/myerr"
	"regexp"
)

var validPath = regexp.MustCompile("^/(reg|auth)?(/process)?$")

func (con *Connect) ValidUrl(fn func(w http.ResponseWriter, r *http.Request)) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		m := validPath.FindStringSubmatch(r.URL.Path)
		if m == nil {
			myerr.ServesError(w, con.KeshTempl, errors.New("not valid path"))
			return
		}
		fn(w, r)
	}
}
