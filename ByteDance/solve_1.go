package main
// https://www.nowcoder.com/practice/c215ba61c8b1443b996351df929dc4d4
import (
	"strconv"
	"fmt"
)

func main() {
	a := solve("1+2")
	fmt.Println(a)
}

func solve( s string ) int {
	numStack, opStack := []int{}, []byte{}

	cal := func() {
		numSize, opSize := len(numStack), len(opStack)
		a, b := numStack[numSize-2], numStack[numSize-1]
		op := opStack[opSize-1]
		numStack = numStack[:numSize-2]
		opStack = opStack[:opSize-1]
		var value int
		if op == '+' {
			value = a+b
		} else if op == '-' {
			value = a-b
		} else if op == '*' {
			value = a*b
		} else {
			value = a/b
		}
		numStack = append(numStack, value)
	}

	cmp := func(op1, op2 byte) bool {
		if op1 == '(' {
			return false
		}
		if (op1 == '+' || op1 == '-') && (op2 == '*' || op2 == '/') {
			return false
		}
		return true
	}

	s += ")"
	opStack = append(opStack, '(')
	nextIsOp := false
	for i := 0; i < len(s); i++ {
		if s[i] == '(' {
			opStack = append(opStack, s[i])
		} else if s[i] == ')' {
			for opStack[len(opStack)-1] != '(' {
				cal()
			}
			opStack = opStack[:len(opStack)-1]
		} else if nextIsOp {
			for cmp(opStack[len(opStack)-1], s[i]) {
				cal()
			}
			opStack = append(opStack, s[i])
			nextIsOp = false
		} else {
			j := i
			if s[i] == '+' || s[i] == '-' {
				i++
			}
			for ;s[i] >= '0' && s[i] <= '9';i++{
			}
			value, _ := strconv.Atoi(s[j:i])
			numStack = append(numStack, value)
			i--
			nextIsOp = true
		}
	}
	return numStack[0]
}
