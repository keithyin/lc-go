package algos

type GraphNode struct {
	value  int
	childs []*GraphNode
}

func NewGraphNode(val int) *GraphNode {
	return &GraphNode{value: val, childs: make([]*GraphNode, 0)}
}

func (this *GraphNode) addChild(child *GraphNode) {
	this.childs = append(this.childs, child)
}

func CanFinish(numCourses int, prerequisites [][]int) bool {
	return canFinish(numCourses, prerequisites)
}

func canFinish(numCourses int, prerequisites [][]int) bool {
	allNodes := make(map[int]*GraphNode)
	for _, item := range prerequisites {
		curCourse := item[0]
		dependenceCourse := item[1]
		var curCourseNode *GraphNode
		if _, ok := allNodes[curCourse]; ok {
			curCourseNode = allNodes[curCourse]
		} else {
			curCourseNode = NewGraphNode(curCourse)
			allNodes[curCourse] = curCourseNode
		}

		var dependenceCourseNode *GraphNode
		if _, ok := allNodes[dependenceCourse]; ok {
			dependenceCourseNode = allNodes[dependenceCourse]
		} else {
			dependenceCourseNode = NewGraphNode(dependenceCourse)
			allNodes[dependenceCourse] = dependenceCourseNode
		}
		dependenceCourseNode.addChild(curCourseNode)
	}
	cache := make(map[*GraphNode]bool)
	for _, node := range allNodes {
		visitMap := make(map[*GraphNode]bool)
		if !isDag(node, visitMap, cache, node, true) {
			return false
		}
	}
	return true

}

func isDag(root *GraphNode, visitMap map[*GraphNode]bool, cache map[*GraphNode]bool, anchor *GraphNode, first bool) bool {
	if !first && root == anchor {
		return false
	}
	if root == nil {
		return true
	}

	if _, ok := visitMap[root]; ok {
		return true
	}

	visitMap[root] = true

	for _, node := range root.childs {
		if !isDag(node, visitMap, cache, anchor, false) {
			return false
		}

	}
	return true
}
