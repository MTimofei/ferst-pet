package web

import (
	"errors"
	"net/http"
	"pet/pkg/myerr"
	"regexp"
)

var validPath = regexp.MustCompile("^/(main)$")

func (con *Connect) ValidUrl(fn func(w http.ResponseWriter, r *http.Request)) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		m := validPath.FindStringSubmatch(r.URL.Path)
		if m == nil {
			myerr.ServesError(w, con.PageKesh, errors.New("not valid path"))
			return
		}
		fn(w, r)
	}
}
