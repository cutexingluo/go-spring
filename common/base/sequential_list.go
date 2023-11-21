package base

// AbstractSequentialList is the abstract of the AbstractList
type AbstractSequentialList struct {
	*AbstractList
}

func NewAbstractSequentialListOverrideBy(list *AbstractList) *AbstractSequentialList {
	return &AbstractSequentialList{
		AbstractList: list,
	}
}

func (_this *AbstractSequentialList) Get(index int) (elem interface{}, err error) {
	if index < 0 || index >= _this.Size() {
		return nil, &ErrOutOfRange{ErrMsg: "index out of range"}
	}
	return _this.ListIteratorIndex(index).Next(), nil
}

func (_this *AbstractSequentialList) Set(index int, item interface{}) (old interface{}, err error) {
	if index < 0 || index >= _this.Size() {
		return nil, &ErrOutOfRange{ErrMsg: "index out of range"}
	}
	it := _this.ListIteratorIndex(index)
	oldVal := it.Next()
	err = it.Set(item)
	if err != nil {
		return oldVal, err
	}
	return oldVal, nil
}

func (_this *AbstractSequentialList) Insert(index int, item interface{}) (bool, error) {
	if index < 0 || index > _this.Size() {
		return false, &ErrOutOfRange{ErrMsg: "index out of range"}
	}
	it := _this.ListIteratorIndex(index)
	err := it.Add(item)
	if err != nil {
		return false, err
	}
	return true, nil
}

func (_this *AbstractSequentialList) RemoveIndex(index int) (elem interface{}, err error) {
	if index < 0 || index >= _this.Size() {
		return nil, &ErrOutOfRange{ErrMsg: "index out of range"}
	}
	it := _this.ListIteratorIndex(index)
	elem = it.Next()
	err = it.Remove()
	if err != nil {
		return nil, err
	}
	return elem, nil
}

func (_this *AbstractSequentialList) InsertAll(index int, collection Collection) (modified bool, err error) {
	if index < 0 || index > _this.Size() {
		return false, &ErrOutOfRange{ErrMsg: "index out of range"}
	}
	it := _this.ListIteratorIndex(index)
	cit := collection.Iterator()
	for cit.HasNext() {
		err = it.Add(cit.Next())
		if err != nil {
			return false, err
		}
		modified = true
	}
	return modified, err
}

func (_this *AbstractSequentialList) Iterator() Iterator {
	return _this.ListIterator()
}
