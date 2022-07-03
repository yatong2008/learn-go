package main

import "fmt"

func variableZeroValue() {
	var a int
	var s string
	//fmt.Println(a, s) //use Println cannot see the empty string value
	fmt.Printf("%d %q\n", a, s) //use Printf can show the empty in quote
}

func main() {
	fmt.Println("hello world")
	variableZeroValue()
}
