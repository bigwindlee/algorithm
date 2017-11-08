package permutation

import (
	"fmt"
	"testing"
)

func TestPermutation(t *testing.T) {
	res := permute([]int{1, 2, 3, 4})
	fmt.Printf("Total: %d\n", len(res))
	for i := range res {
		fmt.Printf("%v\n", res[i])
	}
}
