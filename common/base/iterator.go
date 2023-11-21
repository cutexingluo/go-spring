package base

// ---------------List---------------

// ListIterator iterates over the list
type ListIterator interface {
	Iterator
	// HasPrevious returns true if the iterator has previous
	HasPrevious() bool
	// Previous returns the previous item
	Previous() interface{}
	// NextIndex returns the next index
	NextIndex() int
	// PreviousIndex returns the previous index
	PreviousIndex() int
	// Set sets the item
	Set(item interface{}) error
	// Add adds the item
	Add(item interface{}) error
}

// ListItr iterates over the list
// implements ListIterator
// full methods
type ListItr struct {
	*Itr
}

func NewListItr(itr *Itr) *ListItr {
	return &ListItr{
		Itr: itr,
	}
}
func NewListItrInit(list List) *ListItr {
	listItr := &ListItr{
		Itr: NewItr(list),
	}
	return listItr
}
func NewListItrInitIndex(list List, index int) *ListItr {
	listItr := &ListItr{
		Itr: NewItr(list),
	}
	listItr.cursor = index
	return listItr
}

func (_this *ListItr) HasPrevious() bool {
	return _this.cursor != 0
}

func (_this *ListItr) Previous() interface{} {
	i := _this.cursor - 1
	previous, err := _this.list.Get(i) // in panic outOfRange
	if err != nil {
		panic(err)
	}
	_this.cursor = i
	_this.lastRet = i
	return previous
}

func (_this *ListItr) NextIndex() int {
	return _this.cursor
}

func (_this *ListItr) PreviousIndex() int {
	return _this.cursor - 1
}

func (_this *ListItr) Set(item interface{}) error {
	_, err := _this.list.Set(_this.cursor, item)
	if err != nil {
		return err
	}
	return nil
}

func (_this *ListItr) Add(item interface{}) error {
	i := _this.cursor
	_, err := _this.list.Insert(i, item)
	if err != nil {
		return err
	}
	_this.lastRet = -1
	_this.cursor = i + 1
	return nil
}

// DescendingIterator  iterates over the list in reverse order
// full methods
type DescendingIterator struct {
	*ListItr
}

func NewDescendingIterator(listItr *ListItr) *DescendingIterator {
	listItr.cursor = listItr.list.Size() - 1
	ret := &DescendingIterator{
		ListItr: listItr,
	}
	return ret
}

func NewDescendingIteratorInit(list List) *DescendingIterator {
	ret := &DescendingIterator{
		ListItr: NewListItrInit(list),
	}
	ret.cursor = ret.list.Size() - 1
	return ret
}
