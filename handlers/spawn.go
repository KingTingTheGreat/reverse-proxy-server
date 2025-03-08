package handlers

import (
	"net/http"
	"strings"

	"github.com/kingtingthegreat/reverse-proxy-server/proxy"
	"github.com/kingtingthegreat/reverse-proxy-server/subprocess"
)

func SpawnHandler(p *proxy.Proxy) func(http.ResponseWriter, *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		id := strings.TrimSpace(r.URL.Query().Get("id"))
		if id == "" {
			w.WriteHeader(http.StatusBadRequest)
			w.Write([]byte("id must not be empty"))
			return
		}

		s := &http.Server{}
		if p.ServerFuncId != nil {
			s = p.ServerFuncId(id)
		} else if p.ServerFunc != nil {
			s = p.ServerFunc()
		} else {
			s = p.Server
		}

		subprocess, err := subprocess.Spawn(s)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte("could not create subprocess. try again later"))
			return
		}

		err = p.Insert(id, subprocess)
		if err != nil {
			w.WriteHeader(http.StatusConflict)
			w.Write([]byte("this id already exists"))
			subprocess.Kill <- true
			return
		}

		w.WriteHeader(http.StatusOK)
		w.Write([]byte(subprocess.Url.Host))
	}
}
