package main

import (
	"bufio"
	"fmt"
	"os"
)

const KeyLength = 16
const PlainTextLength = 16
const TotalRounds int = 10
const TotalKeys int = TotalRounds + 1

type Word [4]byte
type WordBlock [4]Word

var SubstituteBox = [][]byte{
	{0x63, 0x7c, 0x77, 0x7b, 0xf2, 0x6b, 0x6f, 0xc5, 0x30, 0x01, 0x67, 0x2b, 0xfe, 0xd7, 0xab, 0x76},
	{0xca, 0x82, 0xc9, 0x7d, 0xfa, 0x59, 0x47, 0xf0, 0xad, 0xd4, 0xa2, 0xaf, 0x9c, 0xa4, 0x72, 0xc0},
	{0xb7, 0xfd, 0x93, 0x26, 0x36, 0x3f, 0xf7, 0xcc, 0x34, 0xa5, 0xe5, 0xf1, 0x71, 0xd8, 0x31, 0x15},
	{0x04, 0xc7, 0x23, 0xc3, 0x18, 0x96, 0x05, 0x9a, 0x07, 0x12, 0x80, 0xe2, 0xeb, 0x27, 0xb2, 0x75},
	{0x09, 0x83, 0x2c, 0x1a, 0x1b, 0x6e, 0x5a, 0xa0, 0x52, 0x3b, 0xd6, 0xb3, 0x29, 0xe3, 0x2f, 0x84},
	{0x53, 0xd1, 0x00, 0xed, 0x20, 0xfc, 0xb1, 0x5b, 0x6a, 0xcb, 0xbe, 0x39, 0x4a, 0x4c, 0x58, 0xcf},
	{0xd0, 0xef, 0xaa, 0xfb, 0x43, 0x4d, 0x33, 0x85, 0x45, 0xf9, 0x02, 0x7f, 0x50, 0x3c, 0x9f, 0xa8},
	{0x51, 0xa3, 0x40, 0x8f, 0x92, 0x9d, 0x38, 0xf5, 0xbc, 0xb6, 0xda, 0x21, 0x10, 0xff, 0xf3, 0xd2},
	{0xcd, 0x0c, 0x13, 0xec, 0x5f, 0x97, 0x44, 0x17, 0xc4, 0xa7, 0x7e, 0x3d, 0x64, 0x5d, 0x19, 0x73},
	{0x60, 0x81, 0x4f, 0xdc, 0x22, 0x2a, 0x90, 0x88, 0x46, 0xee, 0xb8, 0x14, 0xde, 0x5e, 0x0b, 0xdb},
	{0xe0, 0x32, 0x3a, 0x0a, 0x49, 0x06, 0x24, 0x5c, 0xc2, 0xd3, 0xac, 0x62, 0x91, 0x95, 0xe4, 0x79},
	{0xe7, 0xc8, 0x37, 0x6d, 0x8d, 0xd5, 0x4e, 0xa9, 0x6c, 0x56, 0xf4, 0xea, 0x65, 0x7a, 0xae, 0x08},
	{0xba, 0x78, 0x25, 0x2e, 0x1c, 0xa6, 0xb4, 0xc6, 0xe8, 0xdd, 0x74, 0x1f, 0x4b, 0xbd, 0x8b, 0x8a},
	{0x70, 0x3e, 0xb5, 0x66, 0x48, 0x03, 0xf6, 0x0e, 0x61, 0x35, 0x57, 0xb9, 0x86, 0xc1, 0x1d, 0x9e},
	{0xe1, 0xf8, 0x98, 0x11, 0x69, 0xd9, 0x8e, 0x94, 0x9b, 0x1e, 0x87, 0xe9, 0xce, 0x55, 0x28, 0xdf},
	{0x8c, 0xa1, 0x89, 0x0d, 0xbf, 0xe6, 0x42, 0x68, 0x41, 0x99, 0x2d, 0x0f, 0xb0, 0x54, 0xbb, 0x16},
}

