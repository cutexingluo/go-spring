package structure

import (
	"github.com/cutexingluo/go-spring/common/base"
)

// min capacity
const minCapacity = 1 << 3

// ArrayDeque is a double-ended queue with a fixed capacity , like java.util.ArrayDeque
// ArrayDeque extends base.AbstractCollection implements base.Deque
type ArrayDeque struct {
	*base.AbstractCollection
	elem []interface{}
	//size int // number of elements
	head int // index of first element
	tail int // index + 1 of last element(exclusive)
	//cap  int // cap of array
}

func NewArrayDeque() *ArrayDeque {
	return NewArrayDequeCap(16)
}
func NewArrayDequeCap(capacity int) *ArrayDeque {
	arrayQueue := &ArrayDeque{
		elem: allocate(capacity),
		head: 0,
		tail: 0,
	}
	arrayQueue.AbstractCollection = base.NewAbstractCollectionOverrideBy(arrayQueue)
	return arrayQueue
}
func NewArrayDequeCollection(collection base.Collection) *ArrayDeque {
	arrayQueue := NewArrayDequeCap(collection.Size())
	_, err := arrayQueue.AddAll(collection)
	if err != nil {
		return nil
	}
	return arrayQueue
}

// is power of 2
func calculate(capacity int) int {
	initialCapacity := minCapacity
	if capacity >= initialCapacity {
		initialCapacity = capacity
		initialCapacity |= initialCapacity >> 1
		initialCapacity |= initialCapacity >> 2
		initialCapacity |= initialCapacity >> 4
		initialCapacity |= initialCapacity >> 8
		initialCapacity |= initialCapacity >> 16
		initialCapacity++
		if initialCapacity < 0 {
			initialCapacity >>= 1
		}
	}
	return initialCapacity
}
func allocate(capacity int) []interface{} {
	return make([]interface{}, calculate(capacity))
}
func (_this *ArrayDeque) doubleCapacity() error {
	if _this.head != _this.tail {
		return &base.ErrCallNotSatisfied{ErrMsg: "the queue is not full, should not be double capacity"}
	}
	p := _this.head
	n := len(_this.elem)
	r := n - p
	newCap := n << 1
	if newCap < 0 {
		return &base.ErrIllegalState{ErrMsg: "deque cap is too big"}
	}
	newSlice := make([]interface{}, newCap)
	base.ArrayCopy(_this.elem, p, newSlice, 0, r)
	base.ArrayCopy(_this.elem, 0, newSlice, r, p)
	_this.elem = newSlice
	_this.head = 0
	_this.tail = n
	return nil
}
func (_this *ArrayDeque) copyElem(elems []interface{}) []interface{} {
	if _this.head < _this.tail {
		base.ArrayCopy(_this.elem, _this.head, elems, 0, 1)
	} else if _this.head > _this.tail {
		headLen := len(elems) - _this.head
		base.ArrayCopy(_this.elem, _this.head, elems, 0, headLen)
		base.ArrayCopy(_this.elem, 0, elems, headLen, _this.tail)
	}
	return elems
}

// getHeadPrevious the first item's previous element
func (_this *ArrayDeque) getHeadPrevious() int {
	return (_this.head - 1) & (len(_this.elem) - 1)
}

// getHeadNext the first item's next element
func (_this *ArrayDeque) getHeadNext() int {
	return (_this.head + 1) & (len(_this.elem) - 1)
}

// getTailPrevious the last item
func (_this *ArrayDeque) getTailPrevious() int {
	return (_this.tail - 1) & (len(_this.elem) - 1)
}

// getTailNext the tail next index
func (_this *ArrayDeque) getTailNext() int {
	return (_this.tail + 1) & (len(_this.elem) - 1)
}

func (_this *ArrayDeque) AddFirst(item interface{}) {
	if item == nil {
		return
	}
	_this.head = _this.getHeadPrevious()
	_this.elem[_this.head] = item
	if _this.head == _this.tail {
		err := _this.doubleCapacity()
		if err != nil {
			return
		}
	}
}

