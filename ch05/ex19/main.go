package main

import "fmt"

func main() {
	fmt.Println(f())
}

func f() (i int) {
	defer func() {
		if p := recover(); p != nil {
			i = p.(int)
		}
	}()
	panic(100)
}
