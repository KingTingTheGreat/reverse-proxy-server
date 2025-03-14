package handlers

import (
	"log"
	"net/http"
	"net/http/httputil"
	"strings"

	"github.com/kingtingthegreat/reverse-proxy-server/proxy"
)

func ForwardHandler(p *proxy.Proxy) func(http.ResponseWriter, *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		log.Println("forward handler")

		id := r.PathValue("id")
		subprocess := p.Get(id)
		if subprocess == nil {
			log.Println("not okay. invalid id", id)
			w.WriteHeader(http.StatusNotFound)
			w.Write([]byte("not found"))
			return
		}

		// id should be a valid split bc pathvalue correspond to subprocess
		newUrlPath := strings.SplitN(r.URL.Path, id, 2)
		r.URL.Path = newUrlPath[1]

		log.Printf("id: %s target: %s", id, subprocess.Url)
		proxy := httputil.NewSingleHostReverseProxy(subprocess.Url)
		proxy.ServeHTTP(w, r)
	}
}
