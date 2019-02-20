package main

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"io"
	"io/ioutil"
	"log"
	"os"
	"fmt"
	"time"
	"sync"
)

var wg sync.WaitGroup
var key string

func encrypt(plaintext []byte) string {

	// Key
	keyb := []byte(key)

	// Create the AES cipher
	block, err := aes.NewCipher(keyb)
	if err != nil {
		panic(err)
	}
	ciphertext := make([]byte, aes.BlockSize+len(plaintext))

	// Slice of first 16 bytes
	iv := ciphertext[:aes.BlockSize]
	if _, err := io.ReadFull(rand.Reader, iv); err != nil {
		panic(err)
	}
	stream := cipher.NewCFBEncrypter(block, iv)
	stream.XORKeyStream(ciphertext[aes.BlockSize:], plaintext)

	return string(ciphertext)
}

func writeToFile(data string, file string, perm os.FileMode) {
	ioutil.WriteFile(file, []byte(data), perm)
}

func all_process(path string,file os.FileInfo){
	plaintext, err := ioutil.ReadFile(path)
	perms := file.Mode()
	if err != nil {
		panic(err.Error())
	}
	ciphertext := encrypt(plaintext)
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

func main() {
	if len(os.Args) != 3{
		log.Fatal("Invalid number of args: ./encrypt password /root/Desktop/")
	}
	key = os.Args[1]
	dir := os.Args[2]
	if len(key) != 16{
		log.Fatal("Key must be 16 characters(bytes) long, to prevent future decryption problems")
	}
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
	fmt.Printf("Total time %.2fs\n", finish)
}