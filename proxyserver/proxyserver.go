package proxyserver

import (
	"net/http"

	"github.com/kingtingthegreat/reverse-proxy-server/handlers"
	"github.com/kingtingthegreat/reverse-proxy-server/middleware"
	"github.com/kingtingthegreat/reverse-proxy-server/proxy"
)

func ProxyServer(addr string, p *proxy.Proxy) *http.Server {
	router := http.NewServeMux()
	router.HandleFunc("/spawn", handlers.SpawnHandler(p))
	router.HandleFunc("/", handlers.ForwardHandler(p))

	middleware := middleware.CreateStack()

	server := http.Server{
		Addr:    addr,
		Handler: middleware(router),
	}

	return &server
}
