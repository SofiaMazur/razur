package lab2

import (
	"fmt"
	"testing"
	. "gopkg.in/check.v1"
)

func TestImplementation (t *testing.T) { TestingT(t) }

func (s *TestSuite) TestPostfixToInfix (c *C) {
	examples := map[string]string {
		"345 34.902 + 23 /": "(345 + 34.902) / 23",
		"300000004 45 +": "300000004 + 45",
		"26 89 13 +": "too many operands",
		"23,89 22 ^ 87 +": "23,89 ^ 22 + 87",
		"I'm an error: la-la-la": "invalid input expression",
		"9.007 765.9999994 + 56 ^ /": "too many operators",
		"23 15 98.765 21 23 / 90 - ^ 64445 8 - 90 + * * -": "23 - 15 * 98.765 ^ (21 / 23 - 90) * (64445 - 8 + 90)",
		"": "invalid input expression",
		"90 3 6 4 3 0.44 7 - / ^ - * +": "90 + 3 * (6 - 4 ^ (3 / (0.44 - 7)))",
	}

	for postfix, expected := range examples {
		res, err := PostfixToInfix(postfix)
		if err != nil {
			c.Assert(err, ErrorMatches, expected)
		} else {
			c.Assert(res, Equals, expected)
		}
	}
}

func ExamplePostfixToInfix() {
	res, err := PostfixToInfix("34 90.76 + 4 ^")
	if err != nil {
		panic(err)
	} else {
		fmt.Println(res)
	}

	// Output:
	// (34 + 90.76) ^ 4
}

