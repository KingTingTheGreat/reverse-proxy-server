package subprocess

import (
	"fmt"
	"log"
	"net"
	"net/http"
	"net/url"
)

func Spawn(s *http.Server) *Subprocess {
	listener, err := net.Listen("tcp", "127.0.0.1:0")
	if err != nil {
		panic(err)
	}

	active := make(chan bool, 10)
	kill := make(chan bool, 10)

	go func() {
		defer listener.Close()

		go func() {
			err := s.Serve(listener)
			log.Println("error: ", err)
		}()

		<-kill
		s.Close()
	}()

	addr := listener.Addr().(*net.TCPAddr)
	url := &url.URL{
		Scheme: "http",
		Host:   net.JoinHostPort(addr.IP.String(), fmt.Sprintf("%d", addr.Port)),
	}

	return &Subprocess{
		Url:    url,
		Active: active,
		Kill:   kill,
	}
}
