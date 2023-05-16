package main

import (
	"piacheia/mml"
	"fmt"
)

func main() {
		x := mml.ReadFile("text.txt")
		fmt.Println(x.Name())
}
