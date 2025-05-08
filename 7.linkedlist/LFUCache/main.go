package main


type Node struct {
	Val  int
	Key  int
	Next *Node
	Prev *Node
	List *List
}

type List struct {
	Root  *Node
	Next  *List
	Prev  *List
	Count int
}

func NewList(count int) *List {
	root := &Node{}
	root.Next = root
	root.Prev = root
	return &List{
		Root:  root,
		Count: count,
	}
}

type LFUCache struct {
	LFU      *List
	Capacity int
	Cache    map[int]*Node
}

func Constructor(capacity int) LFUCache {
	rootList := NewList(0)
	rootList.Next = rootList
	rootList.Prev = rootList

	return LFUCache{
		LFU:      rootList,
		Capacity: capacity,
		Cache:    make(map[int]*Node),
	}
}

func (this *LFUCache) Get(key int) int {
	node, ok := this.Cache[key]
	if !ok {
		return -1
	}

	// remove node
	node.Next.Prev = node.Prev
	node.Prev.Next = node.Next

	currentList := node.List
	count := currentList.Count + 1
	// if currentList has one element then remove currentList
	if currentList.Root.Next == currentList.Root {
		currentList.Next.Prev = currentList.Prev
		currentList.Prev.Next = currentList.Next

		count = count
		currentList = currentList.Prev
	}

	if count == currentList.Next.Count {
		// root of next list
		nextRoot := currentList.Next.Root
		// move node into tail of next list
		node.Next = nextRoot
		node.Prev = nextRoot.Prev
		nextRoot.Prev = node
		node.Prev.Next = node

		node.List = currentList.Next
	} else {
		newList := NewList(count)

		// append newList after currentList
		newList.Next = currentList.Next
		newList.Prev = currentList
		currentList.Next = newList
		newList.Next.Prev = newList

		// append node into newList
		node.Prev = newList.Root.Prev
		node.Next = newList.Root
		newList.Root.Prev = node
		node.Prev.Next = node
		node.List = newList
	}

	return node.Val
}

func (this *LFUCache) Put(key int, value int) {
	node, ok := this.Cache[key]
	if !ok {
		headList := this.LFU.Next
		if len(this.Cache) >= this.Capacity {
			// remove head node
			headNode := headList.Root.Next
			headNode.Prev.Next = headNode.Next
			headNode.Next.Prev = headNode.Prev

			if headList.Root.Next == headList.Root {
				headList.Prev.Next = headList.Next
				headList.Next.Prev = headList.Prev
			}
			delete(this.Cache, headNode.Key)
		}

		newNode := &Node{Key: key, Val: value}
		// if not exist headList or headList has count > 1
		if headList.Root.Next == headList.Root || headList.Count > 1 {
			newList := NewList(1)
			// insert new list after head list
			rootList := this.LFU
			newList.Next = rootList.Next
			newList.Prev = rootList
			rootList.Next = newList
			newList.Next.Prev = newList

			// apend newNode into list
			newNode.Next = newList.Root
			newNode.Prev = newList.Root.Prev
			newNode.Next.Prev = newNode
			newNode.Prev.Next = newNode

			newNode.List = newList
		} else {
			newNode.Next = headList.Root
			newNode.Prev = headList.Root.Prev
			newNode.Next.Prev = newNode
			newNode.Prev.Next = newNode
			newNode.List = headList
		}

		this.Cache[key] = newNode
	} else {
		currentList := node.List
		count := currentList.Count + 1
		node.Next.Prev = node.Prev
		node.Prev.Next = node.Next
		// remove currentList if having one node
		if currentList.Root.Next == currentList.Root {
			currentList.Next.Prev = currentList.Prev
			currentList.Prev.Next = currentList.Next

			count = count
			currentList = currentList.Prev
		}
		if count == currentList.Next.Count {
			// root of next list
			nextRoot := currentList.Next.Root
			// move node into tail of next list
			node.Next = nextRoot
			node.Prev = nextRoot.Prev
			nextRoot.Prev = node
			node.Prev.Next = node

			node.List = currentList.Next
		} else {
			newList := NewList(count)

			// append newList after currentList
			newList.Next = currentList.Next
			newList.Prev = currentList
			currentList.Next = newList
			newList.Next.Prev = newList

			// append node into newList
			node.Prev = newList.Root.Prev
			node.Next = newList.Root
			newList.Root.Prev = node
			node.Prev.Next = node
			node.List = newList
		}
		node.Val = value
	}
}

