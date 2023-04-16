package web

import (
	"net/http"
	"pet/iternal/s_main/web/validurl"
)

func Rout() *http.ServeMux {
	mux := http.NewServeMux()
	mux.HandleFunc("/", validurl.ValidUrl(hendlerMain))
	return mux
}
