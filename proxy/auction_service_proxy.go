package proxy

import(
	"net/http"
	"net/http/httputil"
	"github.com/gin-gonic/gin"
)

type AuctionServiceProxy struct{
	Target string
}

func NewAutctionServiceProxy() *AuthApiProxy{
	return &AuthApiProxy{
		
	}
}

func (p *AuthApiProxy) ReverseAuctionServiceProxy(req_path string) gin.HandlerFunc {
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
