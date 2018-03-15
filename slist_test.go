package slist

import (
	"math/rand"
	"testing"
)

type testItem struct {
	key   int
	value int
}

func (k testItem) Cmp(other NodeItem) bool {
	return k.key > other.(testItem).key
}

func (k testItem) Eq(other NodeItem) bool {
	return k.key == other.(testItem).key
}

type descTestItem struct {
	key   int
	value int
}

func (k descTestItem) Cmp(other NodeItem) bool {
	return k.key < other.(descTestItem).key
}

func (k descTestItem) Eq(other NodeItem) bool {
	return k.key == other.(descTestItem).key
}

func TestSortsOnInsert(t *testing.T) {
	list := NewSList()
	list.Insert(testItem{5, 50})
	list.Insert(testItem{7, 70})
	list.Insert(testItem{3, 30})
	first := list.Head()

	firstValue := first.item.(testItem).value
	if firstValue != 30 {
		t.Errorf("Expected first key to be 30 was %v", firstValue)
	}

	secondValue := first.Next().item.(testItem).value
	if secondValue != 50 {
		t.Errorf("Expected second key to be 50 was %v", secondValue)
	}

	thirdValue := first.Next().Next().item.(testItem).value
	if thirdValue != 70 {
		t.Errorf("Expected third key to be 70 was %v", thirdValue)
	}
}

func TestSortsDescOnInsert(t *testing.T) {
	list := NewSList()
	list.Insert(descTestItem{5, 50})
	list.Insert(descTestItem{7, 70})
	list.Insert(descTestItem{3, 30})
	first := list.Head()

	firstValue := first.item.(descTestItem).value
	if firstValue != 70 {
		t.Errorf("Expected first key to be 70 was %v", firstValue)
	}

	secondValue := first.Next().item.(descTestItem).value
	if secondValue != 50 {
		t.Errorf("Expected second key to be 50 was %v", secondValue)
	}

	thirdValue := first.Next().Next().item.(descTestItem).value
	if thirdValue != 30 {
		t.Errorf("Expected third key to be 30 was %v", thirdValue)
	}
}

func TestFindsAnItem(t *testing.T) {
	list := NewSList()
	list.Insert(testItem{5, 50})
	list.Insert(testItem{7, 70})
	list.Insert(testItem{3, 30})

	item := list.Find(testItem{key: 5})
	value := item.(testItem).value
	if value != 50 {
		t.Errorf("Expected second key to be 50 was %v", value)
	}
}

func TestCustomOnEqual(t *testing.T) {
	replacer := func(node *Node, item NodeItem) {
		node.item = item
	}
	list := NewSListWithOnEqualHandler(replacer)
	list.Insert(testItem{5, 50})
	list.Insert(testItem{7, 70})
	list.Insert(testItem{3, 30})
	list.Insert(testItem{5, 100})

	item := list.Find(testItem{key: 5})
	value := item.(testItem).value
	if value != 100 {
		t.Errorf("Expected item to be replaced with 100, was %v", value)
	}
}

func TestFindReturnsNilIfNotFound(t *testing.T) {
	list := NewSList()
	item := list.Find(testItem{key: 42})
	if item != nil {
		t.Errorf("Expected item to be nil")
	}

	list.Insert(testItem{key: 1, value: 42})
	item = list.Find(testItem{key: 42})
	if item != nil {
		t.Errorf("Expected item to be nil")
	}
}

func TestRemove(t *testing.T) {
	list := NewSList()
	_, err := list.Remove(testItem{key: 42})
	if err == nil {
		t.Errorf("Removing on empty list should err")
	}

	list.Insert(testItem{key: 5})
	list.Insert(testItem{key: 7})
	list.Insert(testItem{key: 3})
	removedItem, err := list.Remove(testItem{key: 5})
	if removedItem.(testItem).key != 5 || err != nil {
		t.Errorf("Removing an existing item should return the item and nil")
	}
}

func benchmarkInsertSequential(count int, b *testing.B) {
	list := NewSList()
	for n := 0; n < b.N; n++ {
		for i := 0; i < count; i++ {
			list.Insert(testItem{key: i})
		}
	}
}

func BenchmarkInsertSequential10(b *testing.B)     { benchmarkInsertSequential(10, b) }
func BenchmarkInsertSequential100(b *testing.B)    { benchmarkInsertSequential(100, b) }
func BenchmarkInsertSequential1000(b *testing.B)   { benchmarkInsertSequential(1000, b) }
func BenchmarkInsertSequential100000(b *testing.B) { benchmarkInsertSequential(100000, b) }

func benchmarkInsertRandom(count int, b *testing.B) {
	list := NewSList()
	for n := 0; n < b.N; n++ {
		for i := 0; i < count; i++ {
			list.Insert(testItem{key: rand.Intn(count)})
		}
	}
}

func BenchmarkInsertRandom10(b *testing.B)     { benchmarkInsertRandom(10, b) }
func BenchmarkInsertRandom100(b *testing.B)    { benchmarkInsertRandom(100, b) }
func BenchmarkInsertRandom1000(b *testing.B)   { benchmarkInsertRandom(1000, b) }
func BenchmarkInsertRandom100000(b *testing.B) { benchmarkInsertRandom(100000, b) }
