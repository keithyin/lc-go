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

func TestInt(t *testing.T) {
	var a int32
	a = -0x80000000
	fmt.Println(a)
	a = 0x7fffffff
	fmt.Println(a)
}

func TestAlgo98(t *testing.T) {
	root := algos.BuildTree([]int{2, 1, 3})
	algos.IsValidBST(root)
}

func TestSortList(t *testing.T) {
	root := algos.BuildList([]int{4, 2, 1, 3})
	algos.SortList(root)
}

func TestAlgo96(t *testing.T) {
	res := algos.NumTrees(3)
	fmt.Println(res)
}

func TestAlgo84(t *testing.T) {
	res := algos.LargestRectangleArea([]int{1, 1})
	fmt.Println(res)
}

func TestAlgo76(t *testing.T) {
	res := algos.MinCoveredString("ADOBECODEBANC", "ABC")
	fmt.Println(res)
}

func TestCanPartition(t *testing.T) {
	nums := []int{1, 2, 5}
	fmt.Println(algos.CanPartition(nums))
}

func TestAlgo347(t *testing.T) {
	vals := []int{4, 1, -1, 2, -1, 2, 3}
	topK := 2
	fmt.Println(algos.TopKFrequent(vals, topK))
}

func TestAlgo394(t *testing.T) {
	s := "3[a]2[bc]"
	fmt.Printf("result: [%s]", algos.DecodeString(s))
}

func TestAlgo207(t *testing.T) {
	pre := [][]int{{1, 0}, {0, 1}}
	fmt.Println(algos.CanFinish(2, pre))
}

func TestAlgo139(t *testing.T) {

	fmt.Println(algos.WordBreak("leetcode", []string{"leet", "code"}))
}
