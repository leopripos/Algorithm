package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
	"unicode"
)

var CaesarEncytp = "1"
var CaesarDecrypt = "2"

var VigenereEncytp = "3"
var VigenereDecrypt = "4"

var TranspositionEncytp = "5"
var TranspositionDecrypt = "6"

var ExitApplication = "0"

func main() {
	for {
		fmt.Println()
		menu := readMainMenu()
		switch menu {
		case CaesarEncytp:
			cipher := "Caesar"
			plainText := readPlainText(cipher)
			key := readIntKey(cipher)
			cipherText := caesarEncrypt(plainText, key)
			fmt.Println("Cipher\t\t:", cipher)
			fmt.Println("PlainText\t:", plainText)
			fmt.Println("Key\t\t:", key)
			fmt.Println("Result CipherText:", cipherText)
		case CaesarDecrypt:
			cipher := "Caesar"
			cipherText := readCipherText(cipher)
			key := readIntKey(cipher)
			plainText := caesarDecrypt(cipherText, key)
			fmt.Println("Cipher\t\t:", cipher)
			fmt.Println("CipherText\t:", cipherText)
			fmt.Println("Key\t\t:", key)
			fmt.Println("Caesar PlainText\t:", plainText)
		case VigenereEncytp:
			cipher := "Vigenere"
			plainText := readPlainText(cipher)
			key := readStringKey(cipher)
			cipherText := vigenereEncrypt(plainText, key)
			fmt.Println("Cipher\t\t:", cipher)
			fmt.Println("PlainText\t:", plainText)
			fmt.Println("Key\t\t:", key)
			fmt.Println("Vigenere CipherText:", cipherText)
		case VigenereDecrypt:
			cipher := "Vigenere"
			cipherText := readCipherText(cipher)
			key := readStringKey(cipher)
			plainText := vigenereDecrypt(cipherText, key)
			fmt.Println("Cipher\t\t:", cipher)
			fmt.Println("CipherText\t:", cipherText)
			fmt.Println("Key\t\t:", key)
			fmt.Println("Vigenere PlainText\t:", plainText)
		case TranspositionEncytp:
			cipher := "Transposition"
			plainText := readPlainText(cipher)
			key := readIntKey(cipher)
			cipherText := transpositionEncrypt(plainText, key)
			fmt.Println("Cipher\t\t:", cipher)
			fmt.Println("PlainText\t:", plainText)
			fmt.Println("Key\t\t:", key)
			fmt.Println("Transposition CipherText:", cipherText)
		case TranspositionDecrypt:
			cipher := "Transposition"
			cipherText := readCipherText(cipher)
			key := readIntKey(cipher)
			plainText := transpositionDecrypt(cipherText, key)
			fmt.Println("Cipher\t\t:", cipher)
			fmt.Println("CipherText\t:", cipherText)
			fmt.Println("Key\t\t:", key)
			fmt.Println("Transposition PlainText\t:", plainText)
		case ExitApplication:
			os.Exit(0)
		default:
			continue
		}

		scanner := bufio.NewScanner(os.Stdin)
		fmt.Print("\nPress enter...")
		scanner.Scan()
	}
}

func caesarEncrypt(plainText string, key int) string {
	plainText = strings.ReplaceAll(plainText, " ", "")
	key = key % 26
	return strings.Map(func(c rune) rune {
		if c >= 97 {
			return rune((int(c)+key-97)%26 + 97)
		}
		return rune((int(c)+key-65)%26 + 65)
	}, plainText)
}

func caesarDecrypt(cipherText string, key int) string {
	cipherText = strings.ReplaceAll(cipherText, " ", "")
	key = key % 26
	return strings.Map(func(c rune) rune {
		if c >= 97 {
			return rune((int(c)-key+26-97)%26 + 97)
		}
		return rune((int(c)-key+26-65)%26 + 65)
	}, cipherText)
}

func vigenereEncrypt(plainText string, key string) string {
	plainText = strings.ReplaceAll(plainText, " ", "")
	keyLength := len(key)
	cipherText := ""
	for i, c := range plainText {
		currentKey := int(key[i%keyLength])
		if currentKey >= 97 {
			currentKey -= 96
		} else {
			currentKey -= 64
		}
		if c >= 97 {
			cipherText += string(rune((int(c)+currentKey-97)%26 + 97))
		} else {
			cipherText += string(rune((int(c)+currentKey-65)%26 + 65))
		}
	}

	return cipherText
}

