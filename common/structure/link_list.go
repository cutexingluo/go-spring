package structure

import (
	"fmt"
	"github.com/cutexingluo/go-spring/common/base"
)

// GetFirstNode test
func GetFirstNode(linkList *LinkList) *Node {
	return linkList.first
}

// GetLastNode test
func GetLastNode(linkList *LinkList) *Node {
	return linkList.last
}

// Node  the link list node
type Node struct {
	Item interface{}
	Next *Node
	Prev *Node
}

func NewNode(prev *Node, item any, next *Node) *Node {
	return &Node{
		Item: item,
		Next: next,
		Prev: prev,
	}
}

// LinkList the linked list like java LinkedList.
// LinkList extends base.AbstractSequentialList  implement base.Deque
type LinkList struct {
	//base.Cloneable
	//base.Deque
	*base.AbstractSequentialList
	size     int
	first    *Node
	last     *Node
	ModCount int
}

func NewLinkList() *LinkList {
	linkList := &LinkList{}
	// inherit from base.AbstractSequentialList
	DefList := base.NewDefaultListOverrideBy(linkList)
	AbstractCollection := base.NewAbstractCollectionOverrideBy(linkList)
	AbstractList := base.NewAbstractListOverrideBy(AbstractCollection, DefList)
	linkList.AbstractSequentialList = base.NewAbstractSequentialListOverrideBy(AbstractList)
	return linkList
}

func NewLinkListByCollection(collection base.Collection) *LinkList {
	linkList := &LinkList{}
	_, err := linkList.AddAll(collection)
	if err != nil {
		return nil
	}
	return linkList
}

// linkFirst  link the first node
func (_this *LinkList) linkFirst(item interface{}) {
	f := _this.first
	newNode := NewNode(nil, item, f)
	_this.first = newNode
	if f == nil {
		_this.last = newNode
	} else {
		f.Prev = newNode
	}
	_this.size++
	_this.ModCount++
}

// linkLast  link the last node
func (_this *LinkList) linkLast(item interface{}) {
	l := _this.last
	newNode := NewNode(l, item, nil)
	_this.last = newNode
	if l == nil {
		_this.first = newNode
	} else {
		l.Next = newNode
	}
	_this.size++
	_this.ModCount++
}

// linkBefore  Inserts element e before non-null Node curr.
func (_this *LinkList) linkBefore(item interface{}, curr *Node) {
	pred := curr.Prev
	curr.Prev = NewNode(pred, item, curr)
	if pred == nil {
		_this.first = curr.Prev
	} else {
		pred.Next = curr.Prev
	}
	_this.size++
	_this.ModCount++
}

func (_this *LinkList) unlinkFirst(f *Node) (oldElem interface{}) {
	if f == nil {
		return nil
	}
	elem := f.Item
	next := f.Next
	f.Item = nil
	f.Next = nil //GC
	_this.first = next
	if next == nil {
		_this.last = nil
	} else {
		next.Prev = nil
	}
	_this.size--
	_this.ModCount++
	return elem
}
func (_this *LinkList) unlinkLast(l *Node) (oldElem interface{}) {
	if l == nil {
		return nil
	}
	elem := l.Item
	prev := l.Prev
	l.Item = nil
	l.Prev = nil //GC
	_this.last = prev
	if prev == nil {
		_this.first = nil
	} else {
		prev.Next = nil
	}
	_this.size--
	_this.ModCount++
	return elem
}
func (_this *LinkList) unLink(elem *Node) (oldElem interface{}) {
	if elem == nil {
		return nil
	}
	oldElem = elem.Item
	prev := elem.Prev
	next := elem.Next
	if prev == nil {
		_this.first = next
	} else {
		prev.Next = next
		elem.Prev = nil
	}
	if next == nil {
		_this.last = prev
	} else {
		next.Prev = prev
		elem.Next = nil
	}
	elem.Item = nil
	_this.size--
	_this.ModCount++
	return oldElem
}

func (_this *LinkList) GetFirst() (elem interface{}, err error) {
	if _this.first == nil {
		return nil, &base.ErrNoSuchElem{ErrMsg: "no such element"}
	}
	return _this.first.Item, nil
}

func (_this *LinkList) GetLast() (elem interface{}, err error) {
	if _this.last == nil {
		return nil, &base.ErrNoSuchElem{ErrMsg: "no such element"}
	}
	return _this.last.Item, nil
}

func (_this *LinkList) RemoveFirst() (elem interface{}, err error) {
	if _this.first == nil {
		return nil, &base.ErrNoSuchElem{ErrMsg: "no such element"}
	} else {
		return _this.unlinkFirst(_this.first), nil
	}
}

