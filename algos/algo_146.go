package algos

type TaskNode struct {
	key  int
	val  int
	next *TaskNode
	pre  *TaskNode
}

type LRUCache struct {
	capacity int
	curLen   int
	Begin    *TaskNode
	End      *TaskNode

	Key2TaskNode map[int]*TaskNode
}

func LRUCacheConstructor(capacity int) LRUCache {
	cache := LRUCache{
		capacity:     capacity,
		curLen:       0,
		Begin:        new(TaskNode),
		End:          new(TaskNode),
		Key2TaskNode: make(map[int]*TaskNode),
	}

	cache.Begin.next = cache.End
	cache.End.pre = cache.Begin

	return cache
}

func (this *LRUCache) Get(key int) int {

	if val, ok := this.Key2TaskNode[key]; ok {
		firstNode := this.Begin.next
		if firstNode != val {
			val.pre.next = val.next
			val.next.pre = val.pre

			val.next = this.Begin.next
			val.next.pre = val
			val.pre = this.Begin
			this.Begin.next = val

		}
		return val.val
	}
	return -1

}

func (this *LRUCache) Put(key int, value int) {

	if val, ok := this.Key2TaskNode[key]; ok {
		val.val = value
		if val != this.Begin.next {
			val.pre.next = val.next
			val.next.pre = val.pre

			val.next = this.Begin.next
			val.next.pre = val
			val.pre = this.Begin
			this.Begin.next = val

		}
		return

	}

	newTask := new(TaskNode)
	newTask.key = key
	newTask.val = value

	if this.curLen == this.capacity {
		key2del := this.End.pre.key
		prePre := this.End.pre.pre
		prePre.next = this.End
		this.End.pre = prePre
		delete(this.Key2TaskNode, key2del)
	}

	newTask.next = this.Begin.next
	newTask.pre = this.Begin

	this.Begin.next.pre = newTask
	this.Begin.next = newTask
	this.Key2TaskNode[key] = newTask
	this.curLen = len(this.Key2TaskNode)

}
