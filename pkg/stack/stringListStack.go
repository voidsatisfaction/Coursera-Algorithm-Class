package stack

type StringListStack struct {
	Head *element
	Size int
}

type element struct {
	item string
	next *element
}

func (ss *StringListStack) Push(item string) {
	temp := ss.Head
	newElement := &element{
		item: item,
		next: temp,
	}
	ss.Head = newElement
	ss.Size++
}

func (ss *StringListStack) Pop() string {
	if ss.Head != nil {
		head := ss.Head
		ss.Head = ss.Head.next
		ss.Size--
		return head.item
	}
	return ""
}

func (ss *StringListStack) IsEmpty() bool {
	return ss.Head == nil
}
