package proxy

import (
	"log"
	"net/http"
	"net/http/cookiejar"
	"net/url"

	"github.com/gin-gonic/gin"
)

func ServeMiddleware(c *gin.Context) {
	cookie1, err := c.Request.Cookie("userID")
	if err != nil{
		log.Println(err)
	}

	cookie2, err := c.Request.Cookie("access-token")
	if err != nil{
		log.Println(err)
	}

	jar, err := cookiejar.New(nil)
	if err != nil{
		log.Println(err)
	}

	urlObj, _ := url.Parse("http://localhost:8090/")
	var cookies = []*http.Cookie{cookie1, cookie2}
	client := &http.Client{
		Jar: jar,
	}
	client.Jar.SetCookies(urlObj, cookies)

	req, err := http.NewRequest("GET", "http://localhost:8090/check-access", nil)
	if err != nil{
		log.Println(err)
	}
	res, err := client.Do(req)
	if err != nil{
		log.Println(err)
	}

	if res.StatusCode == http.StatusOK{
		c.Next()
	}
}
