package main

import "fmt"

// this is a comment

func main() {

	fmt.Println("Hello World")
	fmt.Println("1 + 1 =", 1+1)
	fmt.Println("1 + 1 =", 1.0+1.0)
	fmt.Println(len("Hello World"))
	fmt.Println("Hello World"[1])
	fmt.Println("Hello " + "World")
	fmt.Println(true && true)
	fmt.Println(true && false)
	fmt.Println(true || true)
	fmt.Println(true || false)
	fmt.Println(!true)
	var Hello string
	Hello = "Held"
	fmt.Println(Hello)
	doStuff("hell", 2)
}

func doStuff(arg string, stuff int) {
	fmt.Println(arg)
	fmt.Println(stuff)
}
