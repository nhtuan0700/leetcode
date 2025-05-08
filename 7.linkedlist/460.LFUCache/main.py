class Node:
    def __init__(self, key=None, value=None):
        self.key = key
        self.value = value
        self.next = None
        self.prev = None
        self.parent = None

class DLL:
    def __init__(self, cnt = 0, head = None):
        if not head:
            head = Node()
        head.next = head
        head.prev = head
        self.prev = None
        self.next = None
        self.head = head
        self.cnt = cnt

    def isEmpty(self):
        return self.head == self.head.prev

    def insert(self, node):
        node.next = self.head.next
        node.prev = self.head
        node.next.prev = node
        node.prev.next = node
        node.parent = self
        
    def append(self, node):
        node.prev = self.head.prev
        node.next = self.head
        node.next.prev = node
        node.prev.next = node
        node.parent = self

    def insertBetween(self, prevNode, nextNode):
        self.prev = prevNode
        self.next = nextNode
        self.prev.next = self
        self.next.prev = self
    
    def remove(self, node):
        node.prev.next = node.next
        node.next.prev = node.prev

    def pop(self):
        node = self.head.prev
        self.remove(node)
        return node



class LFUCache:
    def __init__(self, capacity: int):
        self.capacity = capacity
        self.dictStore = {}
        self.list = DLL(head=DLL())

    def get(self, key: int) -> int:
        if key not in self.dictStore:
            return -1
        node = self.dictStore[key]
        self.moveNode(node)

        return node.value

    def put(self, key: int, value: int) -> None:
        if key in self.dictStore:
            node = self.dictStore[key]
            node.value = value
            self.moveNode(node)

        else:
            self.evict()
            node = Node(key, value)
            # insert node into the dictStore
            self.dictStore[key] = node
            lastParent = self.list.head.prev
            # LFU is empty or last parent's count > 1
            if self.list.isEmpty() or lastParent.cnt > 1:
                # new parent
                newParent = DLL(cnt=1)
                newParent.insert(node)
                self.list.append(newParent)
            else:
                lastParent.insert(node)

    def evict(self):
        if len(self.dictStore) == self.capacity:
            lastParent = self.list.head.prev
            delNode = lastParent.pop()
            if lastParent.isEmpty():
                self.list.remove(lastParent)
            self.dictStore.pop(delNode.key)

    def moveNode(self, node):
        curParent = node.parent
        newCount = curParent.cnt + 1
        
        # remove the node from current parent
        curParent.remove(node)
        if curParent.isEmpty():
            nextParent = curParent.next
            self.list.remove(curParent)
            curParent = nextParent

        prevParent = curParent.prev
        # if found the target parent having count = newCount
        if prevParent.cnt == newCount:
            # insert node into the head
            prevParent.insert(node)
        else:
            # create a new parent
            # insert the new node into it
            # append the new parent to LFU's list
            newParent = DLL(cnt=newCount)
            newParent.insert(node)
            newParent.insertBetween(prevParent, curParent)

# LFU:
# - capacity
# - dictStore
# - head <-> node <-> node <-> ...
#   - node (inner DLL): 
#     - head <-> node <-> node <-> ... <-> tail
#     - cnt
# * remove tail first, freq descreased from head to tail
# * we will assign tail as head.prev
# 
# get(key):
# - if key in dictStore
#   - node = dictStore[key]
#   - parent = node.parent.count, newCount = parent.count + 1
#.  - if parent is empty, then remove the current parent, reasign the current parent by the next one
#.  - move the node into the parent.prev if parent.prev.cnt == count
#     - true: insert the node into the prev parent
#.    - false: create newParent with cnt = newCount and insert the node into it, move newParent to between prevParent and curParent
#   - return node.value
# - otherwise return -1
#
# put(key, value):
# - if key in dictStore
#.  - node = dictStore[key]
#   - parent = node.parent.count, newCount = parent.count + 1
#.  - if parent is empty, then remove the current parent, reasign the current parent by the next one
#.  - move the node into the parent.prev if parent.prev.cnt == count
#     - true: insert the node into the prev parent
#.    - false: create newParent with cnt = newCount and insert the node into it, move newParent to between prevParent and curParent
# - otherwise:
#   - find the parent with count = 1 by checking tail.prev.value == 1
#     - true: insert node into the head of the parent
#.    - false: create new parent having the new node just created and count = 1, append new parent into LFU's list
#  - if over capacity:
#    - targetParent = LFU.tail.prev, remove the last node of the targetParent
#    - if targetParent is empty, then remove it
