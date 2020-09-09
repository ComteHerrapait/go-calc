package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"strconv"
)
func isNumeric(s string) bool {
    _, err := strconv.ParseFloat(s, 64)
    return err == nil
}

func calc(input string) {
	// PARSING
	steps := []string{}
	marker := 0
	for index, character := range input { //create a operation from the input string
		if strings.ContainsRune("+-*/", character) { //check if character is an operator
			number := input[marker:index]
			if !isNumeric(number) {
				fmt.Printf("Error NaN : " + number + " is not a valid number.\n")
				return
			}
			steps = append(steps, number)
			steps = append(steps, string(character))
			marker = index + 1
		}
	}
	number := input[marker:len(input)-2]
	if isNumeric(number) {
		steps = append(steps, number)
	} else {
		fmt.Printf("Error NaN : " + number + " is not a valid number.\n")
		return
	}
	// PROCESSING
	for {
		for index,step := range steps {
			if step == "/" {
				prev, _ :=  strconv.ParseFloat(steps[index-1], 64) 
				next, _ :=  strconv.ParseFloat(steps[index+1], 64) 
				res := fmt.Sprintf("%f", prev / next)
				steps = append(steps[:index-1], res, steps[index+2:len(steps)-1])
				break
			}
			if index == len(steps) - 1 {
				break
			}
		}
	}
	resultat := input
	fmt.Println(strings.Join(steps, ", "))
	fmt.Print("return :" + resultat + "\n")
}

func main() {
	fmt.Printf("let's go\n")
	for true {
		reader := bufio.NewReader(os.Stdin)
		fmt.Print("opération : ")
		userInput, _ := reader.ReadString('\n')
		calc(userInput)
	}
}

