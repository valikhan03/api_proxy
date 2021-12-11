package proxy

import (
	"net/http"
	"net/http/httputil"

	"github.com/gin-gonic/gin"
)

type Service struct {
	ServiceName string `json:"service"`
	Host        string `json:"host"`
	Protocol    string `json:"protocol"`
	RequestPath string `json:"req_path"`
	Method      string `json:"method"`
}

func NewService(service_name string) *Service {
	return ReadServiceConfigs(service_name)
}

func (s *Service) ServeService() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Request.URL.Scheme = s.Protocol
		c.Request.URL.Host = s.Host
		c.Request.URL.Path = s.RequestPath
		director := func(req *http.Request) {
			req = c.Request
		}
		proxy := &httputil.ReverseProxy{Director: director}
		proxy.ServeHTTP(c.Writer, c.Request)
	}
}

func RegisterServiceEndpoint(router *gin.Engine, service *Service, path string) {
	switch service.Method{
	case "POST":
		router.POST(path, service.ServeService())
	case "GET":
		router.GET(path, service.ServeService())
	case "PUT":
		router.PUT(path, service.ServeService())
	case "DELETE":
		router.DELETE(path, service.ServeService())
	case "ANY":
		router.Any(path, service.ServeService())
	}
}
