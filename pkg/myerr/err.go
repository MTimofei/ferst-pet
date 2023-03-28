package myerr

import (
	"fmt"
	"log"
	"net/http"
)

func ServesError(w http.ResponseWriter, err error) {
	log.Println(err)
	http.Error(w, fmt.Sprintf("serve err:%v", err), http.StatusMethodNotAllowed)

}
