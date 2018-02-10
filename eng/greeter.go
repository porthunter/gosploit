package main

import "fmt"

type module string

func (g module) Exploit() {
	fmt.Println("Hello Universe")
}

// exported
var GosploitModule module
