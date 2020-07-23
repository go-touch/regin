package main

import "fmt"

type Container map[string]*Tx
type Tx struct {
	name string
}

type A struct {
	tx *Tx
}

type B struct {
	tx *Tx
}

var c Container

func init() {
	c = Container{"tx": &Tx{"张三"}}
}

func main() {
	testa := A{tx: c["tx"]}
	testb := B{tx: c["tx"]}

	fmt.Printf("testa的值为:%v\n", testa)
	fmt.Printf("testb的值为:%v\n", testb)

	fmt.Printf("c的值为:%v\n", c)
	fmt.Println(c)

	*c["tx"] = nil

	fmt.Printf("testa的值为:%v\n", testa.tx)
	fmt.Printf("testb的值为:%v\n", testb.tx)

}
