package base

/**************************************** 接口 ****************************************/
type WebServer interface {
	Work(*Request) *Result
	Addr() string
	SSLCertPath() string
	ErrorCatch(error)
}
