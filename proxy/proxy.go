package proxy

import (
	"fmt"
	"log"
	"net/http"
	"sync"
	"time"

	"github.com/kingtingthegreat/reverse-proxy-server/subprocess"
)

type Proxy struct {
	m            map[string]*subprocess.Subprocess
	mu           sync.RWMutex
	t            *time.Duration
	Server       *http.Server
	ServerFunc   func() *http.Server
	ServerFuncId func(string) *http.Server
}

func (p *Proxy) Get(id string) *subprocess.Subprocess {
	if id == "" {
		return nil
	}

	p.mu.RLock()
	defer p.mu.RUnlock()

	v, ok := p.m[id]
	if !ok {
		return nil
	}

	// this child is still active
	go func() { v.Active <- true }()

	return v
}

func (p *Proxy) Insert(id string, sub *subprocess.Subprocess) error {
	if id == "" {
		return fmt.Errorf("id must not be empty")
	}

	p.mu.Lock()
	defer p.mu.Unlock()

	if _, ok := p.m[id]; ok {
		return fmt.Errorf("this id is taken")
	}

	p.m[id] = sub
	go func() {
		if p.t == nil {
			log.Println("no timeout set for", id)
			return
		}

		ticker := time.NewTicker(*p.t)
		for {
			select {
			case <-sub.Active:
				fmt.Println("reset ticker due to activity", id)
				ticker = time.NewTicker(*p.t)
				continue
			case <-ticker.C:
				p.mu.Lock()
				l := len(p.m)
				fmt.Println("length", l)
				if l <= 1 {
					p.mu.Unlock()
					fmt.Println("reset ticker due to last one standing:", id)
					ticker = time.NewTicker(*p.t)
					continue
				}
				log.Println("deleting", id)
				delete(p.m, id)
				p.mu.Unlock()
				log.Println("KILLING PROCESS", id)
				sub.Kill <- true
				return
			}
		}
	}()

	fmt.Printf("proxy subserver: %s running at %s\n", id, sub.Url.Host)
	return nil
}

func (p *Proxy) Delete(id string) {
	p.mu.Lock()
	defer p.mu.Unlock()

	delete(p.m, id)
}

func (p *Proxy) Length() int {
	p.mu.RLock()
	defer p.mu.RUnlock()

	return len(p.m)
}

func (p *Proxy) Keys() []string {
	p.mu.RLock()
	defer p.mu.RUnlock()

	i := 0
	keys := make([]string, len(p.m))
	for k := range p.m {
		keys[i] = k
		i++
	}

	return keys
}

func NewProxy(timeout *time.Duration) *Proxy {
	return &Proxy{
		m:  make(map[string]*subprocess.Subprocess),
		mu: sync.RWMutex{},
		t:  timeout,
	}
}

func NewProxyWithServer(server *http.Server, timeout *time.Duration) *Proxy {
	p := NewProxy(timeout)
	p.Server = server
	return p
}

func NewProxyWithServerFunc(serverFunc func() *http.Server, timeout *time.Duration) *Proxy {
	p := NewProxy(timeout)
	p.ServerFunc = serverFunc
	return p
}

func NewProxyWithServerFuncId(serverFuncId func(string) *http.Server, timeout *time.Duration) *Proxy {
	p := NewProxy(timeout)
	p.ServerFuncId = serverFuncId
	return p
}
