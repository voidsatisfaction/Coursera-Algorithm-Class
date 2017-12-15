package queue

type StringListQueue struct {
	Head *element
	Tail *element
	Size int
}

type element struct {
	item string
	next *element
}

func (q *StringListQueue) Enqueue(item string) {
	e := &element{
		item: item,
		next: nil,
	}
	if q.IsEmpty() {
		q.Head = e
		q.Tail = e
	} else {
		q.Tail.next = e
	}
	q.Size++
}

func (q *StringListQueue) Dequeue() string {
	temp := q.Head.item
	if q.IsEmpty() {
		return ""
	} else if q.Size == 1 {
		q.Head = nil
		q.Tail = nil
	} else {
		q.Head = q.Head.next
	}
	q.Size--
	return temp
}

func (q *StringListQueue) IsEmpty() bool {
	return q.Head == nil
}