func (_this *LinkList) RemoveLast() (elem interface{}, err error) {
	if _this.last == nil {
		return nil, &base.ErrNoSuchElem{ErrMsg: "no such element"}
	} else {
		return _this.unlinkLast(_this.last), nil
	}
}

func (_this *LinkList) AddFirst(item interface{}) {
	_this.linkFirst(item)
}

func (_this *LinkList) AddLast(item interface{}) {
	_this.linkLast(item)
}

func (_this *LinkList) Contains(item interface{}) bool {
	return _this.IndexOf(item) != -1
}

func (_this *LinkList) Size() int {
	return _this.size
}

func (_this *LinkList) Add(item interface{}) (isAdd bool, err error) {
	_this.linkLast(item)
	return true, nil
}

// Remove Overrides base.AbstractCollectionRemove
func (_this *LinkList) Remove(item interface{}) (bool, error) {
	if obj, ok := item.(base.Object); ok {
		for x := _this.first; x != nil; x = x.Next {
			if obj.Equals(x.Item) {
				_this.unLink(x)
				return true, nil
			}
		}
	} else {
		for x := _this.first; x != nil; x = x.Next {
			if item == x.Item {
				_this.unLink(x)
				return true, nil
			}
		}
	}
	return false, &base.ErrNoSuchElem{ErrMsg: "no such element"}
}

func (_this *LinkList) AddAll(collection base.Collection) (bool, error) {
	return _this.InsertAll(_this.size, collection)
}

func (_this *LinkList) InsertAll(index int, collection base.Collection) (modified bool, err error) {
	err = _this.checkPositionIndex(index)
	if err != nil {
		return false, err
	}
	c := collection.ToSlice()
	size := len(c)
	if size == 0 {
		return false, nil
	}
	var pred, curr *Node
	if index == _this.size {
		curr = nil
		pred = _this.last
	} else {
		curr = _this.node(index)
		pred = curr.Prev
	}
	for i := 0; i < size; i++ {
		newNode := NewNode(pred, c[i], nil)
		if pred == nil {
			_this.first = newNode
		} else {
			pred.Next = newNode
		}
		pred = newNode
	}
	if curr == nil {
		_this.last = pred
	} else {
		pred.Next = curr
		curr.Prev = pred
	}
	_this.size += size
	_this.ModCount++
	return true, nil
}
func (_this *LinkList) isPositionIndex(index int) bool {
	return index >= 0 && index <= _this.size
}
func (_this *LinkList) isElementIndex(index int) bool {
	return index >= 0 && index < _this.size
}
func (_this *LinkList) checkPositionIndex(index int) error {
	if !_this.isPositionIndex(index) {
		sprintf := fmt.Sprintf("index: %d, size: %d", index, _this.size)
		return &base.ErrOutOfRange{ErrMsg: "index out of bounds. " + sprintf}
	}
	return nil
}
func (_this *LinkList) checkElementIndex(index int) error {
	if !_this.isElementIndex(index) {
		sprintf := fmt.Sprintf("index: %d, size: %d", index, _this.size)
		return &base.ErrOutOfRange{ErrMsg: "index out of bounds. " + sprintf}
	}
	return nil
}

// Returns the (non-null) Node at the specified element index.
func (_this *LinkList) node(index int) *Node {
	err := _this.checkElementIndex(index)
	if err != nil {
		panic(err)
		return nil
	}
	var curr *Node
	if index < _this.size>>1 {
		curr = _this.first
		for i := 0; i < index; i++ {
			curr = curr.Next
		}
	} else {
		curr = _this.last
		for i := _this.size - 1; i > index; i-- {
			curr = curr.Prev
		}
	}
	return curr
}

func (_this *LinkList) Clear() {
	for x := _this.first; x != nil; {
		next := x.Next
		x.Item = nil
		x.Next = nil
		x.Prev = nil
		x = next
	}
	_this.first = nil
	_this.last = nil
	_this.size = 0
	_this.ModCount++
}

func (_this *LinkList) Get(index int) (elem interface{}, err error) {
	err = _this.checkElementIndex(index)
	if err != nil {
		return nil, err
	}
	return _this.node(index).Item, nil
}

func (_this *LinkList) Set(index int, item interface{}) (old interface{}, err error) {
	err = _this.checkElementIndex(index)
	if err != nil {
		return nil, err
	}
	tar := _this.node(index)
	old = tar.Item
	tar.Item = item
	return old, nil
}

func (_this *LinkList) Insert(index int, item interface{}) (bool, error) {
	err := _this.checkPositionIndex(index)
	if err != nil {
		return false, err
	}
	if index == _this.size {
		_this.linkLast(item)
	} else {
		_this.linkBefore(item, _this.node(index))
	}
	return true, nil
}
func (_this *LinkList) RemoveIndex(index int) (elem interface{}, err error) {
	err = _this.checkElementIndex(index)
	if err != nil {
		return nil, err
	}
	return _this.unLink(_this.node(index)), nil
}

