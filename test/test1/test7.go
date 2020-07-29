package main

import (
	"fmt"
)

func main() {
	defer func() {
		if r := recover(); r != nil {
			fmt.Printf("捕获到错误信息:%v\n",r.(string))
			//debug.PrintStack()
			//err = errors.New(r.(string))
		}
	}()
	action1()
}

func action1() {
	action2()
}

func action2() {
	action3()
}

func action3() {
	panic("错误")
}

func errorCatch(){
	if r := recover(); r != nil {
		fmt.Printf("捕获到错误信息:%v\n",r.(string))
		//debug.PrintStack()
		//err = errors.New(r.(string))
	}
}
