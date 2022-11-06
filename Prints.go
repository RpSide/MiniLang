package main

import (
	"fmt"
)

func PrintLnRed(Text string) {

	colorReset := "\033[0m"

	color := "\033[31m"

	fmt.Println(string(color) + Text + string(colorReset))

}

func PrintLnGreen(Text string) {

	colorReset := "\033[0m"

	color := "\033[32m"

	fmt.Println(string(color) + Text + string(colorReset))

}

func PrintLnYellow(Text string) {

	colorReset := "\033[0m"

	color := "\033[33m"

	fmt.Println(string(color) + Text + string(colorReset))

}

func PrintLnBlue(Text string) {

	colorReset := "\033[0m"

	color := "\033[34m"

	fmt.Println(string(color) + Text + string(colorReset))

}

func PrintLnPurple(Text string) {

	colorReset := "\033[0m"

	color := "\033[35m"

	fmt.Println(string(color) + Text + string(colorReset))

}

func PrintLnCyan(Text string) {

	colorReset := "\033[0m"

	color := "\033[36m"

	fmt.Println(string(color) + Text + string(colorReset))

}

func PrintLnWhite(Text string) {

	colorReset := "\033[0m"

	color := "\033[37m"

	fmt.Println(string(color) + Text + string(colorReset))

}
