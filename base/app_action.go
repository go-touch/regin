package base

type AppAction interface {
	//Factory(request *Request) AppAction
	//Request() *Request
	//Result() *Result
	BeforeExec(request *Request) (result *Result)
	Exec(request *Request) (result *Result)
}

type Action struct {
	AppAction
	request *Request
	result  *Result
}

// Get Request.
func (a *Action) Request() *Request {
	return a.request
}

// Get Request.
func (a *Action) Result() *Result {
	return a.result
}

// Before action method.
func (a *Action) BeforeExec(request *Request) (result *Result) {
	return
}

// Action method.
func (a *Action) Exec(request *Request) (result *Result) {
	return
}
