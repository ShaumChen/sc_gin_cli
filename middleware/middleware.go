package middleware

import (
	"fmt"
	"yogo/context"
)

func M1(c *context.Context) {
	fmt.Println("1")
}
func M2(c *context.Context) {
	fmt.Println("2")
}
func M3(c *context.Context) {
	fmt.Println("3")
}
