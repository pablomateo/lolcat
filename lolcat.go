package main

import (
	// These I guess are Go Libraries downloaded when installing Golang
	"bufio"
	"fmt"
	"io"
	"os"
	"strings"

	// This one, on the other hand, had to be downloaded apart:
	// go get -u syreclabs.com/go/faker
	// It gets saved on the go system folder
	"syreclabs.com/go/faker"
	// New imports require to get input
)

func printInGold(output []rune) {
	// RGB values for golden color
	r, g, b := 255, 215, 0

	// For loop through each of the characters found in our output variable
	for j := 0; j < len(output); j++ {
		//Print using the ANSI color escape sequence
		fmt.Printf("\033[38;2;%d;%d;%dm%c\033[0m", r, g, b, output[j])
	}
}

func createHackerPhrase() []rune {
	// We create an array of strings
	var phrases []string
	// Create a for loop (x3)
	for i := 1; i < 3; i++ {
		// Append to the array a "Hacker Phrase" generated from the syreclabs library
		phrases = append(phrases, faker.Hacker().Phrases()...)
	}

	// Print all the string in the array one after the other
	//fmt.Println(strings.Join(phrases[:], ", "))
	// Create Output variable
	output := strings.Join(phrases[:], "; ")
	return []rune(output)
}

func main() {
	//printInGold(createHackerPhrase())
	//println(" ")

	info, _ := os.Stdin.Stat()
	var output []rune

	if info.Mode()&os.ModeCharDevice != 0 {
		fmt.Println("The command is intended to work with Pipes")
		fmt.Println("Usage: fortune | gogold")
	}

	reader := bufio.NewReader(os.Stdin)
	for {
		input, _, err := reader.ReadRune()
		if err != nil && err == io.EOF {
			break
		}
		output = append(output, input)
	}

	printInGold(output)
}
