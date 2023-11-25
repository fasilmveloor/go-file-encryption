package main

import (
	"bytes"
	"fmt"
	"os"
	"golang.org/x/term"
	"github.com/fasilmveloor/go-file-encryption/filecrypt"
)

func main() {
	if len(os.Args) < 2 {
		printHelp()
		os.Exit(0)
	}
	function := os.Args[1]

	switch function {
		case "help":
			printHelp()
		case "encrypt":
			encryptHandle()
		case "decrypt":
			decryptHandle()
		default:
			fmt.Println("Run encrypt to encrypt a file, and decrypt to decrypt a file.")
			os.Exit(1)
	}
}

func printHelp() {
	fmt.Println("File Encryption")
	fmt.Println("Simple file encrypter and decrypter for your day-to-day needs.")
	fmt.Println("")
	fmt.Println("Usage:")
	fmt.Println("")
	fmt.Println("\t go run . encrypt /path/to/your/file")
	fmt.Println("")
	fmt.Println("Commands:")
	fmt.Println("")
	fmt.Println("\t encrypt \t Encrypts a file given a password")
	fmt.Println("\t decrypt \t Tries to decrypt a file using a password")
	fmt.Println("\t help \t\t Prints help text")
	fmt.Println("")
}

func encryptHandle() {
	if len(os.Args) < 3 {
		println("missing the path to the file, For more info run 'go run . help'")
		os.Exit(0)
	}
	file := os.Args[2]
	if !validateFile(file) {
		panic("File not found")
	}
	password := getPassword()
	fmt.Println("\nEncrypting...")
	filecrypt.Encrypt(file, password)
	fmt.Println("\n file successfully protected")
}

func decryptHandle() {
	if len(os.Args) < 3 {
		println("missing the path to the file, For more info run 'go run . help'")
		os.Exit(0)
	}
	file := os.Args[2]
	if !validateFile(file) {
		panic("File not found")
	}
	fmt.Print("Enter Password:")
	password, _ := term.ReadPassword(0)
	fmt.Println("\nDecrypting...")
	filecrypt.Decrypt(file, password)
	fmt.Println("\n file successfully decrypted")
}

func getPassword() []byte {
	fmt.Print("Enter your password: ")
	password, _ := term.ReadPassword(0)
	fmt.Print("\nConfirm Password: ")
	confirmPassword, _ := term.ReadPassword(0)
	if !validatePassword(password, confirmPassword) {
		fmt.Println("\nPasswords do not match, try again")
		return getPassword()
	}
	return password

}

func validatePassword(password []byte, confirmPassword []byte) bool {
	return bytes.Equal(password, confirmPassword)
}

func validateFile(file string) bool {
	if _, err := os.Stat(file); os.IsNotExist(err) {
		return false
	}
	return true
}