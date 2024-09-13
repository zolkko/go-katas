package main

import (
	"bufio"
	"errors"
	"fmt"
	"os"

	"blumen.place/gokatas/utils"
)

func translate(in string) (string, error) {
	result := make([]rune, 0, len(in))
	for _, chr := range in {
		switch chr {
		case 'a':
			result = append(result, 'a')
		case 't':
			result = append(result, 'u')
		case 'g':
			result = append(result, 'g')
		case 'c':
			result = append(result, 'c')
		default:
			return "", errors.New(fmt.Sprintf("unexpected symbol in DNA sequence: %s", string(chr)))
		}
	}
	return string(result), nil
}

func main() {
	rdr := bufio.NewReader(os.Stdin)
	in, err := rdr.ReadString('\n')
	if err != nil {
		fmt.Println(err)
		os.Exit(-1)
	}

	dna := utils.TrimInput(in)

	rna, err := translate(dna)
	if err != nil {
		fmt.Println(err)
		os.Exit(-1)
	}

	fmt.Println(rna)
}
