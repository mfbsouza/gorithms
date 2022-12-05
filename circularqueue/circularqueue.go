// Package circularqueue provides a FIFO algorithm implementation
package circularqueue

// Queue interface represents the common operations of a FIFO object
type Queue interface {
	IsEmpty() bool
	IsFull() bool
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

// function New initializes a new Queue object.
// This method returns nil in case of a error or a initialized Queue object
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

// IsEmpty method checks if a given Queue object is empty
func (p *Queue_t) IsEmpty() bool {
	if p.head == -1 && p.tail == -1 {
		return true
	}
	return false
}

// IsFull method checks if a given Queue object is full
func (p *Queue_t) IsFull() bool {
	if ((p.head + 1) % p.capacity) == p.tail {
		return true
	}
	return false
}

// Enqueue method adds a given v value to the Queue object.
// This method returns if the value was successfully added or not
func (p *Queue_t) Enqueue(v interface{}) bool {
	if p.IsEmpty() {
		p.head = 0
		p.tail = 0
		p.elmts[p.head] = v
		return true
	} else if p.IsFull() {
		return false
	} else {
		p.head = (p.head + 1) % p.capacity
		p.elmts[p.head] = v
		return true
	}
}

// Dequeue method removes a value from the Queue object.
// This method returns nil in case of a empty Queue or a value
func (p *Queue_t) Dequeue() (v interface{}) {
	if p.IsEmpty() {
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

// Peek method returns the next value to be dequeued.
// This method returns nil in canse of a empty Queue or a value
func (p *Queue_t) Peek() interface{} {
	if p.IsEmpty() {
		return nil
	} else {
		return p.elmts[p.tail]
	}
}

// Remaining method returns how many free slots are left in the Queue object
func (p *Queue_t) Remaining() (cnt int) {
	if p.IsEmpty() {
		cnt = p.capacity
	} else {
		for inter := p.head; ((inter + 1) % p.capacity) != p.tail; inter++ {
			cnt++
		}
	}
	return
}
