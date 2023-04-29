package synchron

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"strconv"
	"time"
)

func StartSystem(nameserves *string) {
	req, err := http.NewRequest(http.MethodGet, "http://localhost:7000/", nil)
	if err != nil {
		log.Fatal(err)
	}
	req.Header.Set("User-Agent", fmt.Sprintf("Serves %s", *nameserves))
	client := http.Client{}

	res, err := client.Do(req)
	if err != nil {
		log.Fatal(err)
	}
	body, err := io.ReadAll(res.Body)
	if err != nil {
		log.Fatal(err)
	}
	defer res.Body.Close()

	ts, err := strconv.Atoi(fmt.Sprintf("%s", body))
	if err != nil {
		log.Fatal(err)
	}
	timestart := int64(ts)
	now := time.Now().Unix()
	timeout := timestart - now
	log.Println(time.Duration(timeout) * time.Second)
	time.Sleep(time.Duration(timeout) * time.Second)
}
