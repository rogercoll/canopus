package decrypter

import (
	"crypto/aes"
	"crypto/cipher"
	"github.com/spf13/cobra"
	"io/ioutil"
	"os"
	"log"
	"fmt"
	"time"
	"sync"
	"github.com/rogercoll/canopus/utils"
)

var wg sync.WaitGroup
var key string

var DecryptCmd = &cobra.Command {
	Use:	"decrypt",
	Run:	func(cmd *cobra.Command, args []string) {
		pass := utils.GetCredentials()
		dir := utils.GetDir()
		err := Decrypt(dir, pass)
		if err != nil {
			log.Fatal(err)
		}
	},
}

func Decrypt(dir, password string) error {
	key = password
	allfiles := make([]string,0,200)
	getPaths(&allfiles,dir)
	start := time.Now()	
	
	for _,file := range allfiles{
		wg.Add(1)
		estat,err := os.Stat(file)
		if err != nil{
			panic(err)
		}
		go all_process(file,estat)
	}
	wg.Wait()
	finish := time.Since(start).Seconds()
	fmt.Printf("Total dectypton time %.2fs\n", finish)
	return nil
}


func makeDecryption(ciphertext []byte) string {
	// Key
	keyb := []byte(key)

	// Create the AES cipher
	block, err := aes.NewCipher(keyb)
	if err != nil {
		panic(err)
	}
	// Before even testing the decryption,
	// if the text is too small, then it is incorrect
	if len(ciphertext) < aes.BlockSize {
		panic("Text is too short")
	}
	// Get the 16 byte IV
	iv := ciphertext[:aes.BlockSize]

	// Remove the IV from the ciphertext
	ciphertext = ciphertext[aes.BlockSize:]

	// Return a decrypted stream
	stream := cipher.NewCFBDecrypter(block, iv)

	// Decrypt bytes from ciphertext
	stream.XORKeyStream(ciphertext, ciphertext)

	return string(ciphertext)
}

func writeToFile(data string, file string, perm os.FileMode) {
	ioutil.WriteFile(file, []byte(data), 777)
}

func all_process(path string,file os.FileInfo){
	plaintext, err := ioutil.ReadFile(path)
	perms := file.Mode()
	if err != nil {
		panic(err.Error())
	}
	ciphertext := makeDecryption(plaintext)
	writeToFile(ciphertext, path,perms)
	wg.Done()
}

func getPaths(allfiles *[]string, path string){
	dir, err := ioutil.ReadDir(path)
	if err != nil{
		log.Fatal(err)
	}
	for _,file := range dir{
		if !file.IsDir(){
			*allfiles = append(*allfiles,path+file.Name())
		}else{
			if file.Name() != "." && file.Name() != ".."{
				getPaths(allfiles,path+file.Name()+"/")
			}
		}
	}
}
