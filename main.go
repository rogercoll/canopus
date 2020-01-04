package main


import (
	"os"
	"log"
	"fmt"
	"flag"
	"bufio"
    "strings"
	"syscall"
	"github.com/spf13/cobra"
	"golang.org/x/crypto/ssh/terminal"
	"github.com/rogercoll/dirEncryptor/encrypter"
	"github.com/rogercoll/dirEncryptor/decrypter"
)

var mainCmd = &cobra.Command {
	Use: "canopus",
}

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
	var err error
	boolPtr := flag.Bool("d", false, "Decrypt a directory")
	flag.Parse()
	dir := readDir()
	fmt.Printf("Directory: %s\n", dir)
	pass := credentials()
	fmt.Printf("Password: %s\n", pass)
	if *boolPtr {
		err = decrypter.Decrypt(dir, pass)
		
	} else {
		err = encrypter.Encrypt(dir, pass)
	}
	if err != nil {
		log.Fatal(err)
	}
}