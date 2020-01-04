package utils

import (
	"os"
	"fmt"
	"bufio"
	"syscall"
	"strings"
	"golang.org/x/crypto/ssh/terminal"
)

func GetDir() string {
	reader := bufio.NewReader(os.Stdin)
    fmt.Print("Enter directory: ")
	dir, _ := reader.ReadString('\n')
	return strings.TrimSpace(dir)
}

func GetCredentials() (string) {
    fmt.Print("Enter Password: ")
    bytePassword, err := terminal.ReadPassword(int(syscall.Stdin))
    if err == nil {
        fmt.Println("\nPassword typed: " + string(bytePassword))
    }
    password := string(bytePassword)
    return strings.TrimSpace(password)
}