func (_this *ArrayDeque) AddLast(item interface{}) {
	if item == nil {
		return
	}
	_this.elem[_this.tail] = item
	_this.tail = _this.getTailNext()
	if _this.tail == _this.head {
		err := _this.doubleCapacity()
		if err != nil {
			return
		}
	}
}

func (_this *ArrayDeque) OfferFirst(item interface{}) bool {
	_this.AddFirst(item)
	return true
}

func (_this *ArrayDeque) OfferLast(item interface{}) bool {
	_this.AddLast(item)
	return true
}

func (_this *ArrayDeque) RemoveFirst() (elem interface{}, err error) {
	item := _this.PollFirst()
	if item == nil {
		return nil, &base.ErrEmpty{ErrMsg: "no such item"}
	}
	return item, nil
}

func (_this *ArrayDeque) RemoveLast() (elem interface{}, err error) {
	item := _this.PollLast()
	if item == nil {
		return nil, &base.ErrEmpty{ErrMsg: "deque is empty"}
	}
	return item, nil
}

func (_this *ArrayDeque) PollFirst() interface{} {
	head := _this.head
	res := _this.elem[head]
	if res == nil {
		return nil
	}
	_this.elem[head] = nil
	_this.head = _this.getHeadNext()
	return res
}

func (_this *ArrayDeque) PollLast() interface{} {
	tail := _this.getTailPrevious()
	res := _this.elem[tail]
	if res == nil {
		return nil
	}
	_this.elem[tail] = nil
	_this.tail = tail
	return res
}

func (_this *ArrayDeque) GetFirst() (elem interface{}, err error) {
	res := _this.elem[_this.head]
	if res == nil {
		return nil, &base.ErrEmpty{ErrMsg: "deque is empty"}
	}
	return res, nil
}

func (_this *ArrayDeque) GetLast() (elem interface{}, err error) {
	res := _this.elem[_this.getTailPrevious()]
	if res == nil {
		return nil, &base.ErrEmpty{ErrMsg: "deque is empty"}
	}
	return res, nil
}

func (_this *ArrayDeque) PeekFirst() interface{} {
	return _this.elem[_this.head]
}

func (_this *ArrayDeque) PeekLast() interface{} {
	return _this.elem[_this.getTailPrevious()]
}

func (_this *ArrayDeque) checkInvariants() error {
	if _this.elem[_this.tail] != nil {
		return &base.ErrIllegalState{ErrMsg: "array deque tail is  wrong"}
	}
	var checkHead bool
	if _this.head == _this.tail {
		checkHead = _this.elem[_this.head] == nil
	} else {
		checkHead = _this.elem[_this.head] != nil &&
			_this.elem[_this.getTailPrevious()] != nil
	}
	if !checkHead {
		return &base.ErrIllegalState{ErrMsg: "array deque head and tail is  wrong"}
	}
	if _this.elem[_this.getHeadPrevious()] != nil {
		return &base.ErrIllegalState{ErrMsg: "array deque head is  wrong"}
	}
	return nil
}

func (_this *ArrayDeque) delete(index int) (bool, error) {
	err := _this.checkInvariants()
	if err != nil {
		return false, err
	}
	m := len(_this.elem) - 1
	h := _this.head
	t := _this.tail
	front := (index - h) & m
	back := (t - index) & m
	if front >= ((t - h) & m) {
		return false, &base.ErrIllegalState{ErrMsg: "it panic ConcurrentModificationException"}
	}

	if front < back {
		if h <= index {
			base.ArrayCopy(_this.elem, h, _this.elem, h+1, front)
		} else {
			base.ArrayCopy(_this.elem, 0, _this.elem, 1, index)
			_this.elem[0] = _this.elem[m]
			base.ArrayCopy(_this.elem, h, _this.elem, h+1, m-h)
		}
		_this.elem[h] = nil
		_this.head = (h + 1) & m
		return false, nil
	} else {
		if index < t {
			base.ArrayCopy(_this.elem, index+1, _this.elem, index, back)
			_this.tail = t - 1
		} else {
			base.ArrayCopy(_this.elem, index+1, _this.elem, index, m-index)
			_this.elem[m] = _this.elem[0]
			base.ArrayCopy(_this.elem, 1, _this.elem, 0, t)
			_this.tail = (t - 1) & m
		}
		return true, nil
	}
}

