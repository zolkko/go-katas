package main

import (
	"bufio"
	"math"
	"strings"
	"testing"
)

func TestGcContent(t *testing.T) {
	data := `>Rosalind_6404
CCTGCGGAAGATCGGCACTAGAATAGCCAGAACCGTTTCTCTGAGGCTTCCGGCCTTCCC
TCCCACTAATAATTCTGAGG
>Rosalind_5959
CCATCGGTAGCGCATCCTTAGTCCAATTAAGTCCCTATCCAGGCGCTCCGCCGAAGGTCT
ATATCCATTTGTCAGCAGACACGC
>Rosalind_0808
CCACCCTCGTGGTATGGCTAGGCATTCAGGAACCGGAGAACGCTTCAGACCAGCCCGGAC
TGGGAACCTGCGGGCAGTAGGTGGAAT`

	rdr := strings.NewReader(data)
	scanner := bufio.NewScanner(rdr)

	id, gc := compute(scanner)

	if id != "Rosalind_0808" {
		t.Fatalf("expected Rosalind_0808, got %s", id)
	}

	const expected = 60.919540
	const delta = 0.01
	if math.Abs(gc-expected) > delta {
		t.Fatalf("expected %f +-%f GC content, got %f", expected, delta, gc)
	}
}
