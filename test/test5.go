package main

import (
	"fmt"
	"github.com/go-touch/regin/utils/validator"
)

type PlusUsers struct {
	UserId int    `key:"user_id" require:"true" length:"4|5"`
	Test   string `key:"test" require:"true" length:"0|255"`
}

func main() {
	validator.RegisterForm(&PlusUsers{},"PlusUsers")
	dao := validator.Form("PlusUsers")

	for _, v := range dao.FieldMap {
		fmt.Println(v.Tags)
	}

	result := dao.Verify(&map[string]interface{}{
		"user_id": "测",
	})

	for _, v := range result {
		fmt.Println(v)
	}

}

func verify(m map[string]interface{}) {
	fmt.Println("测试")
}
