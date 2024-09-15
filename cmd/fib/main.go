package main

import (
	"bufio"
	"errors"
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

func parseInput(input string) (int, int, error) {
	line := utils.TrimInput(input)

	chunks := strings.Split(line, " ")
	if len(chunks) != 2 {
		return 0, 0, errors.New("failed to parse input, two numbers expected")
	}

	month, err := strconv.Atoi(chunks[0])
	if err != nil {
		return 0, 0, fmt.Errorf("failed to parse the month parameter: %w", err)
	}
	if month > 40 {
		return 0, 0, errors.New("month must be smaller or equal 40")
	}

	litterSize, err := strconv.Atoi(chunks[1])
	if err != nil {
		return 0, 0, fmt.Errorf("failed to parse litter parameter: %w", err)
	}

	return month, litterSize, nil
}

func main() {
	rdr := bufio.NewReader(os.Stdin)
	line, err := rdr.ReadString('\n')
	if err != nil {
		fmt.Println(err)
		os.Exit(-1)
	}

	month, litterSize, err := parseInput(line)
	if err != nil {
		fmt.Println(err)
		os.Exit(-1)
	}

	result := compute(month, litterSize)
	fmt.Println(result)
}