func vigenereDecrypt(cipherText string, key string) string {
	cipherText = strings.ReplaceAll(cipherText, " ", "")
	keyLength := len(key)
	plainText := ""
	for i, c := range cipherText {
		currentKey := int(key[i%keyLength])
		if currentKey >= 97 {
			currentKey -= 96
		} else {
			currentKey -= 64
		}
		if c >= 97 {
			plainText += string(rune((int(c)-currentKey+26-97)%26 + 97))
		} else {
			plainText += string(rune((int(c)-currentKey+26-65)%26 + 65))
		}
	}

	return plainText
}

func transpositionEncrypt(plainText string, key int) string {
	plainText = strings.ReplaceAll(plainText, " ", "")
	plainTextLength := len(plainText)
	totalColumn := key
	totalRow := int(math.Ceil(float64(plainTextLength) / float64(key)))
	cipherText := ""

	for c := 0; c < totalColumn; c++ {
		for r := 0; r < totalRow; r++ {
			index := c + r*totalColumn
			if index < plainTextLength {
				cipherText += string(plainText[index])
			} else {
				continue
			}
		}
	}

	return cipherText
}

func transpositionDecrypt(cipherText string, key int) string {
	cipherText = strings.ReplaceAll(cipherText, " ", "")
	cipherTextLength := len(cipherText)
	totalRow := key
	totalColumn := int(math.Ceil(float64(cipherTextLength) / float64(key)))
	totalEmptyCell := (totalColumn * totalRow) - cipherTextLength
	firstRowWithEmptyCell := totalRow - totalEmptyCell
	plainText := ""

	for c := 0; c < totalColumn; c++ {
		for r := 0; r < totalRow; r++ {
			index := c + r*totalColumn
			if r >= firstRowWithEmptyCell {
				if c == totalColumn-1 {
					break
				} else {
					index -= (r - firstRowWithEmptyCell)
				}
			}

			plainText += string(cipherText[index])
		}
	}

	return plainText
}

func readMainMenu() string {
	scanner := bufio.NewScanner(os.Stdin)
	fmt.Println("Choose cryptography!")
	fmt.Println("1. Encrypt Caesar")
	fmt.Println("2. Decrypt Caesar")
	fmt.Println("3. Encrypt Vigenere")
	fmt.Println("4. Decrypt Vigenere")
	fmt.Println("5. Encrypt Transposition")
	fmt.Println("6. Decrypt Transposition")
	fmt.Println()
	fmt.Println("0. Exit")
	fmt.Println()
	pilihan := ""
	for pilihan != CaesarEncytp &&
		pilihan != CaesarDecrypt &&
		pilihan != VigenereEncytp &&
		pilihan != VigenereDecrypt &&
		pilihan != TranspositionEncytp &&
		pilihan != TranspositionDecrypt &&
		pilihan != ExitApplication {
		if pilihan != "" {
			fmt.Print("Input is not valid, ")
		}
		fmt.Print("Input option: ")
		scanner.Scan()
		pilihan = scanner.Text()
	}

	return pilihan
}

func isLetterOrSpace(s string) bool {
	for _, r := range s {
		if !unicode.IsLetter(r) && r != ' ' {
			return false
		}
	}
	return true
}

func isLetterOnly(s string) bool {
	for _, r := range s {
		if !unicode.IsLetter(r) {
			return false
		}
	}
	return true
}

func readPlainText(cipher string) string {
	scanner := bufio.NewScanner(os.Stdin)
	text := ""
	for text == "" {
		fmt.Print(cipher, " is chosen, Enter Plain Text: ")
		scanner.Scan()
		text = scanner.Text()
		if !isLetterOrSpace(text) {
			text = ""
			fmt.Println("Only a-z/A-Z or space are allowed.")
		}
	}

	return text
}

func readCipherText(cipher string) string {
	scanner := bufio.NewScanner(os.Stdin)
	text := ""
	for text == "" {
		fmt.Print(cipher, " is chosen, Enter Cipher Text: ")
		scanner.Scan()
		text = scanner.Text()
		if !isLetterOnly(text) {
			text = ""
			fmt.Println("Only a-z/A-Z are allowed.")
		}
	}

	return text
}

func readIntKey(cipher string) int {
	scanner := bufio.NewScanner(os.Stdin)
	key := -1
	var err error
	for key == -1 {
		fmt.Print("Enter ", cipher, " key (number > 0): ")
		scanner.Scan()
		text := scanner.Text()
		key, err = strconv.Atoi(text)
		if err != nil || key <= 0 {
			key = -1
		}
	}

	return key
}

func readStringKey(cipher string) string {
	scanner := bufio.NewScanner(os.Stdin)

	key := ""
	for key == "" {
		fmt.Print("Enter ", cipher, " key (string): ")
		scanner.Scan()
		key = scanner.Text()
		if !isLetterOnly(key) {
			key = ""
			fmt.Println("Only a-z/A-Z are allowed.")
		}
	}

	return key
}
