package webtest

import (
	"fmt"
	"log"
	"net/http"
	"net/url"
)

func Req1(c *http.Client, urlreg, urlauth, urlmain string) {
	for i := 200; i < 1000; i++ {
		name := fmt.Sprintf("test%d", i)
		email := fmt.Sprintf("test%d@test.com", i)
		resp, _ := c.PostForm(urlreg, url.Values{"name": {name}, "email": {email}, "password": {"qwertyuiop"}})
		log.Println(resp.Status, resp.Request.URL, i)
		resp2, _ := c.PostForm(urlauth, url.Values{"name": {name}, "password": {"qwertyuiop"}})
		log.Println(resp2.Status, resp.Request.URL, i)

		r, err := http.NewRequest(http.MethodGet, urlmain, nil)
		if err != nil {
			log.Println(err)
		}
		for _, cookies := range resp2.Cookies() {
			r.AddCookie(cookies)
		}
		res3, _ := c.Do(r)
		log.Println(res3.Status, resp.Request.URL, i)

	}
}

func TestServer() {
	client := &http.Client{}
	urlreg := "http://158.160.60.70/reg/process"
	urlauth := "http://158.160.60.70/auth/process"
	urlmain := "http://158.160.60.70/"

	// urlr := make(chan string)
	// urla := make(chan string)
	// urlm := make(chan string)

	// urlr <- urlreg
	// urla <- urlauth
	// urlm <- urlmain

	Req1(client, urlreg, urlauth, urlmain)

}
