package main

import (
	"fmt"
	"reflect"
)

type Int int

func main() {
	var a Int = 1

	// reflect.TypeOf() 返回 reflect.Type
	t := reflect.TypeOf(a)
	// reflect.ValueOf() 返回 reflect.Value
	v := reflect.ValueOf(a)

	fmt.Printf("t: %v\n", t)           // type: main.Int
	fmt.Printf("kind: %v\n", t.Kind()) // kind: int (底层类型)
	fmt.Printf("v: %v\n", v)           // value: 1

	// Type 和 Value 的常用方法
	fmt.Printf("Type name: %s\n", t.Name())            // 返回string: "Int"
	fmt.Printf("Type string: %s\n", t.String())        // 返回string: "main.Int"
	fmt.Printf("Value type: %s\n", v.Type())           // 返回reflect.Type
	fmt.Printf("Value kind: %s\n", v.Kind())           // 返回reflect.Kind
	fmt.Printf("Value interface: %v\n", v.Interface()) // 返回interface{}: 1
}
