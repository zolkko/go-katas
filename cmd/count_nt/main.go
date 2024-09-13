package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"blumen.place/gokatas/utils"
)

func countNt(line string) (int, int, int, int) {
	var aCount, cCount, gCount, tCount int
	for _, chr := range line {
		switch chr {
		case 'a':
			aCount += 1
		case 't':
			tCount += 1
		case 'g':
			gCount += 1
		case 'c':
			cCount += 1
		default:
			fmt.Printf("unexpected character \"%v\"\n", chr)
			os.Exit(-1)
		}
	}
	return aCount, cCount, gCount, tCount
}

func prepareLine(line string) string {
	line = strings.TrimSuffix(line, "\n")
	line = strings.TrimSuffix(line, "\r")
	line = strings.ToLower(line)
	return line
}

func main() {
	rd := bufio.NewReader(os.Stdin)
	line, err := rd.ReadString('\n')
	if err != nil {
		fmt.Println(err)
		os.Exit(-1)
	}
	line = utils.TrimInput(line)

	aCount, cCount, gCount, tCount := countNt(line)

	fmt.Printf("%d %d %d %d\n", aCount, cCount, gCount, tCount)
}
