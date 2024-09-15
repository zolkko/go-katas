package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"blumen.place/gokatas/utils"
)

func gcContent(value string) float64 {
	cnt := 0
	for _, chr := range value {
		switch chr {
		case 'g', 'c':
			cnt += 1
		case 'a', 't':
		default:
			panic("unexpected base")
		}
	}
	return (float64(cnt) / float64(len(value))) * 100.0
}

func compute(rdr *bufio.Scanner) (string, float64) {
	maxId := ""
	maxGc := 0.0

	lastId := ""
	fasta := ""

	innerGc := func() {
		if lastId != "" {
			lastGc := gcContent(fasta)
			if lastGc > maxGc {
				maxId = lastId
				maxGc = lastGc
			}
		}
	}

	for rdr.Scan() {
		line := rdr.Text()

		if newId, ok := strings.CutPrefix(line, ">"); ok {
			innerGc()

			lastId = newId
			fasta = ""
		} else {
			fasta += utils.TrimInput(line)
		}
	}

	innerGc()

	return maxId, maxGc
}

func main() {
	rdr := bufio.NewScanner(os.Stdin)

	maxId, maxGc := compute(rdr)

	fmt.Println(maxId)
	fmt.Println(maxGc)
}
