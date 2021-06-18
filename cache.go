package main

import "fmt"

type Node struct {
    key, value int
    prev, next *Node
}

type LRUCache struct {
    size       int
    capacity   int
    cache      map[int]*Node
    head, tail *Node
}

func initialize (key, value int) *Node {
    return &Node{
        key:   key,
        value: value,
    }
}

func Constructor(capacity int) LRUCache {
    l := LRUCache{
        capacity: capacity,
        cache:    map[int]*Node{},
        head:     initialize(0, 0),
        tail:     initialize(0, 0),
    }
    l.head.next = l.tail
    l.tail.prev = l.head
    return l
}

func (R *LRUCache) Get(key int) int {
    if node, ok := R.cache[key]; ok {
        R.moveToHead(node)
        return node.value
    }
    return -1
}

func (R *LRUCache) Put(key int, value int) {
    if node, ok := R.cache[key]; ok {
        node.value = value
        R.moveToHead(node)
    } else {
        node := initialize(key, value)
        R.cache[key] = node
        R.addToHead(node)
        R.size++
        if R.size > R.capacity {
            removed := R.removeTail()
            delete(R.cache, removed.key)
            R.size--
        }
    }
}

func (R *LRUCache) moveToHead(node *Node) {
    R.removeNode(node)
    R.addToHead(node)
}

func (R *LRUCache) removeNode(node *Node) {
    node.prev.next = node.next
    node.next.prev = node.prev
}

func (R *LRUCache) addToHead(node *Node) {
    node.prev = R.head
    node.next = R.head.next
    R.head.next.prev = node
    R.head.next = node
}

func (R *LRUCache) removeTail() *Node {
    node := R.tail.prev
    R.removeNode(node)
    return node
}

func main(){
   //var c int
   var k int
   var v int
   var ke int
   //fmt.Println("Enter Capacity :")
   //fmt.Scanf("%d",&c)
   f1 := Constructor(2);
   var i int
   L1: fmt.Println("ENTER 1 for PUT \n2 for GET ")
   fmt.Scanln(&i)
   switch i {
     case 1 :
     fmt.Printf("\nEnter %d Key and %d Value  :",i,i)
     fmt.Scanln(&k)
     fmt.Scanln(&v)
     f1.Put(k,v)
     //f1.Get(k)
     goto L1
     case 2 :
     fmt.Println("\nEnter key to find value :")
     fmt.Scanf("%d",&ke)
     fmt.Println(f1.Get(ke)) 
     default:
     fmt.Println("ERROR")
     }
}
