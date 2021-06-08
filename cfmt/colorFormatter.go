package cfmt

import (
	"fmt"
	"os"
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

func Print(color Color, a ...interface{}) {
	fmt.Print(color)
	fmt.Print(a...)
	fmt.Println(Neutral)
}

// Printf formats according to a format specifier and writes to standard output with a specific color.
func Printf(color Color, format string, a ...interface{}) {
	fmt.Print(color)
	fmt.Fprintf(os.Stdout, format, a...)
	fmt.Println(Neutral)
}
