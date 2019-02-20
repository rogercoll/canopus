# Directory Encryptor/Decryptor
Golang program that encrypts any type of file given a directory. 
PASSWORD MUST BE 16 CHARACTERS LONG

## Usage
### 1. Clone this module:
```sh
$ git clone https://github.com/rogercoll/dirEncryptor.git
```
### 2. Building it:
```sh
$ cd dirEncryptor/encrypter
$ go build encrypter.go
$ cd dirEncryptor/decrypter
$ go build decrypter.go
```
### 3. Run it:
```golang
./encrypter halohalohalohalo /root/Desktop/
./decrypter halohalohalohalo /root/Desktop/
```
