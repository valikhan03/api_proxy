package proxy

import(
	"net/http"
	"net/http/httputil"
	"github.com/gin-gonic/gin"
)

type AuthApiProxy struct{
	Target string
}

func NewAuthApiProxy(target string) *AuthApiProxy{
	return &AuthApiProxy{
		Target: target,
	}
}

func (p *AuthApiProxy) ReverseAuthProxy(req_path string) gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Request.URL.Scheme = "http"
		c.Request.URL.Host = p.Target
		c.Request.URL.Path = req_path
		director := func(req *http.Request) {
			req = c.Request
		}
		proxy := &httputil.ReverseProxy{Director: director}
		proxy.ServeHTTP(c.Writer, c.Request)
	}
}



