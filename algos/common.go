package algos

import (
	"container/list"
	"math/rand"
)

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

const NIL_VALUE int = 999999

func BuildTree(values []int) *TreeNode {
	return BuildTreeCore(values, 0)
}

func BuildTreeCore(values []int, curPos int) *TreeNode {
	var curNode *TreeNode
	if curPos >= len(values) {
		return nil
	}
	if values[curPos] != NIL_VALUE {
		curNode = new(TreeNode)
		curNode.Val = values[curPos]
		curNode.Left = BuildTreeCore(values, 2*curPos+1)
		curNode.Right = BuildTreeCore(values, 2*curPos+2)
	}
	return curNode
}

func (obj *TreeNode) levelTraverse() []int {
	iteratedValues := make([]int, 0)
	mylist := list.New()
	mylist.PushBack(obj)
	for mylist.Len() > 0 {
		front := mylist.Front()
		iteratedValues = append(iteratedValues, front.Value.(*TreeNode).Val)
		if front.Value.(*TreeNode).Left != nil {
			mylist.PushBack(front.Value.(*TreeNode).Left)
		}
		if front.Value.(*TreeNode).Right != nil {
			mylist.PushBack(front.Value.(*TreeNode).Right)
		}
		mylist.Remove(front)
	}
	return iteratedValues
}

type ListNode struct {
	Val  int
	Next *ListNode
}

func BuildList(nums []int) *ListNode {
	root := new(ListNode)
	cursor := root
	for i := 0; i < len(nums); i++ {
		current := new(ListNode)
		current.Val = nums[i]
		cursor.Next = current
		cursor = cursor.Next
	}
	return root.Next
}

func SliceMinInt(values []int) int {
	minVal := values[0]
	for i := 1; i < len(values); i++ {
		if values[i] < minVal {
			minVal = values[i]
		}
	}
	return minVal
}

func SliceMaxInt(values []int) int {
	maxVal := values[0]
	for i := 1; i < len(values); i++ {
		if maxVal < values[i] {
			maxVal = values[i]
		}
	}
	return maxVal
}

func Shuffle(values []int) {
	rand.Seed(2022)
	for i := 1; i < len(values); i++ {
		randomIdx := rand.Intn(i + 1)
		tmp := values[i]
		values[i] = values[randomIdx]
		values[randomIdx] = tmp
	}
}
