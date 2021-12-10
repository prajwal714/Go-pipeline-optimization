package main

import "fmt"

const version="dev"

func main(){
	fmt.Printf("Version: %s\n", version)
	fmt.Print(Hello())
}

func Hello() string{
	return "Hello World"
}

