package bstree

import (
	"testing"

	"github.com/golang-must/must"
)

func TestBSTree(t *testing.T) {

	tree := &BSTree{}

	t.Run("New returns new BSTree", func(t *testing.T) {
		must := must.New(t)

		expected := &BSTree{}
		result := New()

		must.Equal(expected, result)
	})

	t.Run("New with Opts returns new BSTree", func(t *testing.T) {
		must := must.New(t)

		expected := &BSTree{
			isSafe: true,
		}
		result := New(WithSafe())

		must.Equal(expected, result)
	})

	t.Run("can insert and get inserted data", func(t *testing.T) {
		must := must.New(t)

		expected := 1
		tree.Insert("e", expected)
		result := tree.Get("e")

		must.Equal(expected, result)
	})

	t.Run("can insert and get left node", func(t *testing.T) {
		must := must.New(t)

		expected := 1
		tree.Insert("d", expected)
		result := tree.Get("d")

		must.Equal(expected, result)
	})

	t.Run("can insert and get right node", func(t *testing.T) {
		must := must.New(t)

		expected := 1
		tree.Insert("f", expected)
		result := tree.Get("f")

		must.Equal(expected, result)
	})

	t.Run("can re-insert and get data", func(t *testing.T) {
		must := must.New(t)

		expected := 2
		tree.Insert("e", expected)
		result := tree.Get("e")

		must.Equal(expected, result)
	})

	t.Run("can re-insert left and get data", func(t *testing.T) {
		must := must.New(t)

		expected := 2
		tree.Insert("c", expected)
		tree.Insert("b", expected)
		result := tree.Get("b")

		must.Equal(expected, result)
	})

	t.Run("can re-insert right and get data", func(t *testing.T) {
		must := must.New(t)

		expected := 2
		tree.Insert("g", expected)
		tree.Insert("h", expected)
		result := tree.Get("h")

		must.Equal(expected, result)
	})

	t.Run("unknown key returns nil", func(t *testing.T) {
		must := must.New(t)

		result := tree.Get("zzz")
		must.Nil(result)
	})

	t.Run("can convert to slice", func(t *testing.T) {
		must := must.New(t)

		result := tree.ToSlice()
		must.True(len(result) > 0)
	})

}

func BenchmarkBSTree(b *testing.B) {

	b.Run("insert with static key", func(b *testing.B) {
		tree := &BSTree{}
		for i := 0; i < b.N; i++ {
			tree.Insert("a", i)
		}
	})

	b.Run("get with static key", func(b *testing.B) {
		tree := &BSTree{}
		tree.Insert("a", 1)
		for i := 0; i < b.N; i++ {
			tree.Get("a")
		}
	})

}
