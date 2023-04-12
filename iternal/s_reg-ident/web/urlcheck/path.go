package urlcheck

import (
	"errors"
	"log"
	"net/http"
	"pet/pkg/myerr"
	"regexp"
)

var validPath = regexp.MustCompile("^/(reg|auth)?(/process)?$")

func CheckURL(fn func(w http.ResponseWriter, r *http.Request)) http.HandlerFunc {
	log.Println("not valid path")
	return func(w http.ResponseWriter, r *http.Request) {
		m := validPath.FindStringSubmatch(r.URL.Path)
		if m == nil {
			myerr.ServesError(w, errors.New("not valid path"))
			return
		}
		fn(w, r)
	}
}
