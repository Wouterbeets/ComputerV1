package main

import (
	"errors"
	"flag"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type poly struct {
	degree   int
	eqString string
	terms    []term
	operator []string
}

type term struct {
	multip int
	indet  string
	power  int
}

func (t *term) parseMul(mulStr string) error {
	mul, err := strconv.Atoi(mulStr)
	if err != nil {
		return err
	}
	t.multip = mul
	return nil
}

func (t *term) checkMulOp(opStr string) error {
	if opStr == "*" {
		return nil
	} else {
		return errors.New("operator before indeterminate must be \"*\"")
	}
}

func (t *term) parsePow(powStr string) error {
	indPow := strings.Split(powStr, "^")
	t.indet = indPow[0]
	p, err := strconv.Atoi(indPow[1])
	if err != nil {
		return err
	}
	t.power = p
	return nil
}

func main() {
	flag.Parse()
	if flag.NArg() != 1 {
		usage(nil)
	} else {
		eq := &poly{}
		eq.parseEq(flag.Arg(0))
		fmt.Printf("%+v\n", eq)
	}
}

func errorHandel(err error) {
	if err != nil {
		usage(err)
	}
}

func usage(err error) {
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("usage example: ./ComputerV1 \"2 * X^2 + 3 * X^1 * 14 * X^2\"")
	os.Exit(-1)
}

func (eq *poly) parseEq(str string) {
	termStrs := strings.Split(str, " ")
	fmt.Printf("%+v\n", str)
	for i := 0; i < len(termStrs); i++ {
		t := term{}
		err := t.parseMul(termStrs[i])
		errorHandel(err)
		err = t.checkMulOp(termStrs[i+1])
		errorHandel(err)
		err = t.parsePow(termStrs[i+2])
		errorHandel(err)
		if err != nil {
			panic(err)
		}
		i = i + 3
		if ((i+1)%4) == 0 && i != 0 && i < len(termStrs) {
			eq.operator = append(eq.operator, termStrs[i])
		}
		eq.terms = append(eq.terms, t)
	}
}
