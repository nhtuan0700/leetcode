# https://leetcode.com/problems/lru-cache/

# TC of problems: O(capacity), SC: O(capacity)
# get: O(1)
# put: O(1)
class Node:
    def __init__(self, key = '', value = None):
        self.next = None
        self.value = value
        self.key = key
        self.prev = None

class DoublyList:
    def __init__(self):
        head = Node()
        tail = Node()
        head.next = tail
        tail.prev = head
        self.head = head
        self.tail = tail

    # add into the head
    def insert(self, newNode):
        newNode.prev = self.head
        newNode.next = self.head.next
        newNode.next.prev = newNode
        newNode.prev.next = newNode
    def delete(self, node):
        node.next.prev = node.prev
        node.prev.next = node.next

    # remove the node of the tail
    def pop(self):
        node = self.tail.prev
        self.delete(self.tail.prev)
        return node


class LRUCache:
    def __init__(self, capacity: int):
        # key: string, value: Node
        self.dictStore = {}
        self.capacity = capacity
        self.list = DoublyList()        

    # O(1)
    def get(self, key: int) -> int:
        if key in self.dictStore:
            node = self.dictStore[key]
            self.list.delete(node)
            self.list.insert(node)
            return node.value
        return -1
        

    # O(1)
    def put(self, key: int, value: int) -> None:
        if key in self.dictStore:
            node = self.dictStore[key]
            node.value = value
            self.list.delete(node)
        else:
            node = Node(key, value)
            self.dictStore[key] = node
        self.list.insert(node)
        if len(self.dictStore) > self.capacity:
            node = self.list.pop()
            self.dictStore.pop(node.key)

# Your LRUCache object will be instantiated and called as such:
# obj = LRUCache(capacity)
# param_1 = obj.get(key)
# obj.put(key,value)
