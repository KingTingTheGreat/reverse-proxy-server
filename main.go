package main

import (
	"fmt"
	"log"
	"time"

	"github.com/kingtingthegreat/reverse-proxy-server/mock"
	"github.com/kingtingthegreat/reverse-proxy-server/proxy"
	"github.com/kingtingthegreat/reverse-proxy-server/proxyserver"
	"github.com/kingtingthegreat/reverse-proxy-server/subprocess"
)

const n = 5

func main() {
	t := time.Minute * 15
	p := proxy.NewProxyWithServerFuncId(mock.Server, &t)

	fmt.Println("creating a number of subprocesses", n)
	for _, id := range mock.Id(n) {
		s := subprocess.Spawn(mock.Server(id))
		p.Insert(id, s)
	}

	ps := proxyserver.ProxyServer(":8080", p)
	err := ps.ListenAndServe()
	log.Fatal(err)
}
