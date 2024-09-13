package main

import (
	"bufio"
	"errors"
	"fmt"
	"os"

	"blumen.place/gokatas/utils"
)

func reverse(dna string) (string, error) {
	dnaLen := len(dna)
	result := make([]rune, dnaLen)

	offset := dnaLen - 1
	for i, chr := range dna {
		insertionIndex := offset - i

		var inversedChr rune
		switch chr {
		case 'a':
			inversedChr = 't'
		case 't':
			inversedChr = 'a'
		case 'g':
			inversedChr = 'c'
		case 'c':
			inversedChr = 'g'
		default:
			return "", errors.New(fmt.Sprintf("unexpected character in the DNA string: %s", string(chr)))
		}

		result[insertionIndex] = inversedChr
	}

	return string(result), nil
}

func main() {
	rdr := bufio.NewReader(os.Stdin)
	line, err := rdr.ReadString('\n')
	if err != nil {
		fmt.Println(err)
		os.Exit(-1)
	}

	dna := utils.TrimInput(line)

	reversedDna, err := reverse(dna)
	if err != nil {
		fmt.Println(err)
		os.Exit(-1)
	}
	fmt.Println(reversedDna)
}
