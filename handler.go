package lab2

import (
	"strings"
	"io"
	"fmt"
	"io/ioutil"
)

// ComputeHandler prefix to infix interpreter
type ComputeHandler struct {
	Input io.Reader
	Output io.Writer
}

// Compute prefix to infix wrapper
func (ch *ComputeHandler) Compute() error {
	if ch.Input == nil {
		return fmt.Errorf("input is not specified");
	}
	if ch.Output == nil {
		return fmt.Errorf("output is not specified");
	}
	buf, readErr := ioutil.ReadAll(ch.Input)
	if readErr != nil {
		return readErr
	}
	strInput := strings.Trim(string(buf), "\n")
	computed, computeErr := PostfixToInfix(strInput)
	if computeErr != nil {
		return computeErr
	}
	res := []byte(computed + "\n")
	_, writeErr := ch.Output.Write(res)
	if writeErr != nil {
		return writeErr
	}
	return nil
}
