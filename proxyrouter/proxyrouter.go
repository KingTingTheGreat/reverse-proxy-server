package proxyrouter

import (
	"net/http"

	"github.com/kingtingthegreat/reverse-proxy-server/handlers"
	"github.com/kingtingthegreat/reverse-proxy-server/proxy"
)

func ProxyRouter(p *proxy.Proxy) *http.ServeMux {
	router := http.NewServeMux()

	router.HandleFunc("/spawn", handlers.SpawnHandler(p))
	router.HandleFunc("/sub/{id}/", handlers.ForwardHandler(p))
	router.HandleFunc("/", handlers.HomeHandler(p))

	return router
}
