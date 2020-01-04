package utils

import (
	"syscall"
	"strings"
	"golang.org/x/crypto/ssh/terminal"
)


func GetCredentials() (string) {
    fmt.Print("Enter Password: ")
    bytePassword, err := terminal.ReadPassword(int(syscall.Stdin))
    if err == nil {
        fmt.Println("\nPassword typed: " + string(bytePassword))
    }
    password := string(bytePassword)
    return strings.TrimSpace(password)
}
