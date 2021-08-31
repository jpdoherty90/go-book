package main

import (
	"fmt"
	"os"
)

func main() {
	fmt.Println(os.Args[0])
	for i, val := range os.Args[1:] {
		fmt.Printf("%d : %s\n", i, val)
	}
}

func echo1() {
	var s, sep string
	for i := 1; i < len(os.Args); i++ {
		s += sep + os.Args[i]
		sep = " "
	}
	fmt.Println(s)
}

func echo2() {
	s, sep := "", ""
	for _, arg := range os.Args[1:] {
		s += sep + arg
		sep = " "
	}
	fmt.Println(s)
}
