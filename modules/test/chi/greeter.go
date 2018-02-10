package main

import "fmt"

type greeting string

func (g greeting) Exploit() {
	fmt.Println("你好宇宙")
}

var GosploitModule greeting
