package handlers

import (
	"fmt"
	"log"
	"net/http"
	"net/http/httputil"
	"strings"

	"github.com/kingtingthegreat/reverse-proxy-server/proxy"
)

func makeAnchor(href string) string {
	return fmt.Sprintf("<a href=\"/sub/%s\">%s</a>", href, href)
}

func HomeHandler(p *proxy.Proxy) func(http.ResponseWriter, *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		log.Println("forward handler")
		id := r.PathValue("id")

		subprocess := p.Get(id)
		if subprocess == nil {
			log.Println("not okay. invalid id", id)
			w.WriteHeader(http.StatusNotFound)

			keys := p.Keys()
			s := ""
			for _, k := range keys {
				s += makeAnchor(k)
			}

			w.Write([]byte("<!doctype html>please pick a lobby " + s))
			return
		}

		newUrlPath := strings.SplitN(r.URL.Path, id, 2)
		r.URL.Path = newUrlPath[1]

		log.Printf("id: %s target: %s", id, subprocess.Url)
		proxy := httputil.NewSingleHostReverseProxy(subprocess.Url)
		proxy.ServeHTTP(w, r)
	}
}
