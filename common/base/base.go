package base

// Cloneable is the interface that wraps the Clone method.
type Cloneable interface {
	// Clone clones the object.
	Clone() interface{}
}

// Object is the base interface that like Java Object
type Object interface { // extends Mashibin implement Yushengjun  qaq lol
	// String returns the string representation
	//String() string

	// Equals returns true if the object is equal to the other
	Equals(interface{}) bool
}

// Objects is the base struct that like Java Object
// full methods
type Objects struct { // extends Mashibin implement Yushengjun  qaq lol
	//Object
}

func NewObjects() *Objects {
	return &Objects{}
}

// Equals returns true if the object is equal to the other, including EqualsAddress and reflect.DeepEqual
func (o *Objects) Equals(i interface{}) bool {
	return EqualsObject(o, i) // compare by address and deep equal
}

// Iterable  is the interface that can use the Iterator
type Iterable interface {
	// Iterator returns a new iterator
	Iterator() Iterator

	// ForEach iterates the collection.
	ForEach(consumer func(item interface{}))
}

// Iterator is the interface that wraps the HasNext and Next methods.
type Iterator interface {
	// HasNext checks whether the collection has next.
	HasNext() bool

	// Next returns the next item in the collection.
	Next() interface{}

	// Remove removes the item from the collection
	Remove() error

	// ForEachRemaining iterates the collection.
	ForEachRemaining(consumer func(item interface{}))
}

// Comparable is the interface that wraps the CompareTo method.
type Comparable interface {
	//Object
	//comparable

	// CompareTo return less than 0 if o1 < o2, 0 if o1 == o2, greater than 0 if o1 > o2
	CompareTo(o interface{}) int
}

// Comparator is the interface that wraps the Compare method.
// You can use IComparator that implemented the Compare method.
type Comparator interface {
	// Compare return less than 0 if o1 < o2, 0 if o1 == o2, greater than 0 if o1 > o2
	Compare(o1, o2 interface{}) int
}

// IComparator is the struct implemented the Compare method.
//
// e.g.
//
//	IComparator(func(o1, o2 interface{}) int {
//					return o1.(int) - o2.(int)
//	 })
type IComparator func(o1, o2 interface{}) int

func (f IComparator) Compare(o1, o2 interface{}) int {
	return f(o1, o2)
}

// AbstractIterable is the interface that wraps the Iterator interface
type AbstractIterable struct {
	Iterable
}

// NewAbstractIterableOverrideBy creates a new AbstractIterable
func NewAbstractIterableOverrideBy(iterable Iterable) *AbstractIterable {
	return &AbstractIterable{
		Iterable: iterable,
	}
}

func (_this *AbstractIterable) ForEach(consumer func(item interface{})) {
	iterator := _this.Iterator()
	iterator.ForEachRemaining(consumer)
}

// AbstractIterator is the abstract implemented Iterator interface
type AbstractIterator struct {
	Iterator
}

//func NewAbstractIterator(iterator Iterator) *AbstractIterator {
//	return &AbstractIterator{
//		Iterator: iterator,
//	}
//}

// NewAbstractIteratorOverrideBy override method
func NewAbstractIteratorOverrideBy(iterator Iterator) *AbstractIterator {
	return &AbstractIterator{
		Iterator: iterator,
	}
}

func (_this *AbstractIterator) ForEachRemaining(consumer func(item interface{})) {
	RequireNonNil(consumer)
	for _this.HasNext() {
		consumer(_this.Next())
	}
}

//--------------implementation ------------------

// Itr is the struct implemented Iterator
// full methods
type Itr struct {
	*AbstractIterator
	list    List
	cursor  int // index of current item
	lastRet int
}

// NewItr returns an iterator
func NewItr(list List) *Itr {
	itr := &Itr{
		list:    list,
		cursor:  0,
		lastRet: -1,
	}
	//override AbstractIterator
	itr.AbstractIterator = NewAbstractIteratorOverrideBy(itr) //   Reverse inheritance, a feature of go
	return itr
}

func (_this *Itr) HasNext() bool {
	return _this.cursor != _this.list.Size()
}

func (_this *Itr) Next() interface{} {
	i := _this.cursor
	next, err := _this.list.Get(i) // in panic
	if err != nil {
		panic(err)
	}
	_this.lastRet = i
	_this.cursor = i + 1
	return next
}

func (_this *Itr) Remove() error {
	if _this.lastRet < 0 {
		return &ErrIllegalArgument{ErrMsg: "no last element"}
	} else {
		_, err := _this.list.RemoveIndex(_this.lastRet)
		if err != nil {
			return err
		}
		if _this.lastRet < _this.cursor {
			_this.cursor--
		}
		_this.lastRet = -1
		return nil
	}
}
