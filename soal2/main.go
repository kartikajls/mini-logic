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

	fmt.Print("Enter an integer: ")
	input, _ := reader.ReadString('\n')
	input = strings.TrimSpace(input)

	// Konversi string ke integer
	n, err := strconv.Atoi(input)
	if err != nil {
		fmt.Println("Error") //jika user tidak sengaja menuliskan string maka perintah akan error
		return
	}

	// Validasi angka negatif
	if n < 0 {
		fmt.Println("Invalid Input")
		return
	}

	// Menghitung jumlah 1 sampai n
	sum := 0
	for i := 1; i <= n; i++ {
		sum += i
	}

	fmt.Printf("The sum of number from 1 to %d is %d\n", n, sum)
}
