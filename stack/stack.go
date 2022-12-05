// Package stack provides a LIFO algorithm implemetation
package stack

// Stack interface represents the common operations of LIFO object
type Stack interface {
	IsEmpty() bool
	IsFull() bool
	Push(interface{}) bool
	Pop() interface{}
	Peek() interface{}
	Remaining() int
}

// Stack_t implements the Stack interface
type Stack_t struct {
	pool     []interface{}
	capacity int
	top      int
}

// function New initializes a new Queue object
// It will return nil in case of a error or a initialized Queue object
func New(size int) (p *Stack_t) {
	if size <= 0 {
		return nil
	} else {
		return &Stack_t{
			pool:     make([]interface{}, size),
			capacity: size,
			top:      -1,
		}
	}
}

// IsEmpty method check if a given Stack object is empty
func (p *Stack_t) IsEmpty() bool {
	if p.top == -1 {
		return true
	}
	return false
}

// IsFull method check if a given Stack object is full
func (p *Stack_t) IsFull() bool {
	if p.top == p.capacity-1 {
		return true
	}
	return false
}

// Push method adds a given v value to the Stack object.
// This method returns if the value was successfully added or not
func (p *Stack_t) Push(v interface{}) bool {
	if p.IsFull() {
		return false
	} else {
		p.top++
		p.pool[p.top] = v
		return true
	}
}

// Pop method removes a value from the Stack object.
// This method returns nil, in case of a empty Stack, or a value
func (p *Stack_t) Pop() (v interface{}) {
	if p.IsEmpty() {
		v = nil
	} else {
		v = p.pool[p.top]
		p.top--
	}
	return
}

// Peek method returns the next value to be poped.
// This method returns nil, in case of a empty Stack, or a value
func (p *Stack_t) Peek() (v interface{}) {
	if p.IsEmpty() {
		v = nil
	} else {
		v = p.pool[p.top]
	}
	return
}

// Remaining method returns how many free slots are left in the Stack object
func (p *Stack_t) Remaining() (cnt int) {
	if p.IsEmpty() {
		return p.capacity
	} else {
		for inter := p.top; inter < p.capacity-1; inter++ {
			cnt++
		}
	}
	return
}
