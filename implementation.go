package lab2

import (
	"regexp"
	"strings"
	"fmt"
)

// Main export function PostfixToInfix takes as an argument
// polish postfix notation string, return result (if input
// argument is correct, default empty string) and particular
// error if computing was failed (defalut nil)

// struct for postfix input checking
type validator struct {
	ValidOperatorExp string				// operators regexp string
	ValidOperandExp string				// operands regexp string
}

// checks input string
func (v *validator) VaildInput(input string) bool {
	validator := fmt.Sprintf(`^((%s|%s)\s){2,}(%s\s){0,}%s$`, v.ValidOperandExp, v.ValidOperatorExp, v.ValidOperatorExp, v.ValidOperatorExp)
	isValid, _ := regexp.MatchString(validator, input)
	return isValid
}

// checks slipped prefix string (array of operand and operators)
func (v *validator) CheckArgsAmount(args []string) error {
	operators, operands := 0, 0
	macthOperator := fmt.Sprintf(`^%s$`, v.ValidOperatorExp)
	macthOperand := fmt.Sprintf(`^%s$`, v.ValidOperandExp)
	for _, arg := range args {
		if isOperator, _ := regexp.MatchString(macthOperator, arg); isOperator {
			operators++
		} else if isOperand, _ := regexp.MatchString(macthOperand, arg); isOperand {
			operands++
		}
	}
	switch {
	case operators + operands != len(args):
		return fmt.Errorf("invalid expression argument(s)")
	case operands > operators + 1:
		return fmt.Errorf("too many operands")
	case operators > operands - 1:
		return fmt.Errorf("too many operators")
	default:
		return nil
	}
}

// checks if given value has any operator (e.g. "a + b...")
func (v *validator) IncludesOperator(str string) bool {
	includes, _ := regexp.MatchString(v.ValidOperatorExp, str)
	return includes
}

// checks if given value is any operator
func (v *validator) IsOperator(str string) bool {
	macthOperator := fmt.Sprintf(`^%s$`, v.ValidOperatorExp)
	includes, _ := regexp.MatchString(macthOperator, str)
	return includes
}

// PostfixToInfix function for converting polish notation
func PostfixToInfix(postfixStr string) (infixStr string, err error) {
	v := validator{ValidOperatorExp: `[-\+\*\^\/]`, ValidOperandExp: `(\d+|(\d+[,\.]\d+))`}
	if !v.VaildInput(postfixStr) {
		err = fmt.Errorf("invalid input expression")
		return
	}
	var operatorsStack []string
	var infixHeap []string
	operatorsPriorities := map[string]uint8 {
		"+": 1,
		"-": 1,
		"*": 2,
		"/": 2,
		"^": 3,	
	}
	postfixArgs := strings.Split(postfixStr, " ")
	if agrsErr := v.CheckArgsAmount(postfixArgs); agrsErr != nil {
		err = agrsErr
		return
	}
	// calcucaling cycle if given string is valid
	for _, arg := range postfixArgs {
		if !v.IsOperator(arg) {
			infixHeap = append(infixHeap, arg)
			continue
		}
		operator := arg
		operatorsStack = append(operatorsStack, operator)
		slicedEnd := len(infixHeap) - 2
		sliced := infixHeap[slicedEnd:]
		infixHeap = infixHeap[:slicedEnd]
		operand1, operand2 := sliced[0], sliced[1]
		// expression has more than one calculating operation
		if len(operatorsStack) > 1 {
			prevOperatorIndex := len(operatorsStack) - 2
			prevOperator := operatorsStack[prevOperatorIndex]
			// brackets wrapping
			isPowerOperators := operatorsPriorities[operator] == 3 && operatorsPriorities[prevOperator] == 3
			higherPriority := operatorsPriorities[operator] > operatorsPriorities[prevOperator]
			if higherPriority || isPowerOperators {
				if v.IncludesOperator(operand2) {
					operand2 = "(" + operand2 + ")"
				} else {
					operand1 = "(" + operand1 + ")"
				} 
			}
		}
		// appending new operand (joined two previous operands)
		operand := fmt.Sprintf("%s %s %s", operand1, operator, operand2)
		infixHeap = append(infixHeap, operand)
	}
	// the firs and only element in stack is the final result
	infixStr = infixHeap[0]
	return
}
