package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/Brotiger/go_xor_crypto/internal/xor"
)

var (
	input    string
	password string
	shift    int
)

func init() {
	flag.IntVar(&shift, "shift", 1, "shift")
}

func main() {
	flag.Parse()

	fmt.Print("Введите строку: ")
	fmt.Fscan(os.Stdin, &input)
	fmt.Print("Введите пароль: ")
	fmt.Fscan(os.Stdin, &password)

	fmt.Print(xor.EncryptDecrypt(input, password, shift))
}