var RoundConstants = [TotalRounds]Word{
	{0x01, 0x00, 0x00, 0x00},
	{0x02, 0x00, 0x00, 0x00},
	{0x04, 0x00, 0x00, 0x00},
	{0x08, 0x00, 0x00, 0x00},
	{0x10, 0x00, 0x00, 0x00},
	{0x20, 0x00, 0x00, 0x00},
	{0x40, 0x00, 0x00, 0x00},
	{0x80, 0x00, 0x00, 0x00},
	{0x1B, 0x00, 0x00, 0x00},
	{0x36, 0x00, 0x00, 0x00},
}

func main() {
	fmt.Println("============== AES-128 ==============")

	key := readKey()
	plainText := readPlainText()

	fmt.Print("\nKey        : ")
	printArrayAsFormat("%002X ", key[0:])

	fmt.Print("\nPlainText  : ")
	printArrayAsFormat("%002X ", plainText[0:])

	fmt.Println()

	fmt.Println("\nKey Expansion:")
	roundKeys := expandRoundKeys(key)
	for r, key := range roundKeys {
		if r == 0 {
			fmt.Println("- Initial Key:")
		} else {
			fmt.Println("- Round", r, "Key")
		}
		printBlock(key, "  ")
	}

	fmt.Println("Initial State :")
	state := mapToWordBlock(plainText)
	printBlock(state, "")

	fmt.Println("After Initial AddRoundKey:")
	state = addRoundKey(state, roundKeys[0])
	printBlock(state, "")

	for r := 1; r <= TotalRounds; r++ {
		fmt.Println("==============", "Round", r, "==============")

		fmt.Printf("- [R%d] After SubstituteBytes:\n", r)
		state = substituteBytes(state)
		printBlock(state, "  ")

		fmt.Printf("- [R%d] After ShiftRows:\n", r)
		state = shiftRows(state)
		printBlock(state, "  ")

		if r < TotalRounds {
			fmt.Printf("- [R%d] After MixColumns:\n", r)
			state = mixColumns(state)
			printBlock(state, "  ")
		}

		fmt.Printf("- [R%d] After AddRoundKey:\n", r)
		state = addRoundKey(state, roundKeys[r])
		printBlock(state, "  ")
	}

	cipherText := mapToArray(state)
	fmt.Print("CipherText  : ")
	printArrayAsFormat("%002X ", cipherText[0:])

	scanner := bufio.NewScanner(os.Stdin)
	fmt.Print("\n\nPress enter...")
	scanner.Scan()
}

func readKey() [KeyLength]byte {
	scanner := bufio.NewScanner(os.Stdin)

	keyText := ""
	for keyText == "" {
		fmt.Print("Enter Key (16 characters): ")
		scanner.Scan()
		keyText = scanner.Text()

		if len(keyText) != 16 {
			keyText = ""
			fmt.Println("Key must be 16 characters (128 bit)")
		}
	}

	var key [KeyLength]byte
	for i, val := range keyText {
		key[i] = byte(val)
	}

	return key
}

func readPlainText() [PlainTextLength]byte {
	scanner := bufio.NewScanner(os.Stdin)

	plainText := ""
	for plainText == "" {
		fmt.Print("Enter PlainText (max 16 characters): ")
		scanner.Scan()
		plainText = scanner.Text()

		if len(plainText) >= 16 {
			plainText = ""
			fmt.Println("PlainText max is 16 characters (128 bit)")
		}
	}
	var plain [PlainTextLength]byte
	for i, val := range plainText {
		plain[i] = byte(val)
	}

	return plain
}

func addRoundKey(state WordBlock, roundKey WordBlock) WordBlock {
	var newState WordBlock
	for i := 0; i < 4; i++ {
		newState[i] = exclusiveOr(state[i], roundKey[i])
	}
	return newState
}

func substituteBytes(state WordBlock) WordBlock {
	var newState WordBlock
	for i := 0; i < 4; i++ {
		newState[i] = substituteWord(state[i])
	}
	return newState
}

func shiftRows(state WordBlock) WordBlock {
	var newState WordBlock
	newState[0][0] = state[0][0]
	newState[1][0] = state[1][0]
	newState[2][0] = state[2][0]
	newState[3][0] = state[3][0]

	newState[0][1] = state[1][1]
	newState[1][1] = state[2][1]
	newState[2][1] = state[3][1]
	newState[3][1] = state[0][1]

	newState[0][2] = state[2][2]
	newState[1][2] = state[3][2]
	newState[2][2] = state[0][2]
	newState[3][2] = state[1][2]

	newState[0][3] = state[3][3]
	newState[1][3] = state[0][3]
	newState[2][3] = state[1][3]
	newState[3][3] = state[2][3]
	return newState
}

