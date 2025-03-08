package main

import (
	"fmt"
	"log"

	"github.com/kingtingthegreat/reverse-proxy-server/mock"
	"github.com/kingtingthegreat/reverse-proxy-server/proxy"
	"github.com/kingtingthegreat/reverse-proxy-server/proxyserver"
	"github.com/kingtingthegreat/reverse-proxy-server/subprocess"
)

const n = 5

func main() {
	// t := time.Minute * 15
	p := proxy.NewProxyWithServerFuncId(mock.MockServer, nil)

	fmt.Println("creating a number of subprocesses", n)
	for _, id := range mock.Id(n) {
		s, err := subprocess.Spawn(mock.MockServer(id))
		if err != nil {
			panic("could not spawn subprocess: " + err.Error())
		}
		p.Insert(id, s)
	}

	ps := proxyserver.ProxyServer(":8080", p)
	err := ps.ListenAndServe()
	log.Fatal(err)
}
