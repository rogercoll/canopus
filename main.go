package main


import (
    "fmt"
    "strings"
    "syscall"
    "golang.org/x/crypto/ssh/terminal"
)

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
	pass := credentials()
	fmt.Printf("Password: %s\n", pass)
}