// Search Operations

func (_this *LinkList) IndexOf(item interface{}) int {
	index := 0
	if obj, ok := item.(base.Object); ok {
		for x := _this.first; x != nil; x = x.Next {
			if obj.Equals(x.Item) {
				return index
			}
			index++
		}
	} else {
		for x := _this.first; x != nil; x = x.Next {
			if item == x.Item {
				return index
			}
			index++
		}
	}
	return -1
}

func (_this *LinkList) LastIndexOf(item interface{}) int {
	index := _this.size
	if obj, ok := item.(base.Object); ok {
		for x := _this.last; x != nil; x = x.Prev {
			index--
			if obj.Equals(x.Item) {
				return index
			}
		}
	} else {
		for x := _this.last; x != nil; x = x.Prev {
			index--
			if item == x.Item {
				return index
			}
		}
	}
	return -1
}

// Queue operations.

func (_this *LinkList) Peek() interface{} {
	return _this.PeekFirst()
}

func (_this *LinkList) Elem() (elem interface{}, err error) {
	return _this.GetFirst()
}

func (_this *LinkList) Poll() interface{} {
	return _this.PollFirst()
}

func (_this *LinkList) RemoveHead() (elem interface{}, err error) {
	return _this.RemoveFirst()
}

func (_this *LinkList) Offer(item interface{}) bool {
	add, err := _this.Add(item)
	if err != nil {
		return false
	}
	return add
}

func (_this *LinkList) OfferFirst(item interface{}) bool {
	_this.AddFirst(item)
	return true
}

func (_this *LinkList) OfferLast(item interface{}) bool {
	_this.AddLast(item)
	return true
}

func (_this *LinkList) PeekFirst() interface{} {
	f := _this.first
	if f != nil {
		return f.Item
	} else {
		return nil
	}
}

func (_this *LinkList) PeekLast() interface{} {
	l := _this.last
	if l != nil {
		return _this.unlinkLast(l)
	} else {
		return nil
	}
}

func (_this *LinkList) PollFirst() interface{} {
	f := _this.first
	if f != nil {
		return _this.unlinkFirst(f)
	} else {
		return nil
	}
}

func (_this *LinkList) PollLast() interface{} {
	l := _this.last
	if l != nil {
		return _this.unlinkLast(l)
	} else {
		return nil
	}
}

func (_this *LinkList) Push(item interface{}) (err error) {
	_this.AddFirst(item)
	return nil
}

func (_this *LinkList) Pop() (elem interface{}, err error) {
	return _this.RemoveFirst()
}

func (_this *LinkList) RemoveFirstOccurrence(item interface{}) (bool, error) {
	return _this.Remove(item)
}

func (_this *LinkList) RemoveLastOccurrence(item interface{}) (bool, error) {
	if obj, ok := item.(base.Object); ok {
		for x := _this.last; x != nil; x = x.Prev {
			if obj.Equals(x.Item) {
				_this.unLink(x)
				return true, nil
			}
		}
	} else {
		for x := _this.last; x != nil; x = x.Prev {
			if item == x.Item {
				_this.unLink(x)
				return true, nil
			}
		}
	}
	return false, &base.ErrNoSuchElem{ErrMsg: "no such element"}
}

func (_this *LinkList) ListIteratorIndex(index int) base.ListIterator {
	err := _this.checkPositionIndex(index)
	if err != nil {
		return nil
	}
	return base.NewListItrInitIndex(_this, index)
}

func (_this *LinkList) DescendingIterator() base.Iterator {
	return base.NewDescendingIteratorInit(_this)
}

func (_this *LinkList) Clone() interface{} {
	clone := NewLinkList()
	clone.last = nil
	clone.first = nil
	clone.size = 0
	clone.ModCount = 0
	for x := _this.first; x != nil; x = x.Next {
		_, err := clone.Add(x.Item)
		if err != nil {
			return nil
		}
	}
	return clone
}
func (_this *LinkList) ToSlice() []interface{} {
	slice := make([]interface{}, 0, _this.size)
	for x := _this.first; x != nil; x = x.Next {
		slice = append(slice, x.Item)
	}
	return slice
}

func (_this *LinkList) String() string {
	str := "LinkList"
	str += "["
	for x := _this.first; x != nil; x = x.Next {
		str += fmt.Sprintf("%v", x.Item)
		if x.Next != nil {
			str += ","
		}
	}
	str += "]"
	return str
}

func (_this *LinkList) ForEach(consumer func(item interface{})) {
	Iterator := _this.Iterator()
	Iterator.ForEachRemaining(consumer)
}