/*
Node {
	Val int
	Key int
	List *List
	Prev *Node
	Next *Node
}

List {
	Root *Node
	Count int
	Next *List
	Prev *List
}

RootList(cnt=0) <-> List(cnt=1) <-> List(cnt=...)
===

LFU:
- List: value, store the usage order with count -> recently used
	- head: least recently used
	- tail: most recently used
- Cache: map[key] -> node of DLL
- Capacity
- RootList List

get(int key):
- if key not in cache:
    - return -1
- else:
    - if CurrentList.Next != nil && CurrentList.Count + 1 == CurrentList.Next.Count:
				- Append node into NextList
    - else:
				- If CurrentList has one element need to remove CurrentList and assign CurrentList with Previous
        - NewList with count++
        - Append node into NewList
				- NewList.Prev = CurrentList
				- NewList.Next = CurrentList.Next
        - CurrentList.Next = NewList
        - NewList.Next.Prev = NewList
		- return node's value

put(int key, int value):
- if key not in cache:
    - if current capacity == capcity:
        - remove the head node of HeadList
		- if HeadList count > 1:
        - NewList with count=1
        - Append node into NewList
				- NewList.Prev = HeadList
				- NewList.Next = HeadList.Next
        - HeadList.Next = NewList
				- NewList.Next.Prev = NewList
    - newNode: push value into back of list
    - cached[key] = newNode
- else:
    - if CurrentList.Next != nil && CurrentList.Count + 1 == CurrentList.Next.Count:
				- Append node into NextList
    - else:
				- If CurrentList has one element need to remove CurrentList and assign CurrentList with Previous
        - NewList with count++
        - Append node into NewList
				- NewList.Prev = CurrentList
				- NewList.Next = CurrentList.Next
        - CurrentList.Next = NewList
        - NewList.Next.Prev = NewList

[[2],[1,1],[2,2],[1],[3,3],[2],[3],[4,4],[1],[3],[4]]

capacity[2], LFU <->
- put [1,1] LFU <-> [root<->1](1)
- put [2,2] LFU <-> [root<->1<->2](1)
- get [1] LFU <-> [root<->2](1) <-> [root<->1](2)
- put [3,3] LFU <-> [root<->3](1) <-> [root<->1](2)
- get [2] LFU <-> [root<->3](1) <-> [root<->1](2) => -1
- get [3] LFU <-> [root<->1<->3](2)
- put [4] LFU <-> [root<->4](1) <-> [root<->3](2)
- get [1] LFU <-> [root<->4](1) <-> [root<->3](2) => -1
- get [3] LFU <-> [root<->4](1) <-> [root<->3](3)
- get [4] LFU <-> [root<->4](2) <-> [root<->3](3)

[[2],[2,1],[2,2],[2],[1,1],[4,1],[2]]
capacity[2], LFU ->
- put [2,1] LFU <-> [root<->2](1)
- put [2,2] LFU <-> [root<->2](2)
- get [2] LFU <-> [root<->2](3) => 2
- put [1,1] LFU <-> [root<->1](1) <-> [root<->2](3)
- put [4,1] LFU <-> [root<->4](1) <-> [root<->2](3)
- get [2]  <-> [root<->4](1) <-> [root<->2](4) => 2


*/

/**

Your LFUCache object will be instantiated and called as such:
obj := Constructor(capacity);
param_1 := obj.Get(key);
obj.Put(key,value); */
