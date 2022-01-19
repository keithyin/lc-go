package algos_test

import (
	"fmt"
	"testing"

	"github.com/keithyin/lc-go/algos"
)

func TestBuildTree(t *testing.T) {

}

func TestAlgo448(t *testing.T) {
	nums := []int{4, 3, 2, 7, 8, 2, 3, 1}
	algos.FindDisappearedNumbers(nums)
}

func TestFindKthMax(t *testing.T) {
	nums := []int{3, 2, 1, 5, 6, 4}
	res := algos.FindKthLargest(nums, 2)
	fmt.Println(res)
}
