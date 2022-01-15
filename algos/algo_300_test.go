package algos_test

import (
	"fmt"
	"testing"

	"github.com/keithyin/lc-go/algos"
)

func TestAlgo300(t *testing.T) {

	values := []int{0, 1, 0, 3, 2, 3}

	val := algos.LengthOfLIS(values)
	fmt.Println(val)
}
