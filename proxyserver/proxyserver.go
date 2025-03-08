package proxyserver

import (
	"net/http"

	"github.com/kingtingthegreat/reverse-proxy-server/middleware"
	"github.com/kingtingthegreat/reverse-proxy-server/proxy"
	"github.com/kingtingthegreat/reverse-proxy-server/proxyrouter"
)

func ProxyServer(addr string, p *proxy.Proxy) *http.Server {
	router := proxyrouter.ProxyRouter(p)

	middleware := middleware.Stack()

	server := http.Server{
		Addr:    addr,
		Handler: middleware(router),
	}

	return &server
}
