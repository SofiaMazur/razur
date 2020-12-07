package main

import (
	"flag"
	"fmt"
	"io"
	"os"
	"strings"
	
	lab2 "github.com/SofiaMazur/razur_lab_2"
)

func getFlagsValues() (inputExpression, inputFilename, outputFilename *string) {
	defer flag.Parse()
	// -e key used to enter expression through the command line
	inputExpression = flag.String("e", "", "Expression to compute")
	// -f key used to enter expression through the file
	// -e and -f can not be combined!
	inputFilename = flag.String("f", "", "input file")
	// -o key used to output the result to the file mentioned
	outputFilename = flag.String("o", "", "output file")
	return
}

func main() {
	var inputExpression, inputFilename, outputFilename *string = getFlagsValues()
	// If both filename and expression are full-filled
	if *inputFilename != "" && *inputExpression != "" {
		err := fmt.Errorf("only one source of expression needed")
		panic(err)
	}
	var input io.Reader
	var output io.Writer
	if *inputExpression != "" {
		input = strings.NewReader(*inputExpression)
	} else if *inputFilename != "" {
		input, _ = os.Open(*inputFilename)
	}
	if *outputFilename != "" {
		output, _ = os.Create(*outputFilename)
	} else {
		output = os.Stdout
	}
	handler := lab2.ComputeHandler{Input: input, Output: output}
	err := handler.Compute()
	if err != nil {
		panic(err)
	}
}
