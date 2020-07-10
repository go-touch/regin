package main

import (
	"fmt"
	"github.com/go-touch/regin/utils/multitype"
)

type Admin struct {
}

func main() {
	any := &multitype.AnyMap{"username": "admin"}

	fmt.Println(any)

	any.Set("password", "123456")

	fmt.Println(any)
	fmt.Printf("内存地址是:%p", any)

	fmt.Println(any.Get("password").ToInt())
	fmt.Printf("类型转换成:%T\n", any.Get("password1").ToInt())
	fmt.Println(any.Get("password1").ToInt())

	any.Set("username", "张三")
	fmt.Println(any)
	fmt.Printf("内存地址是:%p\n", any)
	fmt.Printf("内存地址是:%T\n", any)
	fmt.Printf("内存地址是:%T\n", any.Get().ToStringMap())
	fmt.Printf("内存地址是:%v\n", any.Get().ToStringMap())

	c := multitype.AnyMap{"username": "admin"}

	fmt.Printf("转换后的类型为:%T\n", multitype.Eval(c).ToStringMap())
	fmt.Printf("转换后的值为:%v\n", multitype.Eval(c).ToStringMap())

	//d := []interface{}{"admin"}

	//fmt.Println(d[1])

	/*if 1 == "1" {

	}*/

	t()

}

func t(arg ...int) {
	//fmt.Println(arg[0])
	fmt.Printf("参数类型:%T", arg)

	if arg == nil {
		fmt.Println("nil")
	}

	fmt.Println(len(arg))
}
