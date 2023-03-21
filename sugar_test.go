package orderedmap

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestOrderedMap_Keys(t *testing.T) {
	om := New[string, int](WithInitialData[string, int](
		Pair[string, int]{Key: "foo", Value: 97},
		Pair[string, int]{Key: "bar", Value: 3},
		Pair[string, int]{Key: "baz", Value: 249},
	))
	assert.Equalf(t, []string{"foo", "bar", "baz"}, om.Keys(), "Keys()")
}

func TestOrderedMap_Values(t *testing.T) {
	om := New[string, int](WithInitialData[string, int](
		Pair[string, int]{Key: "foo", Value: 97},
		Pair[string, int]{Key: "bar", Value: 3},
		Pair[string, int]{Key: "baz", Value: 249},
	))
	assert.Equalf(t, []int{97, 3, 249}, om.Values(), "Keys()")
}

func TestOrderedMap_Index(t *testing.T) {
	om := New[string, int](WithInitialData[string, int](
		Pair[string, int]{Key: "foo", Value: 97},
		Pair[string, int]{Key: "bar", Value: 3},
		Pair[string, int]{Key: "baz", Value: 249},
	))
	k, v := om.Index(1)
	assert.Equalf(t, "bar", k, "Index() key")
	assert.Equalf(t, 3, v, "Index() value")
}

func TestOrderedMap_Map(t *testing.T) {
	om := New[string, int](WithInitialData[string, int](
		Pair[string, int]{Key: "foo", Value: 97},
		Pair[string, int]{Key: "bar", Value: 3},
		Pair[string, int]{Key: "baz", Value: 249},
	))
	m := om.Map()
	for k, v := range m {
		assert.Equalf(t, v, om.Value(k), "Map(%s)", k)
	}
}

func TestOrderedMap_Each(t *testing.T) {
	om := New[string, int](WithInitialData[string, int](
		Pair[string, int]{Key: "foo", Value: 97},
		Pair[string, int]{Key: "bar", Value: 3},
		Pair[string, int]{Key: "baz", Value: 249},
	))
	ks := make([]string, 0, om.Len())
	vs := make([]int, 0, om.Len())
	err := om.Each(func(k string, v int) error {
		ks = append(ks, k)
		vs = append(vs, v)
		return nil
	})
	assert.NoError(t, err)
	assert.Equalf(t, []string{"foo", "bar", "baz"}, ks, "Each() keys")
	assert.Equalf(t, []int{97, 3, 249}, vs, "Each() keys")
}

func TestOrderedMap_SortKeys(t *testing.T) {
	om := FromMap(map[string]int{"foo": 97, "bar": 3, "baz": 249})
	om.SortKeys(func(s string, s2 string) bool {
		return s < s2
	})
	assert.Equalf(t, []string{"bar", "baz", "foo"}, om.Keys(), "SortKeys() keys")
	assert.Equalf(t, []int{3, 249, 97}, om.Values(), "SortKeys() values")
}

func TestOrderedMap_SortValues(t *testing.T) {
	om := FromMap(map[string]int{"foo": 97, "bar": 3, "baz": 249})
	om.SortValues(func(i, i2 int) bool {
		return i < i2
	})
	assert.Equalf(t, []string{"bar", "foo", "baz"}, om.Keys(), "SortValues() keys")
	assert.Equalf(t, []int{3, 97, 249}, om.Values(), "SortValues() values")
}

func TestOrderedMap_Sort(t *testing.T) {
	om := FromMap(map[string]int{"foo": 97, "bar": 3, "baz": 249})
	om.Sort(func(p *Pair[string, int], p2 *Pair[string, int]) bool {
		return p.Key < p2.Key
	})
	assert.Equalf(t, []string{"bar", "baz", "foo"}, om.Keys(), "Sort() keys")
	assert.Equalf(t, []int{3, 249, 97}, om.Values(), "Sort() values")
}
