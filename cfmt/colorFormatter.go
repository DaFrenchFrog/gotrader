package cfmt

import (
	"fmt"
)

type Color string

const (
	Neutral Color = "\033[0m"
	Red     Color = "\033[31m"
	Green   Color = "\u001B[32m"
	Yellow  Color = "\033[33m"
	Blue    Color = "\033[34m"
	Purple  Color = "\033[35m"
	Cyan    Color = "\033[36m"
	White   Color = "\033[37m"
)

// Println :
func Println(color Color, a ...interface{}) {
	fmt.Print(color)
	fmt.Print(a...)
	fmt.Println(Neutral)
}
