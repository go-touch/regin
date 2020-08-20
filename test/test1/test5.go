package main

import (
	"fmt"
	"github.com/go-touch/regin/utils/validator"
)

type PlusUsers struct {
	UserId  int    `key:"user_id" require:"true" length:"0|5"`
	Account string `key:"account" require:"true" length:"0|20"`
}

func main() {
	validator.RegisterForm(&PlusUsers{}, "PlusUsers")
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
