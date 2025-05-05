package main

import (
	"fmt"
	"io"
	"os"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func main() {
	if len(os.Args) == 2 {
		fileName := os.Args[1]
		contents, err := os.ReadFile(fileName)
		check(err)
		fmt.Print(string(contents))
	} else if len(os.Args) == 1 {
		for {
			var input string
			fmt.Print("> ")
			_, err := fmt.Scan(&input)
			if err == io.EOF {
				fmt.Println("\nExiting")
				break
			} else if err != nil {
				panic(err)
			}
			fmt.Println(input)
		}
	} else {
		fmt.Println("Incorrect Usage!")
	}

}
