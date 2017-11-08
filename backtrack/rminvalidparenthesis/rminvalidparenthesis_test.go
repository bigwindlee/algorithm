package rminvalid

import (
	"fmt"
	"testing"
)

func TestRemoveInvalidParenthesis(t *testing.T) {
	fmt.Printf("%v\n", removeInvalidParenthesis("()())()"))
	fmt.Printf("%v\n", removeInvalidParenthesis("(v)())()"))
}
