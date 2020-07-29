package app

import "github.com/go-touch/regin/base"

type Https struct {
	base.WebServer
	addr string
}

func NewHttps() *Https {
	return &Https{
		addr: "127.0.0.1:443",
	}
}

//
func (hs *Https) Work(request *base.Request) *base.Result {
	return &base.Result{}
}

func (hs *Https) Addr() string {
	return hs.addr
}

func (hs *Https) SSLCertPath() string {
	return hs.addr
}

func (hs *Https) ErrorCatch() error {
	return nil
}
