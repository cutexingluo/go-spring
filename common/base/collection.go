package base

type Collection interface {
	Object
	Iterable

	// Size returns the number of items in the collection.
	Size() int

	// IsEmpty returns true if the collection is empty.
	IsEmpty() bool // AbstractCollection

	// Contains checks whether the collection contains the given item.
	Contains(item interface{}) bool // AbstractCollection

	// ContainsAll returns true if the collection contains
	// all the items in the given collection.
	ContainsAll(collection Collection) bool // AbstractCollection

	// ToSlice  converts the collection to a slice.
	ToSlice() []interface{} // AbstractCollection

	// Clear removes all items from the collection.
	Clear()

	// Add adds an item to the collection. Returns an error if the item is not added.
	Add(item interface{}) (bool, error)

	// AddAll adds all the items in the given collection to the collection. return true if changed,  Returns an error if the items are not added.
	AddAll(collection Collection) (modified bool, err error) // AbstractCollection

	// Remove removes the given item from the collection. Returns an error if the item is not removed.
	Remove(item interface{}) (bool, error) // AbstractCollection

	// RemoveIf removes all the items that satisfy the given predicate.
	RemoveIf(predicate func(item interface{}) bool) (modified bool, err error) // AbstractCollection

	// RemoveAll removes all the items in the given collection.return true if changed,    Returns an error if the items are not removed.
	RemoveAll(collection Collection) (modified bool, err error) // AbstractCollection

	// RetainAll removes all the items that are not in the given collection. return true if changed,   Returns an error if the items are not removed.
	RetainAll(collection Collection) (modified bool, err error) // AbstractCollection
}

//------------------abstract ----------------

// AbstractCollection implements the Collection interface ,
type AbstractCollection struct {
	Collection
	//size int // Size of the collection
}

// NewAbstractCollectionOverrideBy creates a new AbstractCollection
func NewAbstractCollectionOverrideBy(collection Collection) *AbstractCollection {
	return &AbstractCollection{
		collection,
	}
}

func (_this *AbstractCollection) Equals(i interface{}) bool {
	return EqualsObject(_this, i)
}

//func (_this *AbstractCollection) Iterator() Iterator {
//	panic("implement me")
//}

func (_this *AbstractCollection) ForEach(consumer func(item interface{})) {
	Iterator := _this.Iterator()
	Iterator.ForEachRemaining(consumer)
}

//func (_this *AbstractCollection) Size() int {
//	//return _this.size
//	panic("implement me")
//}

func (_this *AbstractCollection) IsEmpty() bool {
	return _this.Size() == 0
}

func (_this *AbstractCollection) Contains(item interface{}) bool {
	it := _this.Iterator()
	if obj, ok := item.(Object); ok {
		for it.HasNext() {
			if obj.Equals(it.Next()) {
				return true
			}
		}
	} else {
		for it.HasNext() {
			if item == it.Next() {
				return true
			}
		}
	}

	return false
}

func (_this *AbstractCollection) ContainsAll(collection Collection) bool {
	RequireNonNil(collection)
	it := collection.Iterator()
	for it.HasNext() {
		if !_this.Contains(it.Next()) {
			return false
		}
	}
	return true
}

func (_this *AbstractCollection) ToSlice() []interface{} {
	slice := make([]interface{}, 0, _this.Size()) // 先保证容量
	it := _this.Iterator()
	for it.HasNext() {
		slice = append(slice, it.Next())
	}
	return slice
}

//
//func (_this *AbstractCollection) Clear() {
//	_this.size = 0
//}

//func (_this *AbstractCollection) Add(item interface{}) (bool, error) {
//	panic("implement me")
//}

func (_this *AbstractCollection) AddAll(collection Collection) (modified bool, err error) {
	RequireNonNil(collection)
	it := collection.Iterator()
	for it.HasNext() {
		changed, errs := _this.Add(it.Next())
		if changed {
			modified = true
		}
		if errs != nil {
			err = errs
			break
		}
	}
	return modified, err
}

func (_this *AbstractCollection) Remove(item interface{}) (bool, error) {
	it := _this.Iterator()
	if obj, ok := item.(Object); ok {
		for it.HasNext() {
			if obj.Equals(it.Next()) {
				err := it.Remove()
				if err != nil {
					return false, err
				}
				return true, err
			}
		}
	} else {
		for it.HasNext() {
			if item == it.Next() {
				err := it.Remove()
				if err != nil {
					return false, err
				}
				return true, err
			}
		}
	}
	return false, nil
}

func (_this *AbstractCollection) RemoveIf(predicate func(item interface{}) bool) (modified bool, err error) {
	RequireNonNil(predicate)
	it := _this.Iterator()
	for it.HasNext() {
		if predicate(it.Next()) {
			err := it.Remove()
			if err != nil {
				return modified, err
			}
			modified = true
		}
	}
	return modified, err
}

func (_this *AbstractCollection) RemoveAll(collection Collection) (modified bool, err error) {
	RequireNonNil(collection)
	it := _this.Iterator()
	for it.HasNext() {
		if collection.Contains(it.Next()) {
			err := it.Remove()
			if err != nil {
				return modified, err
			}
			modified = true
		}
	}
	return modified, err
}

func (_this *AbstractCollection) RetainAll(collection Collection) (modified bool, err error) {
	RequireNonNil(collection)
	it := _this.Iterator()
	for it.HasNext() {
		if !collection.Contains(it.Next()) {
			err := it.Remove()
			if err != nil {
				return modified, err
			}
			modified = true
		}
	}
	return modified, err
}
