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
	var err error // membuat variable error ketika input yang dimasukan bukan angka

	reader := bufio.NewReader(os.Stdin)

	// membuat perintah 1
	for {
		fmt.Print("Enter the first integer: ")

		input, _ := reader.ReadString('\n')
		input = strings.TrimSpace(input)

		a, err = strconv.Atoi(input)
		if err != nil {
			fmt.Println("Invalid Input : Input yang dimasukan harus angka")
			continue
		}
		break

	}

	// membuat perintah 2
	for {
		fmt.Print("Enter second integer: ")

		input, _ := reader.ReadString('\n')
		input = strings.TrimSpace(input)

		b, err = strconv.Atoi(input)
		if err != nil {
			fmt.Println("Invalid Input : Input yang dimasukan harus angka")
			continue
		}
		break
	}

	// membuat perintah 3
	for {
		fmt.Print("Enter the third integer: ")

		input, _ := reader.ReadString('\n')
		input = strings.TrimSpace(input)

		c, err = strconv.Atoi(input)
		if err != nil {
			fmt.Println("Invalid Input : Input yang dimasukan harus angka")
			continue
		}
		break
	}

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
