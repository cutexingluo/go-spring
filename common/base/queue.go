package base

type Queue interface {
	// Collection extends the Collection interface
	// Collection.Add adds a new item to the queue. Returns an error if the item is not added.
	Collection

	// Offer adds a new item to the queue. Will not throw exceptions
	Offer(item interface{}) bool

	// RemoveHead removes the first item from the queue. Returns an error if the queue is empty.
	RemoveHead() (elem interface{}, err error)

	// Poll removes the first item from the queue. Returns nil if the queue is empty.
	Poll() interface{}

	// Elem returns the first item from the queue. Returns an error if the queue is empty.
	Elem() (elem interface{}, err error)

	// Peek returns the first item from the queue. Returns nil if the queue is empty.
	Peek() interface{}
}

//------------------abstract ----------------

// AbstractQueue is abstract of  the Queue interface
type AbstractQueue struct {
	Queue               // interface
	*AbstractCollection //struct
}

func NewAbstractQueueOverrideBy(queue Queue, collection *AbstractCollection) *AbstractQueue {
	return &AbstractQueue{
		Queue:              queue,
		AbstractCollection: collection,
	}
}

func (_this *AbstractQueue) Add(item interface{}) (bool, error) {
	if _this.Offer(item) {
		return true, nil
	} else {
		return false, &ErrFull{ErrMsg: "queue is full"}
	}
}

//func (_this *AbstractQueue) Offer(item interface{}) bool {
//	panic("implement me")
//}

func (_this *AbstractQueue) RemoveFirst() (elem interface{}, err error) {
	ret := _this.Poll()
	if ret != nil {
		return ret, nil
	} else {
		return nil, &ErrEmpty{ErrMsg: "queue is empty"}
	}
}

//func (_this *AbstractQueue) Poll() interface{} {
//	panic("implement me")
//}

func (_this *AbstractQueue) Elem() (elem interface{}, err error) {
	ret := _this.Peek()
	if ret != nil {
		return ret, nil
	} else {
		return nil, &ErrNoSuchElem{ErrMsg: "queue have no such element"}
	}
}

//func (_this *AbstractQueue) Peek() interface{} {
//	panic("implement me")
//}

// Clear all elements
func (_this *AbstractQueue) Clear() {
	for _this.Poll() != nil {
	}
}

// AddAll adds all elements
func (_this *AbstractQueue) AddAll(collection Collection) (modified bool, err error) {
	RequireNonNil(collection)
	if EqualsAddress(_this, collection) {
		return false, &ErrIllegalArgument{ErrMsg: "can not add self"}
	}
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
