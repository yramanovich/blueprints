package blueprints

type node[E comparable] struct {
	item E
	prev *node[E]
	next *node[E]
}

type linkedList[E comparable] struct {
	first *node[E]
	last  *node[E]
	size  int
}

func (ll *linkedList[E]) add(e E) {
	ll.linkLast(e)
}

func (ll *linkedList[E]) linkLast(e E) {
	l := ll.last
	newNode := &node[E]{item: e, prev: l, next: nil}
	ll.last = newNode
	if l == nil {
		ll.first = newNode
	} else {
		l.next = newNode
	}
	ll.size++
}

func (ll *linkedList[E]) len() int {
	return ll.size
}

func (ll *linkedList[E]) remove(e E) {
	for n := ll.first; n != nil; n = n.next {
		if n.item == e {
			ll.unlink(n)
		}
	}
}

func (ll *linkedList[E]) unlink(n *node[E]) {
	prev, next := n.prev, n.next
	if prev == nil {
		ll.first = next
	} else {
		prev.next = next
		n.prev = nil
	}

	if next == nil {
		ll.last = prev
	} else {
		next.prev = prev
		n.next = nil
	}
	//var zero E
	//n.item = zero
	ll.size--
}
