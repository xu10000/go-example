package main

import (
	"fmt"
	"reflect"
)

const LEN int = 3

func SetArrValue(iArr []interface{}) {
	for i := 0; i < len(iArr); i++ {

		v := reflect.ValueOf(iArr[i])
		// fmt.Println("------ print kind", v.Kind())
		if v.Kind() != reflect.Ptr && v.Kind() != reflect.Interface {
			panic("iArr item must be ptr or interface")
		}

		// 直接设置string类型,如果多种要判断
		v.Elem().SetString("set string i ")
	}
}

func main() {
	strArr := make([]string, LEN)
	iArr := make([]interface{}, LEN)

	for i := 0; i < len(iArr); i++ {
		strArr[i] = "zero"
		iArr[i] = &strArr[i]
	}
	// 通过iArr的item引用来修改strArr的值
	SetArrValue(iArr)
	fmt.Println("------ print strArr", strArr)
}
