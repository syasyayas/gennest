package main

import (
	"flag"
	"fmt"
	"os"
)

func main() {
	level := 1
	flag.IntVar(&level, "level", 0, "levels of nest")
	funcName := flag.String("funcName", "wtf", "function name")

	file := flag.String("file", "", "file to write to")
	flag.Parse()
	write := "package main\nimport \"fmt\"\n" + Generate(level, "fmt.Println(\"Hello\")", *funcName)
	err := os.WriteFile(*file, []byte(write), 0644)
	if err != nil {
		panic(err)
	}
	c := GenerateFuncCall(level, *funcName)
	fmt.Println("function call bozo :\n", c)

}
