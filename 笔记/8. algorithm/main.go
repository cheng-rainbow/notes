package main

import (
	"fmt"
	"strings"
)

type ValidationError struct {
	Field   string
	Value   interface{}
	Message string
}

// 实现 Error 接口
func (e *ValidationError) Error() string {
	return fmt.Sprintf("validation failed for field '%s' with value '%v': %s",
		e.Field, e.Value, e.Message)
}

func validateEmail(email string) error {
	if !strings.Contains(email, "@") {
		return &ValidationError{ // 创建一个ValidationError对象返回，会自动调用 Error()
			Field:   "email",
			Value:   email,
			Message: "must contain @ symbol",
		}
	}
	return nil
}
func main() {

	var n *int
	fmt.Printf("n: %v\n", n)

	n = new(int)
	fmt.Printf("n: %v\n", n)

	x := 1
	n = &x
	*n = 10

	fmt.Printf("n: %v\n", x)
}
