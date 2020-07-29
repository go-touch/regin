package main

import (
	"fmt"
	"github.com/go-touch/mtype"
	"github.com/go-touch/regin/utils/curl"
)

func main()  {

	post := curl.Post()
	ret := post.Call("http://127.0.0.1:8080/demo/index",mtype.AnyMap{})
	fmt.Println(ret.ToAnyMap())
}
