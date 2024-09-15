package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"

	"blumen.place/gokatas/utils"
)

func compute(month, litterSize int) int {
	// Fn - represents the number of rabbit pairs alive after the n-th month
	// F(n−1) - previous month;
	// F(n−2) - number of offspring in any month is equal to the number of rabbits that were alive two months prior;
	//
	// Rabbits were alive the previous month, plus any new offspring
	// 	Fn = F(n − 1) + F(n − 2)
	switch month {
	case 1:
		offspring := 1
		prevMonth := 1
		return offspring + prevMonth
	case 2:
		offspring := 1 * litterSize
		prevMonth := compute(month-1, litterSize)
		return offspring + prevMonth
	default:
		offspring := compute(month-2, litterSize)
		prevMonth := compute(month-1, litterSize)
		return offspring + prevMonth
	}
}

func main() {
	rdr := bufio.NewReader(os.Stdin)
	line, err := rdr.ReadString('\n')
	if err != nil {
		fmt.Println(err)
		os.Exit(-1)
	}

	line = utils.TrimInput(line)

	chunks := strings.Split(line, " ")
	if len(chunks) != 2 {
		fmt.Println("two numbers expected")
		os.Exit(-1)
	}

	month, err := strconv.Atoi(chunks[0])
	if err != nil {
		fmt.Println(err)
		os.Exit(-1)
	}
	if month > 40 {
		fmt.Println("month must be smaller or equal 40")
		os.Exit(-1)
	}

	litterSize, err := strconv.Atoi(chunks[1])
	if err != nil {
		fmt.Println(err)
		os.Exit(-1)
	}

	result := compute(month, litterSize)
	fmt.Println(result)
}
