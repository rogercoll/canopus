package main


import (
	"os"
	"fmt"
	"bufio"
    "strings"
    "syscall"
	"golang.org/x/crypto/ssh/terminal"
	"github.com/rogercoll/dirEncryptor/encrypter"
)

func readDir() string {
	reader := bufio.NewReader(os.Stdin)

    fmt.Print("Enter directory: ")
	dir, _ := reader.ReadString('\n')
	return strings.TrimSpace(dir)
}

func credentials() (string) {
    fmt.Print("Enter Password: ")
    bytePassword, err := terminal.ReadPassword(int(syscall.Stdin))
    if err == nil {
        fmt.Println("\nPassword typed: " + string(bytePassword))
    }
    password := string(bytePassword)
    return strings.TrimSpace(password)
}

func main() {
	dir := readDir()
	fmt.Printf("Directory: %s\n", dir)
	pass := credentials()
	fmt.Printf("Password: %s\n", pass)
	err := encrypter.Encrypt(dir, pass)
	if err != nil {
		fmt.Println("Directory encripted correctly!")
	}
}