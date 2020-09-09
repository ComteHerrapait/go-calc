package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	fmt.Printf("let's go\n")
	for true {
		reader := bufio.NewReader(os.Stdin)
		fmt.Print("opération : ")
		input, _ := reader.ReadString('\n')

		steps := []string{}
		marker := 0
		for _, character := range input {
			if strings.Contains("+-*/", character) { //check if character is an operator

			}
		}
		resultat := input
		fmt.Print("return :" + resultat + "\n")
	}
}
