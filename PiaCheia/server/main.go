package main

import (
	"piacheia/mml"
	"fmt"
)

func main() {
		x := mml.ReadFile("text.txt")
		//fmt.Println("name:", x.Name())
		//fmt.Println("current:", string(x.Current()))
		//fmt.Println("buffer:", string(x.Buffer))
		//fmt.Println("stack:", x.TagStack)
		fmt.Print("")
		x.Parse()
}
