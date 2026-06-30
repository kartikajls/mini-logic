package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	// membuat variable a,b,c
	var a, b, c int

	reader := bufio.NewReader(os.Stdin)

	// membuat perintah 1
	fmt.Print("Enter the first integer: ")
	input, _ := reader.ReadString('\n')
	input = strings.TrimSpace(input)
	a, _ = strconv.Atoi(input)

	// membuat perintah 2
	fmt.Print("Enter second integer: ")
	input, _ = reader.ReadString('\n')
	input = strings.TrimSpace(input)
	b, _ = strconv.Atoi(input)

	// membuat perintah 3
	fmt.Print("Enter the third integer: ")
	input, _ = reader.ReadString('\n')
	input = strings.TrimSpace(input)
	c, _ = strconv.Atoi(input)

	//logic 1 jika semua angka sama
	if a == b && b == c {
		fmt.Println("All Three are equal.")
	} else if (a == b && b != c) || // logic 2 jika salah satu dari 3 angka tidak sama
		(a == c && a != b) ||
		(b == c && a != b) {
		fmt.Println("Two are equals and one is different.")
	} else {
		fmt.Println("All Three are different") // semua angka berbeda
	}
}
