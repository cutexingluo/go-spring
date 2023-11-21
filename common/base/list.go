package base

import (
	"sort"
)

type List interface {
	Collection
	// Get returns the item at the given index
	Get(index int) (elem interface{}, err error)
	// Set sets the item at the given index
	Set(index int, item interface{}) (old interface{}, err error)

	// RemoveIndex removes the item at the given index
	RemoveIndex(index int) (elem interface{}, err error)

	// Insert inserts the item at the given index, like addIndex
	Insert(index int, item interface{}) (bool, error)

	// InsertAll inserts the items in the given collection at the given index
	InsertAll(index int, collection Collection) (modified bool, err error)

	// IndexOf returns the index of the given item
	IndexOf(item interface{}) int

	// LastIndexOf returns the last index of the given item
	LastIndexOf(item interface{}) int

	// SubList returns a sublist of the list
	//SubList(fromIndex, toIndex int) List

	// ForEachWithIndex iterates over the collection and passes each item to the given callback
	ForEachWithIndex(consumer func(index int, item interface{}))

	// ListIterator returns a ListIterator for the list
	ListIterator() ListIterator

	// ListIteratorIndex returns a ListIterator for the list
	ListIteratorIndex(index int) ListIterator
}

// DefList  has  the implementation of the List interface
type DefList interface {
	List
	// ReplaceAll replaces all items in the list with the result of the given unary operation
	ReplaceAll(unaryOp func(item interface{}) interface{})

	// Sort sorts the list using the given comparator
	Sort(comparator Comparator)
}

// DefaultList  has  the abstract of the DefList interface, it is like Java.util List interface
type DefaultList struct {
	List
	//DefList  // it should not be used
}

func NewDefaultListOverrideBy(list List) *DefaultList {
	return &DefaultList{
		list,
	}
}

func (_this *DefaultList) ReplaceAll(unaryOp func(item interface{}) interface{}) {
	RequireNonNil(unaryOp)
	it := _this.ListIterator()
	for it.HasNext() {
		_ = it.Set(unaryOp(it.Next())) // ignore error
	}
}

func (_this *DefaultList) Sort(comparator Comparator) {
	slice := _this.ToSlice()
	sort.Slice(slice, func(i, j int) bool {
		return comparator.Compare(slice[i], slice[j]) < 0
	})
	it := _this.ListIterator()
	for _, e := range slice {
		it.Next()
		_ = it.Set(e) // ignore error
	}
}

//------------------abstract ----------------

// AbstractList has the abstract of the List
type AbstractList struct {
	//Collection
	*AbstractCollection              // Single inheritance. 单继承
	DefList             *DefaultList // Side methods, implemented by default.  (Like the default implementation interface for Java 8) 副方法，默认实现
}

// NewAbstractListOverrideBy one is AbstractCollection and the other is DefaultList, they should be the same type
func NewAbstractListOverrideBy(collection *AbstractCollection, list *DefaultList) *AbstractList {
	return &AbstractList{
		AbstractCollection: collection,
		DefList:            list,
	}
}

func (_this *AbstractList) Iterator() Iterator {
	return NewItr(_this.DefList)
}

func (_this *AbstractList) removeRange(fromIndex, toIndex int) {
	it := _this.ListIteratorIndex(fromIndex)
	n := toIndex - fromIndex
	for i := 0; i < n; i++ {
		it.Next()
		_ = it.Remove() // no catch
	}
}
func (_this *AbstractList) Clear() {
	_this.removeRange(0, _this.Size())
}

func (_this *AbstractList) Add(item interface{}) (bool, error) {
	return _this.DefList.Insert(_this.Size(), item)
}

func (_this *AbstractList) InsertAll(index int, collection Collection) (modified bool, err error) {
	err = _this.rangeCheckForAdd(index)
	if err != nil {
		return
	}
	it := _this.Iterator()
	for it.HasNext() {
		changed, errs := _this.DefList.Insert(index, it.Next())
		if changed {
			modified = true
		}
		if errs != nil {
			err = errs
			break
		}
	}
	return
}

// IndexOf returns the index of the given item in the given collection, or -1 if not found
func (_this *AbstractList) IndexOf(item interface{}) int {
	it := _this.ListIterator()
	if obj, ok := item.(Object); ok {
		for it.HasNext() {
			if obj.Equals(it.Next()) {
				return it.NextIndex()
			}
		}
	} else {
		for it.HasNext() {
			if item == it.Next() {
				return it.NextIndex()
			}
		}
	}
	return -1
}

func (_this *AbstractList) LastIndexOf(item interface{}) int {
	it := _this.ListIteratorIndex(_this.Size())
	if obj, ok := item.(Object); ok {
		for it.HasPrevious() {
			if obj.Equals(it.Previous()) {
				return it.NextIndex()
			}
		}
	} else {
		for it.HasPrevious() {
			if item == it.Previous() {
				return it.NextIndex()
			}
		}
	}
	return -1
}

func (_this *AbstractList) ForEachWithIndex(consumer func(index int, item interface{})) {
	it := _this.ListIterator()
	for it.HasNext() {
		consumer(it.NextIndex(), it.Next())
	}
}

func (_this *AbstractList) ListIterator() ListIterator {
	return _this.ListIteratorIndex(0)
}

func (_this *AbstractList) ListIteratorIndex(index int) ListIterator {
	err := _this.rangeCheckForAdd(index)
	if err != nil {
		panic(err)
		return nil
	}
	return NewListItrInit(_this.DefList)
}

func (_this *AbstractList) rangeCheckForAdd(index int) error {
	if index < 0 || index > _this.Size() {
		return &ErrIllegalArgument{ErrMsg: "index out of range"}
	}
	return nil
}

//func (_this *AbstractList) Size() int {
//	panic("implement me")
//}
//func (_this *AbstractList) Get(index int) (elem interface{}, err error) {
//	panic("implement me")
//}
//func (_this *AbstractList) Set(index int, item interface{}) (old interface{}, err error) {
//	panic("implement me")
//}
//func (_this *AbstractList) RemoveIndex(index int) (elem interface{}, err error) {
//	panic("implement me")
//}
//func (_this *AbstractList) Insert(index int, item interface{}) (bool, error) {
//	panic("implement me")
//}

//func (_this *AbstractList) Remove(item interface{}) (bool, error) {
//	panic("implement me")
//}

//func (_this *AbstractList) RemoveAll(collection Collection) (modified bool, err error) {
//	panic("implement me")
//}
//func (_this *AbstractList) RetainAll(collection Collection) (modified bool, err error) {
//	panic("implement me")
//}
//func (_this *AbstractList) ContainsAll(collection Collection) bool {
//	panic("implement me")
//}
