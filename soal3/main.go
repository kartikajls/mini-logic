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
		// perintah 1
		fmt.Print("Masukkan sebuah angka: ")
		input, _ := reader.ReadString('\n')
		input = strings.TrimSpace(input)

		//jika user memasukan selain angka maka akan error
		angka, err := strconv.Atoi(input)
		if err != nil {
			fmt.Println("Error: Input harus berupa angka!")
			return
		}

		if angka%2 == 0 {
			fmt.Println(angka, "adalah bilangan Genap.")
		} else {
			fmt.Println(angka, "adalah bilangan Ganjil.")
		}

		//perintah 2
		fmt.Println("Apakah Anda Ingin Melanjutkan (y/n)")
		endInput, _ := reader.ReadString('\n')
		input = strings.TrimSpace(strings.ToLower(endInput))

		if input == "n" {
			break
		}
	}

}
