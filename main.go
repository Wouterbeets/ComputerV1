package main

import (
	"flag"
	"fmt"
	"github.com/Wouterbeets/poly"
	"os"
)

func main() {
	flag.Parse()
	if flag.NArg() != 1 {
		usage(nil)
	} else {
		eq := &poly.Poly{}
		err := eq.ParseEq(flag.Arg(0))
		errorHandel(err)
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
