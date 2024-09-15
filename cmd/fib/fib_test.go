package main

import (
	"testing"
)

func TestCompute(t *testing.T) {
	result := compute(5, 3)
	if result != 19 {
		t.Fatal("for inputs month=3 and litter=3 expected result is 19")
	}
}

func TestCanParseInput(t *testing.T) {
	month, litter, err := parseInput("5 3")
	if err != nil {
		t.Fatalf("failed to parse input: %s", err)
	}

	if month != 5 || litter != 3 {
		t.Fatal("month must be equal to 5 and the litter to 3")
	}

	_, _, err = parseInput("50 3")
	if err == nil {
		t.Fatalf("the month parameter must be smaller or equal to 40")
	}
}
