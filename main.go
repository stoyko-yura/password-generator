package main

import (
	"bufio"
	"fmt"
	"math/rand"
	"os"
	"strconv"
	"strings"
)

type PasswordOptions struct {
	LowerCase bool
	UpperCase bool
	Digits    bool
	Symbols   bool
}

type PasswordConfig struct {
	Length  int
	Options PasswordOptions
}

var (
	lowercase = "abcdefghijklmnopqrstuvwxyz"
	uppercase = "ABCDEFGHIJKLMNOPQRSTUVWXYZ"
	digits    = "0123456789"
	symbols   = "!@#$%^&*()-_=+[]{};:,.<>?/|\\"
)

var defaultOptions = PasswordOptions{
	LowerCase: true,
	UpperCase: false,
	Digits:    false,
	Symbols:   false,
}

func printMenu(options PasswordOptions, length int) {
	fmt.Println("====== Password Generator ======")
	fmt.Printf("[1] Toggle lowercase  [%t]\n", options.LowerCase)
	fmt.Printf("[2] Toggle uppercase  [%t]\n", options.UpperCase)
	fmt.Printf("[3] Toggle digits     [%t]\n", options.Digits)
	fmt.Printf("[4] Toggle symbols    [%t]\n", options.Symbols)
	fmt.Printf("[5] Set length        [%d]\n", length)
	fmt.Println()
	fmt.Println("[6] Generate password")
	fmt.Println("================================")
	fmt.Print("Select an option: ")
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

	for i := 0; i < length; i++ {
		char := charset[rand.Intn(len(charset))]

		password[i] = char
	}

	return string(password)
}

func main() {
	passwordConfig := PasswordConfig{
		Length:  8,
		Options: defaultOptions,
	}

	for {
		printMenu(passwordConfig.Options, passwordConfig.Length)

		stdinReader := bufio.NewReader(os.Stdin)
		userInput, err := stdinReader.ReadString('\n')
		if err != nil {
			fmt.Println("Error reading input:", err)
			return
		}

		userInput = strings.TrimSpace(userInput)

		switch userInput {
		case "1":
			passwordConfig.Options.LowerCase = !passwordConfig.Options.LowerCase
		case "2":
			passwordConfig.Options.UpperCase = !passwordConfig.Options.UpperCase
		case "3":
			passwordConfig.Options.Digits = !passwordConfig.Options.Digits
		case "4":
			passwordConfig.Options.Symbols = !passwordConfig.Options.Symbols
		case "5":
			fmt.Println("Input password's length")

			userInput, _ = stdinReader.ReadString('\n')

			userPasswordLength, err := strconv.Atoi(strings.TrimSpace(userInput))
			if err != nil {
				fmt.Println("Error reading input:", err)
				return
			}

			passwordConfig.Length = userPasswordLength
		case "6":
			charset, err := generateCharset(passwordConfig.Options)
			if err != nil {
				fmt.Println("Error with charset:", err)
				continue
			}

			fmt.Printf("Generated password: %s\n", generatePassword(passwordConfig.Length, charset))

			fmt.Println("\nPress Enter to exit...")
			bufio.NewReader(os.Stdin).ReadBytes('\n')

			return
		default:
			fmt.Println("Unusual command")
		}
	}
}
