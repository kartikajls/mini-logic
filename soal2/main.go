package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	reader := bufio.NewReader(os.Stdin)

	for {
		fmt.Print("Enter an integer: ")

		input, _ := reader.ReadString('\n')
		input = strings.TrimSpace(input)

		n, err := strconv.Atoi(input)
		if err != nil || n < 0 {
			fmt.Println("Invalid Input.") //jika user tidak sengaja menuliskan string maka perintah akan error
			continue
		}

		// Menghitung jumlah 1 sampai n
		sum := 0
		for i := 1; i <= n; i++ {
			sum += i
		}

		fmt.Printf("The sum of number from 1 to %d is %d\n", n, sum)
	}
}
