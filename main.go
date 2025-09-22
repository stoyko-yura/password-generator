package main

import (
	"bufio"
	"fmt"
	"math/rand"
	"os"
	"strconv"
	"strings"
)

var (
	lowercase = "abcdefghijklmnopqrstuvwxyz"
	uppercase = "ABCDEFGHIJKLMNOPQRSTUVWXYZ"
	digits    = "0123456789"
	symbols   = "!@#$%^&*()-_=+[]{};:,.<>?/|\\"
)

type PasswordOptions struct {
	LowerCase bool
	UpperCase bool
	Digits    bool
	Symbols   bool
}

var defaultOptions = PasswordOptions{
	LowerCase: true,
	UpperCase: false,
	Digits:    false,
	Symbols:   false,
}

type PasswordConfig struct {
	Length  int
	Options PasswordOptions
}

func generateCharset(passwordOptions PasswordOptions) (string, error) {
	charset := ""

	if passwordOptions.LowerCase {
		charset += lowercase
	}

	if passwordOptions.UpperCase {
		charset += uppercase
	}

	if passwordOptions.Digits {
		charset += digits
	}

	if passwordOptions.Symbols {
		charset += symbols
	}

	if len(charset) == 0 {
		return "", fmt.Errorf("no character sets selected")
	}

	return charset, nil
}

func generatePassword(length int, charset string) string {
	password := make([]byte, length)

	for i := range length {
		char := charset[rand.Intn(len(charset))]

		password[i] = char
	}

	return string(password)
}

func main() {
	passwordData := PasswordConfig{
		Length:  8,
		Options: defaultOptions,
	}

	for {
		fmt.Println(fmt.Sprintf("[1] - toggle lowercase(%t)\n"+
			"[2] - toggle uppercase(%t)\n"+
			"[3] - toggle digits(%t)\n"+
			"[4] - toggle symbols(%t)\n"+
			"[5] - change length(%d)\n\n"+
			"[6] - generate password",
			passwordData.Options.LowerCase, passwordData.Options.UpperCase, passwordData.Options.Digits, passwordData.Options.Symbols, passwordData.Length))

		stdinReader := bufio.NewReader(os.Stdin)
		userInput, err := stdinReader.ReadString('\n')
		if err != nil {
			fmt.Println("Error reading input:", err)
			return
		}

		userInput = strings.TrimSpace(userInput)

		switch userInput {
		case "1":
			passwordData.Options.LowerCase = !passwordData.Options.LowerCase
		case "2":
			passwordData.Options.UpperCase = !passwordData.Options.UpperCase
		case "3":
			passwordData.Options.Digits = !passwordData.Options.Digits
		case "4":
			passwordData.Options.Symbols = !passwordData.Options.Symbols
		case "5":
			fmt.Println("Input password's length")

			stdinReader = bufio.NewReader(os.Stdin)
			userInput, _ = stdinReader.ReadString('\n')

			userPasswordLength, err := strconv.Atoi(strings.TrimSpace(userInput))
			if err != nil {
				fmt.Println("Error reading input:", err)
				return
			}

			passwordData.Length = userPasswordLength
		case "6":
			charset, err := generateCharset(passwordData.Options)
			if err != nil {
				fmt.Println("Error with charset:", err)
				continue
			}

			fmt.Println("Generated password: " + generatePassword(passwordData.Length, charset))
			return
		default:
			fmt.Println("Unusual command")
		}
	}
}
