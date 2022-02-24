package algos

import (
	"container/list"
	"fmt"
	"strconv"
	"strings"
)

type Codec struct {
}

func TreeSerilizationConstructor() Codec {
	return Codec{}
}

// Serializes a tree to a single string.
func (this *Codec) serialize(root *TreeNode) string {
	result := ""

	myQueue := list.New()
	if root == nil {
		result += "null"
		return result
	}
	result = fmt.Sprintf("%s,%d", result, root.Val)
	myQueue.PushBack(root)
	for myQueue.Len() > 0 {
		node := myQueue.Front()
		myQueue.Remove(node)
		nodeVal := node.Value.(*TreeNode)
		if nodeVal.Left == nil {
			result = fmt.Sprintf("%s,null", result)
		} else {
			result = fmt.Sprintf("%s,%d", result, nodeVal.Left.Val)
			myQueue.PushBack(nodeVal.Left)
		}

		if nodeVal.Right == nil {
			result = fmt.Sprintf("%s,null", result)
		} else {
			result = fmt.Sprintf("%s,%d", result, nodeVal.Right.Val)
			myQueue.PushBack(nodeVal.Right)
		}
	}
	return result

}

// Deserializes your encoded data to tree.
func (this *Codec) deserialize(data string) *TreeNode {
	if data == "null" {
		return nil
	}
	idx := 0
	items := strings.Split(data, ",")
	root := new(TreeNode)
	tmp, _ := strconv.ParseInt(items[idx], 10, 32)
	root.Val = int(tmp)
	myQueue := list.New()
	myQueue.PushBack(root)
	for idx < len(items) && myQueue.Len() > 0 {
		curNode := myQueue.Front()
		myQueue.Remove(curNode)
		curNodeVal := curNode.Value.(*TreeNode)
		if (idx + 1) >= len(items) {
			break
		}
		if items[idx+1] != "null" {
			left := new(TreeNode)
			tmp, _ = strconv.ParseInt(items[idx+1], 10, 32)
			left.Val = int(tmp)
			curNodeVal.Left = left
			myQueue.PushBack(left)
		}

		if (idx + 2) >= len(items) {
			break
		}
		if items[idx+2] != "null" {
			right := new(TreeNode)
			tmp, _ = strconv.ParseInt(items[idx+2], 10, 32)
			right.Val = int(tmp)
			curNodeVal.Right = right
			myQueue.PushBack(right)
		}
		idx += 2
	}
	return root
}
