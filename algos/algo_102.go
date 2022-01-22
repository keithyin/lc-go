package algos

import "container/list"

func levelOrder(root *TreeNode) [][]int {
	queue1 := list.New()
	queue2 := list.New()
	results := make([][]int, 0)
	queue1Tracer := make([]int, 0)
	queue2Tracer := make([]int, 0)

	if root == nil {
		return results
	}

	queue1.PushBack(root)
	for queue1.Len() > 0 || queue2.Len() > 0 {
		for queue1.Len() > 0 {
			node := queue1.Front()
			nodeVal := node.Value.(*TreeNode)
			queue1Tracer = append(queue1Tracer, nodeVal.Val)
			if nodeVal.Left != nil {
				queue2.PushBack(nodeVal.Left)
			}
			if nodeVal.Right != nil {
				queue2.PushBack(nodeVal.Right)
			}
			queue1.Remove(node)
		}

		if len(queue1Tracer) > 0 {
			results = append(results, queue1Tracer)
			queue1Tracer = make([]int, 0)
		}

		for queue2.Len() > 0 {
			node := queue2.Front()
			nodeVal := node.Value.(*TreeNode)
			queue2Tracer = append(queue2Tracer, nodeVal.Val)
			if nodeVal.Left != nil {
				queue1.PushBack(nodeVal.Left)
			}
			if nodeVal.Right != nil {
				queue1.PushBack(nodeVal.Right)
			}
			queue2.Remove(node)
		}
		if len(queue2Tracer) > 0 {
			results = append(results, queue2Tracer)
			queue2Tracer = make([]int, 0)
		}
	}

	return results

}
