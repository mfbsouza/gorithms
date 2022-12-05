// Package circularqueue provides a Circular Buffer
// implementation following the FIFO algorithm
package circularqueue

// Queue interface represents the common operations
// of a circular queue
type Queue interface {
	Empty() bool
	Full() bool
	Enqueue(interface{}) bool
	Dequeue() interface{}
	Peek() interface{}
	Remaining() int
}

// Queue_t implements the Queue interface
type Queue_t struct {
	elmts    []interface{}
	capacity int
	head     int
	tail     int
}

// function New initializes a new Queue data type
// It will return nil in case of a error
// or a initialized Queue object
func New(size int) *Queue_t {
	if size <= 0 {
		return nil
	} else {
		return &Queue_t{
			elmts:    make([]interface{}, size),
			capacity: size,
			head:     -1,
			tail:     -1,
		}
	}
}

// The Empty method checks if a given Queue_t is empty
func (p *Queue_t) Empty() bool {
	if p.head == -1 && p.tail == -1 {
		return true
	}
	return false
}

// The Full method checks if a given Queue_t is full
func (p *Queue_t) Full() bool {
	if ((p.head + 1) % p.capacity) == p.tail {
		return true
	}
	return false
}

// The Enqueue method adds a given v value to the queue
// and returns if the value was added successfully or not
func (p *Queue_t) Enqueue(v interface{}) bool {
	if p.Empty() {
		p.head = 0
		p.tail = 0
		p.elmts[p.head] = v
		return true
	} else if p.Full() {
		return false
	} else {
		p.head = (p.head + 1) % p.capacity
		p.elmts[p.head] = v
		return true
	}
}

// The Dequeue method removes a value from the queue
// and returns this removed value or a nil in case
// of a empty queue
func (p *Queue_t) Dequeue() (v interface{}) {
	if p.Empty() {
		v = nil
	} else {
		v = p.elmts[p.tail]
		if p.tail == p.head {
			p.tail = -1
			p.head = -1
		} else {
			p.tail = (p.tail + 1) % p.capacity
		}
	}
	return
}

// The Peek method returns the next value
// to be removed from the queue or nil
func (p *Queue_t) Peek() interface{} {
	if p.Empty() {
		return nil
	} else {
		return p.elmts[p.tail]
	}
}

// The Remaining method returns how many elements are
// left empty to be used in the Queue
func (p *Queue_t) Remaining() (cnt int) {
	if p.Empty() {
		cnt = p.capacity
	} else if p.Full() {
		cnt = 0
	} else {
		for inter := p.head; ((inter + 1) % p.capacity) != p.tail; inter++ {
			cnt++
		}
	}
	return
}
