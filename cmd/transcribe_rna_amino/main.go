package main

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"strings"
)

func prepareString(line string) string {
	line = strings.TrimSuffix(line, "\n\r\t ")
	line = strings.TrimPrefix(line, "\t ")
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
		panic("got unknown amino acid" + fmt.Sprint(a))
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

	c, ok := CODON_MAP[codon]
	if !ok {
		return unknown, errors.New("invalid nucliotide sequence")
	}

	*input = (*input)[3:]

	return c, nil
}

func (input *rna) nonEmpty() bool {
	return len(*input) >= 3
}

func main() {
	rdr := bufio.NewReader(os.Stdin)
	line, err := rdr.ReadString('\n')
	if err != nil {
		fmt.Println(err)
		os.Exit(-1)
	}

	var rnaLine rna = rna(line)

	for rnaLine.nonEmpty() {
		v, err := rnaLine.decodeNextAminoAcid()
		if err != nil {
			fmt.Println(err)
			return
		}
		fmt.Println(v)
	}
}
