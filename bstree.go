package bstree

import "sync"

type BSTree struct {
	m    sync.RWMutex
	root *node
}

func (tree *BSTree) Insert(key string, data interface{}) {
	tree.m.Lock()
	defer tree.m.Unlock()

	newn := &node{
		key:  key,
		data: data,
	}

	if tree.root == nil {
		tree.root = newn
	} else {
		tree.root.insert(newn)
	}
}

func (tree *BSTree) Get(key string) interface{} {
	tree.m.RLock()
	defer tree.m.RUnlock()

	return tree.root.get(key)
}

func (tree *BSTree) ToSlice() []interface{} {
	var datum []interface{}

	tree.m.RLock()
	defer tree.m.RUnlock()

	tree.root.ToSlice(&datum)
	return datum
}
