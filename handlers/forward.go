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
		segments := strings.Split(r.URL.Path, "/")
		if len(segments) < 2 {
			log.Println("not segments")
			w.WriteHeader(http.StatusNotFound)
			w.Write([]byte("not found"))
			return
		}
		log.Println("segs", segments)

		id := strings.TrimSpace(segments[1])
		trimmedPath := "/" + strings.Join(segments[2:], "/")
		r.URL.Path = trimmedPath

		subprocess := p.Get(id)
		if subprocess == nil {
			log.Println("not okay. invalid id", id)
			w.WriteHeader(http.StatusNotFound)
			w.Write([]byte("not found"))
			return
		}

		log.Printf("id: %s target: %s", id, subprocess.Url)
		proxy := httputil.NewSingleHostReverseProxy(subprocess.Url)
		proxy.ServeHTTP(w, r)
	}
}
