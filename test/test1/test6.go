package main

import (
	"fmt"
	"github.com/go-touch/mtype"
	"math"
	"math/big"
)

func main() {
	slice := mtype.AnyMapSlice{
		{"username": "admin1"},
		{"username": "admin2"},
		{"username": "admin3"},
		{"username": "admin4"},
		{"username": "admin5"},
		{"username": "admin6"},
		{"username": "admin7"},
		{"username": "admin8"},
		{"username": "admin9"},
	}

	sliceGroup := LimitAnyMapSlice(slice, 4)
	fmt.Printf("value:%v\n", sliceGroup)
	fmt.Printf("num:%d\n", len(sliceGroup))
}

// 计算分页数组
func LimitAnyMapSlice(slice mtype.AnyMapSlice, limit int) (sliceGroup []mtype.AnyMapSlice) {
	length := len(slice)
	if length < limit {
		sliceGroup = append(sliceGroup, slice)
		return sliceGroup
	}
	// 计算分段
	bigRat := big.NewRat(int64(length), int64(limit))
	divValue, _ := bigRat.Float64()
	count := int(math.Ceil(divValue))
	for i := 1; i <= count; i++ {
		offset := (i - 1) * limit
		maxLimit := i * limit
		if maxLimit > length {
			maxLimit = length
		}
		sliceGroup = append(sliceGroup, slice[offset:maxLimit])
	}
	return sliceGroup
}
