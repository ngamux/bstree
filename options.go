package bstree

func WithSafe() func(*BSTree) {
	return func(b *BSTree) {
		b.isSafe = true
	}
}
