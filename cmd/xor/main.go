package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"
	"strconv"

	"github.com/Brotiger/go_xor_crypto/internal/xor"
)

var (
	shift int
)

func init() {
	flag.IntVar(&shift, "shift", 1, "shift")
}

func main() {
	flag.Parse()
	myscanner := bufio.NewScanner(os.Stdin)

	fmt.Print("Введите строку: ")

	myscanner.Scan()
	input := myscanner.Text()
	fmt.Print("Введите пароль: ")
	myscanner.Scan()
	password := myscanner.Text()
	fmt.Print("Результат: ")
	crypto := xor.EncryptDecrypt(input, password, shift)
	fmt.Print("\n")
	fmt.Print(strconv.Quote(crypto))

	fmt.Print("\n")
}