func (_this *ArrayDeque) RemoveFirstOccurrence(item interface{}) (bool, error) {
	if item == nil {
		return false, &base.ErrIllegalArgument{ErrMsg: "item is nil"}
	}
	m := len(_this.elem) - 1
	p := _this.head
	tar := _this.elem[p]
	for i := p; tar != nil; tar = _this.elem[i] {
		if base.Equals(item, tar) {
			b, err := _this.delete(i)
			if err != nil {
				return b, err
			}
			return b, nil
		}
		i = (i + 1) & m
	}
	return false, nil
}

func (_this *ArrayDeque) RemoveLastOccurrence(item interface{}) (bool, error) {
	if item == nil {
		return false, &base.ErrIllegalArgument{ErrMsg: "item is nil"}
	}
	m := len(_this.elem) - 1
	p := _this.getTailPrevious()
	tar := _this.elem[p]
	for i := p; tar != nil; tar = _this.elem[i] {
		if base.Equals(item, tar) {
			b, err := _this.delete(i)
			if err != nil {
				return b, err
			}
			return b, nil
		}
		i = (i - 1) & m
	}
	return false, nil
}

func (_this *ArrayDeque) Add(item interface{}) (bool, error) {
	_this.AddLast(item)
	return true, nil
}

func (_this *ArrayDeque) Offer(item interface{}) bool {
	return _this.OfferLast(item)
}

func (_this *ArrayDeque) RemoveHead() (elem interface{}, err error) {
	return _this.RemoveFirst()
}

func (_this *ArrayDeque) Poll() interface{} {
	return _this.PollFirst()
}

func (_this *ArrayDeque) Elem() (elem interface{}, err error) {
	return _this.GetFirst()
}

func (_this *ArrayDeque) Peek() interface{} {
	return _this.PeekFirst()
}

func (_this *ArrayDeque) Push(item interface{}) (err error) {
	_this.AddFirst(item)
	return
}

func (_this *ArrayDeque) Pop() (elem interface{}, err error) {
	return _this.RemoveFirst()
}

// *** Collection Methods ***

func (_this *ArrayDeque) Size() int {
	return (_this.tail - _this.head) & (len(_this.elem) - 1)
}

func (_this *ArrayDeque) IsEmpty() bool {
	return _this.head == _this.tail
}

func (_this *ArrayDeque) Contains(item interface{}) bool {
	if item == nil {
		return false
	}
	m := len(_this.elem) - 1
	p := _this.head
	tar := _this.elem[p]
	for i := p; tar != nil; tar = _this.elem[i] {
		if base.Equals(item, tar) {
			return true
		}
		i = (i + 1) & m
	}
	return false
}

func (_this *ArrayDeque) Remove(item interface{}) (bool, error) {
	return _this.RemoveFirstOccurrence(item)
}
func (_this *ArrayDeque) Clear() {
	h := _this.head
	t := _this.tail
	if h != t {
		_this.head = 0
		_this.tail = 0
		m := len(_this.elem) - 1
		for i := h; i != t; i = (i + 1) & m {
			_this.elem[i] = nil
		}
	}
}

func (_this *ArrayDeque) ToSlice() []interface{} {
	return _this.copyElem(make([]interface{}, _this.Size()))
}

func (_this *ArrayDeque) Iterator() base.Iterator {
	return NewDeqIterator(_this)
}

