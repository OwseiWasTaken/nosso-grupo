package main

import (
	"piacheia/mml"
	"fmt"
)

func main() {
		x := mml.ReadFile("text.txt")
		fmt.Println(x.Name())
		fmt.Println(string(x.Current()))
		fmt.Println(string(x.Buffer))
		fmt.Println(x)
}