func mixColumns(s WordBlock) WordBlock {
	var newState WordBlock
	for w := 0; w < 4; w++ {
		newState[w][0] = gMul(0x02, s[w][0]) ^ gMul(0x03, s[w][1]) ^ s[w][2] ^ s[w][3]
		newState[w][1] = s[w][0] ^ gMul(0x02, s[w][1]) ^ gMul(0x03, s[w][2]) ^ s[w][3]
		newState[w][2] = s[w][0] ^ s[w][1] ^ gMul(0x02, s[w][2]) ^ gMul(0x03, s[w][3])
		newState[w][3] = gMul(0x03, s[w][0]) ^ s[w][1] ^ s[w][2] ^ gMul(0x02, s[w][3])
	}

	return newState
}

// Galois Field (256) Multiplication of two Bytes
func gMul(a byte, b byte) byte {
	var p byte = 0

	for counter := 0; counter < 8; counter++ {
		if (b & 1) != 0 {
			p ^= a
		}

		var hiBitSet bool = (a & 0x80) != 0
		a <<= 1
		if hiBitSet {
			a ^= 0x1B /* x^8 + x^4 + x^3 + x + 1 */
		}
		b >>= 1
	}

	return p
}

func expandRoundKeys(key [KeyLength]byte) [TotalKeys]WordBlock {
	var roundKeys [TotalKeys]WordBlock

	roundKeys[0] = mapToWordBlock(key)
	for r := 1; r <= 10; r++ {
		roundKeys[r][0] = exclusiveOr(roundKeys[r-1][0], exclusiveOr(substituteWord(rotateWord(roundKeys[r-1][3])), RoundConstants[r-1]))
		roundKeys[r][1] = exclusiveOr(roundKeys[r-1][1], roundKeys[r][0])
		roundKeys[r][2] = exclusiveOr(roundKeys[r-1][2], roundKeys[r][1])
		roundKeys[r][3] = exclusiveOr(roundKeys[r-1][3], roundKeys[r][2])
	}

	return roundKeys
}

func rotateWord(word Word) Word {
	var newWord Word
	newWord[0] = word[1]
	newWord[1] = word[2]
	newWord[2] = word[3]
	newWord[3] = word[0]

	return newWord
}

func substituteWord(word Word) Word {
	var newWord Word
	for i := 0; i < 4; i++ {
		x := word[i] >> 4
		y := word[i] & 0b00001111
		newWord[i] = SubstituteBox[x][y]
	}
	return newWord
}

func exclusiveOr(word1 Word, word2 Word) Word {
	var newWord Word
	for i := 0; i < 4; i++ {
		newWord[i] = word1[i] ^ word2[i]
	}
	return newWord
}

func mapToWordBlock(key [16]byte) WordBlock {
	var block WordBlock
	for i, val := range key {
		block[i/4][i%4] = val
	}

	return block
}

func mapToArray(block WordBlock) [16]byte {
	var arr [16]byte
	for w, word := range block {
		arr[w*4+0] = word[0]
		arr[w*4+1] = word[1]
		arr[w*4+2] = word[2]
		arr[w*4+3] = word[3]
	}
	return arr
}

func printBlock(block WordBlock, prefix string) {
	fmt.Printf(prefix+"%002X %002X %002X %002X\n", block[0][0], block[1][0], block[2][0], block[3][0])
	fmt.Printf(prefix+"%002X %002X %002X %002X\n", block[0][1], block[1][1], block[2][1], block[3][1])
	fmt.Printf(prefix+"%002X %002X %002X %002X\n", block[0][2], block[1][2], block[2][2], block[3][2])
	fmt.Printf(prefix+"%002X %002X %002X %002X\n", block[0][3], block[1][3], block[2][3], block[3][3])
	fmt.Println()
}

func printArrayAsFormat(format string, data []byte) {
	for _, val := range data {
		fmt.Printf(format, val)
	}
}
