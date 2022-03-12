package bstree

type node struct {
	key  string
	data interface{}

	left  *node
	right *node
}

func (n *node) insert(newn *node) {
	switch {
	case n.key == newn.key:
		*n = *newn
	case n.key > newn.key:
		if n.left == nil {
			n.left = newn
		} else {
			n.left.insert(newn)
		}
	case n.key < newn.key:
		if n.right == nil {
			n.right = newn
		} else {
			n.right.insert(newn)
		}
	}
}

func (n *node) get(key string) interface{} {
	if n == nil {
		return nil
	}
	if n.key == key {
		return n.data
	}
	if n.key > key {
		return n.left.get(key)
	}
	if n.key < key {
		return n.right.get(key)
	}

	return nil
}

func (n *node) ToSlice(datum *[]interface{}) {
	if n == nil {
		return
	}

	n.left.ToSlice(datum)
	*datum = append(*datum, n.data)
	n.right.ToSlice(datum)
}
