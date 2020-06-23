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
func (this *Https) Work(request *base.Request) base.Result {
	return base.Result{}
}

func (this *Https) Addr() string {
	return this.addr
}

/*func (this *Https) SSLCertPath() string {
	return this.addr
}*/