package main


import (
	"log"
	"github.com/spf13/cobra"
	"github.com/rogercoll/canopus/encrypter"
	"github.com/rogercoll/canopus/decrypter"
)

var rootCmd = &cobra.Command {
	Use: "canopus",
}

func init() {
	rootCmd.AddCommand(encrypter.EncryptCmd)
	rootCmd.AddCommand(decrypter.DecryptCmd)
}

func main() {
	if err := rootCmd.Execute(); err != nil {
		log.Fatal(err)
	}
}