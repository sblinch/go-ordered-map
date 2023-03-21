package orderedmap

import (
	"golang.org/x/exp/constraints"
)

// FromMap creates a new OrderedMap from the contents of an existing Golang map.
func FromMap[K constraints.Ordered, V any](src map[K]V) *OrderedMap[K, V] {
	m := New[K, V](len(src))
	for k, v := range src {
		m.Set(k, v)
	}
	return m
}

// Each calls f() for each element in the OrderedMap in order.
func (om *OrderedMap[K, V]) Each(f func(K, V) error) error {
	for node := om.list.Front(); node != nil; node = node.Next() {
		if err := f(node.Value.Key, node.Value.Value); err != nil {
			return err
		}
	}
	return nil
}

// Map returns a copy of the OrderedMap as an unordered Golang map.
func (om *OrderedMap[K, V]) Map() map[K]V {
	r := make(map[K]V, om.Len())
	for node := om.list.Front(); node != nil; node = node.Next() {
		r[node.Value.Key] = node.Value.Value
	}
	return r
}

// Keys returns a slice containing the keys of the OrderedMap.
func (om *OrderedMap[K, V]) Keys() []K {
	r := make([]K, 0, om.list.Len())
	for node := om.list.Front(); node != nil; node = node.Next() {
		r = append(r, node.Value.Key)
	}
	return r
}

// Values returns a slice containing the values of the OrderedMap.
func (om *OrderedMap[K, V]) Values() []V {
	r := make([]V, 0, om.list.Len())
	for node := om.list.Front(); node != nil; node = node.Next() {
		r = append(r, node.Value.Value)
	}
	return r
}

// Index returns the Pair at the ith index in order.
func (om *OrderedMap[K, V]) Index(i int) (k K, v V) {
	n := 0
	for e := om.list.Front(); e != nil; e = e.Next() {
		if n == i {
			return e.Value.Key, e.Value.Value
		}
		n++
	}
	return
}

// Sort sorts the OrderedMap using the provided less function.
func (om *OrderedMap[K, V]) Sort(less func(*Pair[K, V], *Pair[K, V]) bool) {
	outer := om.list.Front()
	if outer == nil {
		return
	}

	for outer != nil {
		inner := outer.Next()
		for inner != nil {
			if less(inner.Value, outer.Value) {
				outer.Value, inner.Value = inner.Value, outer.Value
			}
			inner = inner.Next()
		}
		outer = outer.Next()
	}
}

// SortKeys sorts the OrderedMap by key using the provided less function.
func (om *OrderedMap[K, V]) SortKeys(less func(K, K) bool) {
	outer := om.list.Front()
	if outer == nil {
		return
	}

	for outer != nil {
		inner := outer.Next()
		for inner != nil {
			if less(inner.Value.Key, outer.Value.Key) {
				outer.Value, inner.Value = inner.Value, outer.Value
			}
			inner = inner.Next()
		}
		outer = outer.Next()
	}
}

// SortValues sorts the OrderedMap by value using the provided less function.
func (om *OrderedMap[K, V]) SortValues(less func(V, V) bool) {
	outer := om.list.Front()
	if outer == nil {
		return
	}

	for outer != nil {
		inner := outer.Next()
		for inner != nil {
			if less(inner.Value.Value, outer.Value.Value) {
				outer.Value, inner.Value = inner.Value, outer.Value
			}
			inner = inner.Next()
		}
		outer = outer.Next()
	}
}
