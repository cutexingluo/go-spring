package base

type Deque interface {
	Queue
	// AddFirst adds a new item to the head of this deque.
	AddFirst(item interface{})
	// AddLast adds a new item to the tail of this deque.
	AddLast(item interface{})
	// OfferFirst adds a new item to the head of this deque.
	// Returns true if the item was added, false otherwise.
	OfferFirst(item interface{}) bool
	// OfferLast adds a new item to the tail of this deque.
	// Returns true if the item was added, false otherwise.
	OfferLast(item interface{}) bool
	// RemoveFirst removes and returns the head of this deque.
	// Returns nil,err if this deque is empty.
	RemoveFirst() (elem interface{}, err error)
	// RemoveLast removes and returns the tail of this deque.
	// Returns nil,err if this deque is empty.
	RemoveLast() (elem interface{}, err error)
	// PollFirst removes and returns the head of this deque.
	// Returns nil if this deque is empty.
	PollFirst() interface{}
	// PollLast removes and returns the tail of this deque.
	// Returns nil if this deque is empty.
	PollLast() interface{}
	// GetFirst returns the head of this deque without removing it.
	// Returns nil,err if this deque is empty.
	GetFirst() (elem interface{}, err error)
	// GetLast returns the tail of this deque without removing it.
	// Returns nil,err if this deque is empty.
	GetLast() (elem interface{}, err error)
	// PeekFirst returns the head of this deque without removing it.
	// Returns nil if this deque is empty.
	PeekFirst() interface{}
	// PeekLast returns the tail of this deque without removing it.
	// Returns nil if this deque is empty.
	PeekLast() interface{}
	// RemoveFirstOccurrence removes the first occurrence of the specified element from this deque.
	// If the deque does not contain the element, it is unchanged.
	// Returns true if this deque contained the specified element (or equivalently, if this deque changed as a result of the call).
	RemoveFirstOccurrence(item interface{}) (bool, error)
	// RemoveLastOccurrence removes the last occurrence of the specified element from this deque.
	// If the deque does not contain the element, it is unchanged.
	// Returns true if this deque contained the specified element (or equivalently, if this deque changed as a result of the call).
	RemoveLastOccurrence(item interface{}) (bool, error)

	// Push adds a new item to the tail of this deque.
	Push(item interface{}) (err error)

	// Pop removes and returns the tail of this deque.
	// Returns nil,err if this deque is empty.
	Pop() (elem interface{}, err error)

	// DescendingIterator returns an iterator over the elements in this deque in reverse sequential order.
	DescendingIterator() Iterator
}
