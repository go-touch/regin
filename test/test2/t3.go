package main

import (
	"fmt"
	"time"
)

func main() {
	count := 0
	startTime := time.Now()

	for {
		count++
		if count > 100 {
			break
		}
		fmt.Printf("输出数字:%d\n", count)
	}

	// endTime := time.Now()


	d := time.Since(startTime).String()

	fmt.Println(d)


	/*fmt.Println("开始时间", startTime, "\nsss")
	fmt.Println("结束时间", endTime, "\nsss")

	// rat := big.NewRat(endTime,startTime)

	a := big.NewInt(startTime)
	b := big.NewInt(endTime)
	z := big.NewInt(1)

	ret := z.Sub(b, a)

	fmt.Println("相减", ret, "\nsss")
	fmt.Printf("类型:%T\n", ret)

	c := big.NewInt(1000000000)

	fmt.Println("除法", ret.Div(ret, c), "\nsss")*/
}
