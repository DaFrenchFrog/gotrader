package cfmt

import (
	"fmt"
)

type Color string

const (
	Neutral Color = "\033[0m"
	Red     Color = "\033[31m"
	Green   Color = "\u001B[32m"
)

func Println(color Color, a ...interface{}) {
	fmt.Print(color)
	fmt.Print(a...)
	fmt.Println(Neutral)
}
