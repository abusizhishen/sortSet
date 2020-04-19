package src

import "sync"

// base function add update remove size range
//addOrUpdate exists remove size topK topRange rangeRemove reverse current safe

type SortedSetInterface interface {
	Add(node Node) bool
	Update(node Node) bool
	AddOrUpdate(Node)
	AddOrUpdateMany([]Node)
	Remove(key string)
	RemoveMany(key string)
	Exists(key string)bool
	RemoveRange(start,end int)
	Size()int
	TopK(int2 int)[]Node
	TopByRange(start,end int)[]Node
}

type SortedSet struct {
	length int
	Head *Node
	Tail *Node
	dict map[string]*Node
	sync.RWMutex
}

func (s *SortedSet)AddOrUpdate(condition Condition)  {

}

type Condition struct {
	Start int
	Length int
	Reverse bool
}



