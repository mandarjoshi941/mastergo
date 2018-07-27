package intqueue

import "appliedgo/queue"

// IntegerQueue bla bla bla
type IntegerQueue struct {
	queue.Queue
}

// Get get E
func (i *IntegerQueue) Get() (int, error) {
	elem, err := i.GetAny()
	if err != nil {
		return 0, err
	}

	intelem := elem.(int)
	return intelem, err
}

// Put bla bla bla
func (i *IntegerQueue) Put(elem int) {
	i.PutAny(elem)
}
