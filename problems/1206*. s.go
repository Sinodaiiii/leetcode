package problems

const maxLevel = 32
const pFactor = 0.25

type SkipListNode struct {
	next  []*SkipListNode
	value int
}

type SkipList struct {
	head  *SkipListNode
	level int
}

func Constructor() SkipList {
	NewList := SkipList{&SkipListNode{make([]*SkipListNode, maxLevel), -1}, 0}
	return NewList
}

func (this *SkipList) Search(target int) bool {

}

func (s *SkipList) Add(num int) {
	var tail = make([]*SkipListNode, maxLevel)
	for i := range tail {
		tail[i].next[i] = s.head
		tail[i].value = -1
	}
	for i := s.level - 1; i >= 0; i-- {
		for tail[i].value < num && tail[i].next[i] != nil {
			tail[i] = tail[i].next[i]
		}
	}

}

func (this *SkipList) Erase(num int) bool {

}
