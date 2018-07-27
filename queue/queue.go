package queue

import "errors"

// Queue jo sab ke liye kam kare
type Queue []interface{}

// PutAny adds an element to the queue.
func (c *Queue) PutAny(elem interface{}) {
	*c = append(*c, elem)
}

// GetAny removes an element from the queue.
// If the queue is empty, GetAny returns an error.
func (c *Queue) GetAny() (interface{}, error) {
	//TODO: fetch the first element's value, and then remove the first element from the queue.
	//If the queue is already empty, return the zero value of interface{} and an error.
	if len(*c) > 0 {
		elem := (*c)[0]
		*c = (*c)[1:]
		return elem, nil
	}
	return 0, errors.New("POP failed")
}