func (_this *ArrayDeque) DescendingIterator() base.Iterator {
	return NewDescendingIterator(_this)
}

// ---------------DeqIterator---------------

type DeqIterator struct {
	//base.Iterator
	*ArrayDeque
	cursor int
	// Tail recorded at construction (also in remove), to stop iterator and also to check for comodification
	fence   int
	lastRet int
}

func NewDeqIterator(deque *ArrayDeque) *DeqIterator {
	ret := &DeqIterator{
		ArrayDeque: deque,
		cursor:     deque.head,
		fence:      deque.tail,
		lastRet:    -1,
	}
	return ret
}

func (_this *DeqIterator) HasNext() bool {
	return _this.cursor != _this.fence
}

func (_this *DeqIterator) Next() interface{} {
	if _this.cursor == _this.fence {
		panic(&base.ErrNoSuchElem{ErrMsg: "no such element"})
	}
	res := _this.elem[_this.cursor]
	if _this.tail != _this.fence || res == nil {
		panic(&base.ErrIllegalState{ErrMsg: "ConcurrentModificationException"})
	}
	_this.lastRet = _this.cursor
	_this.cursor = (_this.cursor + 1) & (len(_this.elem) - 1)
	return res
}

func (_this *DeqIterator) Remove() error {
	if _this.lastRet < 0 {
		return &base.ErrIllegalState{ErrMsg: "no such element to remove"}
	}
	b, err := _this.delete(_this.lastRet)
	if err != nil {
		return err
	}
	if b { // deleted
		_this.cursor = (_this.cursor - 1) & (len(_this.elem) - 1)
		_this.fence = _this.tail
	}
	_this.lastRet = -1
	return nil
}

func (_this *DeqIterator) ForEachRemaining(consumer func(item interface{})) {
	base.RequireNonNil(consumer)
	a := _this.elem
	m := len(a) - 1
	f := _this.fence
	i := _this.cursor
	_this.cursor = f
	for i != f {
		e := a[i]
		i = (i + 1) & m
		if e == nil {
			panic(&base.ErrIllegalState{ErrMsg: "ConcurrentModificationException"})
		}
		consumer(a[i])
	}
}

//---------------DescendingIterator---------------

// DescendingIterator is a reverse iterator for DeqIterator
// full methods
type DescendingIterator struct {
	//base.Iterator
	*ArrayDeque
	*base.AbstractIterator
	cursor int
	// Tail recorded at construction (also in remove), to stop iterator and also to check for comodification
	fence   int
	lastRet int
}

func NewDescendingIterator(deque *ArrayDeque) *DescendingIterator {
	ret := &DescendingIterator{
		ArrayDeque: deque,
		cursor:     deque.tail,
		fence:      deque.head,
		lastRet:    -1,
	}
	ret.AbstractIterator = base.NewAbstractIteratorOverrideBy(ret)
	return ret
}

func (_this *DescendingIterator) HasNext() bool {
	return _this.cursor != _this.fence
}

func (_this *DescendingIterator) Next() interface{} {
	if _this.cursor == _this.fence {
		panic(&base.ErrNoSuchElem{ErrMsg: "no such element"})
	}
	_this.cursor = (_this.cursor - 1) & (len(_this.elem) - 1)
	res := _this.elem[_this.cursor]
	if _this.tail != _this.fence || res == nil {
		panic(&base.ErrIllegalState{ErrMsg: "ConcurrentModificationException"})
	}
	_this.lastRet = _this.cursor
	return res
}

func (_this *DescendingIterator) Remove() error {
	if _this.lastRet < 0 {
		return &base.ErrIllegalState{ErrMsg: "no such element to remove"}
	}
	b, err := _this.delete(_this.lastRet)
	if err != nil {
		return err
	}
	if !b { // not deleted
		_this.cursor = (_this.cursor + 1) & (len(_this.elem) - 1)
		_this.fence = _this.head
	}
	_this.lastRet = -1
	return nil
}
