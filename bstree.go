package bstree

import (
	"sync"
)

type BSTree struct {
	m sync.RWMutex

	isSafe bool

	root *node
}

func New(opts ...func(*BSTree)) *BSTree {
	tree := &BSTree{}

	for _, opt := range opts {
		opt(tree)
	}

	return tree
}

func (tree *BSTree) Insert(key string, data interface{}) {
	if tree.isSafe {
		tree.m.Lock()
		defer tree.m.Unlock()
	}

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
	if tree.isSafe {
		tree.m.RLock()
		defer tree.m.RUnlock()
	}

	return tree.root.get(key)
}

func (tree *BSTree) ToSlice() []interface{} {
	var datum []interface{}

	if tree.isSafe {
		tree.m.RLock()
		defer tree.m.RUnlock()
	}

	tree.root.ToSlice(&datum)
	return datum
}

func (tree *BSTree) Keys() []string {
	var keys []string

	if tree.isSafe {
		tree.m.RLock()
		defer tree.m.RUnlock()
	}

	tree.root.Keys(&keys)
	return keys
}
