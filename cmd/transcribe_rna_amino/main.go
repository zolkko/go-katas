package main

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"strings"

	"blumen.place/gokatas/utils"
)

func prepareString(line string) string {
	line = strings.TrimSpace(line)
	line = strings.ToLower(line)
	return line
}

type codon string

type amino int

const (
	unknown amino = iota
	stop
	phe
	leu
	ser
	tyr
	cys
	trp
	pro
	his
	gln
	arg
	lle
	met
	thr
	asn
	lys
	val
	ala
	asp
	glu
	gly
)

func (a amino) String() string {
	switch a {
	case unknown:
		return "Unknown"
	case stop:
		return "Stop"
	case phe:
		return "Phe"
	case leu:
		return "Leu"
	case ser:
		return "Ser"
	case tyr:
		return "Tyr"
	case cys:
		return "Cys"
	case trp:
		return "Trp"
	case pro:
		return "Pro"
	case his:
		return "His"
	case gln:
		return "Gln"
	case arg:
		return "Arg"
	case lle:
		return "Lle"
	case met:
		return "Met"
	case thr:
		return "Thr"
	case asn:
		return "Asn"
	case lys:
		return "Lys"
	case val:
		return "Val"
	case ala:
		return "Ala"
	case asp:
		return "Asp"
	case glu:
		return "Glu"
	case gly:
		return "Gly"
	default:
		msg := fmt.Sprintf("got unknown amino acid code: %d", a)
		panic(msg)
	}
}

var CODON_MAP = map[rna]amino{
	"uuu": phe,
	"uuc": phe,
	"uua": leu,
	"uug": leu,
	//
	"ucu": ser,
	"ucc": ser,
	"uca": ser,
	"ucg": ser,
	//
	"uau": tyr,
	"uac": tyr,
	"uaa": stop,
	"uag": stop,
	//
	"ugu": cys,
	"ugc": cys,
	"uga": stop,
	"ugg": trp,
	//
	"cuu": leu,
	"cuc": leu,
	"cua": leu,
	"cug": leu,
	//
	"ccu": pro,
	"ccc": pro,
	"cca": pro,
	"ccg": pro,
	//
	"cau": his,
	"cac": his,
	"caa": gln,
	"cag": gln,
	//
	"cgu": arg,
	"cga": arg,
	"cgg": arg,
	"cgc": arg,
	//
	"auu": lle,
	"auc": lle,
	"aua": lle,
	"aug": met,
	//
	"acu": thr,
	"acc": thr,
	"aca": thr,
	"acg": thr,
	//
	"aau": asn,
	"aac": asn,
	"aaa": lys,
	"aag": lys,
	//
	"agu": ser,
	"agc": ser,
	"aga": arg,
	"agg": arg,
	//
	"guu": val,
	"guc": val,
	"gua": val,
	"gug": val,
	//
	"gcu": ala,
	"gcc": ala,
	"gca": ala,
	"gcg": ala,
	//
	"gau": asp,
	"gac": asp,
	"gaa": glu,
	"gag": glu,
	//
	"ggu": gly,
	"ggc": gly,
	"gga": gly,
	"ggg": gly,
}

type rna string

func (input *rna) decodeNextAminoAcid() (amino, error) {
	if len(*input) < 3 {
		return unknown, errors.New("one codon require 3 nucliotides")
	}

	codon := (*input)[:3]
	*input = (*input)[3:]

	aminoAcid, ok := CODON_MAP[codon]
	if !ok {
		return unknown, errors.New("invalid nucliotide sequence")
	}

	return aminoAcid, nil
}

func (input *rna) nonEmpty() bool {
	return len(*input) >= 3
}

type rnaIter struct {
	r     rna
	index int
}

func newRnaIter(r rna) *rnaIter {
	return &rnaIter{
		r:     r,
		index: 0,
	}
}

func (self *rnaIter) hasNext() bool {
	return self.index <= len(self.r)-3
}

func (self *rnaIter) next() (amino, error) {
	if !self.hasNext() {
		return unknown, errors.New("one codon require 3 nucliotides")
	}

	codon := self.r[self.index : self.index+3]
	self.index += 3

	aminoAcid, ok := CODON_MAP[codon]
	if !ok {
		return unknown, errors.New("invalid nucliotide sequence")
	}

	return aminoAcid, nil
}

func main() {
	rdr := bufio.NewReader(os.Stdin)
	line, err := rdr.ReadString('\n')
	if err != nil {
		fmt.Println(err)
		os.Exit(-1)
	}

	line = utils.TrimInput(line)

	var rnaLine rna = rna(line)

	peptide, err := processRna(rnaLine)
	if err != nil {
		return
	}

	fmt.Println(peptide)
}

func processRna(rnaLine rna) ([]amino, error) {
	peptide := make([]amino, 0, len(rnaLine)/3)

	ri := newRnaIter(rnaLine)
	for ri.hasNext() {
		aminoAcid, err := ri.next()
		if err != nil {
			return nil, err
		}
		peptide = append(peptide, aminoAcid)
	}
	return peptide, nil
}
