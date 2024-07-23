package xtract

import "iter"

// compatibility checks
var (
	_ Extractor[int, string] = (*SliceExtractor[string])(nil)
	_ Extractor[string, int] = (*MapExtractor[string, int])(nil)
)

// Extractor provides features to extract values from a collection.
type Extractor[K comparable, V any] interface {
	// ByValue filters the values of the collection by their values.
	ByValue(condition func(V) bool) Extractor[K, V]
	// ByKey filters the values of the collection by their keys.
	ByKey(condition func(K) bool) Extractor[K, V]
	// ByKeyAndValue filters the values of the collection by their keys and values.
	ByKeyAndValue(condition func(K, V) bool) Extractor[K, V]
	// Values returns a sequence of values.
	Values() iter.Seq[V]
}

// SliceExtractor is implementation of Extractor for slice.
type SliceExtractor[V any] struct {
	seq iter.Seq2[int, V]
}

// ByValue See: Extractor.ByValue
func (x *SliceExtractor[V]) ByValue(condition func(V) bool) Extractor[int, V] {
	return &SliceExtractor[V]{
		seq: byValue(x.seq, condition),
	}
}

// ByKey See: Extractor.ByKey
func (x *SliceExtractor[V]) ByKey(condition func(int) bool) Extractor[int, V] {
	return &SliceExtractor[V]{
		seq: byKey(x.seq, condition),
	}
}

// ByKeyAndValue See: Extractor.ByKeyAndValue
func (x *SliceExtractor[V]) ByKeyAndValue(condition func(int, V) bool) Extractor[int, V] {
	return &SliceExtractor[V]{
		seq: byKeyAndValue(x.seq, condition),
	}
}

// Values See: Extractor.Values
func (x *SliceExtractor[V]) Values() iter.Seq[V] {
	return values(x.seq)
}

// FromSlice returns Extractor for a slice.
func FromSlice[V any](s []V) Extractor[int, V] {
	return &SliceExtractor[V]{
		seq: func(yield func(int, V) bool) {
			for i, v := range s {
				if !yield(i, v) {
					return
				}
			}
		}}
}

// MapExtractor is implementation of Extractor for map.
type MapExtractor[K comparable, V any] struct {
	seq iter.Seq2[K, V]
}

// ByValue See: Extractor.ByValue
func (x MapExtractor[K, V]) ByValue(condition func(V) bool) Extractor[K, V] {
	return &MapExtractor[K, V]{
		seq: byValue(x.seq, condition),
	}
}

// ByKey See: Extractor.ByKey
func (x MapExtractor[K, V]) ByKey(condition func(K) bool) Extractor[K, V] {
	return &MapExtractor[K, V]{
		seq: byKey(x.seq, condition),
	}
}

// ByKeyAndValue See: Extractor.ByKeyAndValue
func (x MapExtractor[K, V]) ByKeyAndValue(condition func(K, V) bool) Extractor[K, V] {
	return &MapExtractor[K, V]{
		seq: byKeyAndValue(x.seq, condition),
	}
}

// Values See: Extractor.Values
func (x MapExtractor[K, V]) Values() iter.Seq[V] {
	return values(x.seq)
}

// FromMap returns Extractor for a map.
func FromMap[K comparable, V any](m map[K]V) Extractor[K, V] {
	return &MapExtractor[K, V]{
		seq: func(yield func(K, V) bool) {
			for k, v := range m {
				if !yield(k, v) {
					return
				}
			}
		},
	}
}

func byValue[K comparable, V any](seq iter.Seq2[K, V], condition func(V) bool) iter.Seq2[K, V] {
	return func(yield func(K, V) bool) {
		seq(func(k K, v V) bool {
			if condition(v) && !yield(k, v) {
				return false
			}
			return true
		})
	}
}

func byKey[K comparable, V any](seq iter.Seq2[K, V], condition func(K) bool) iter.Seq2[K, V] {
	return func(yield func(K, V) bool) {
		seq(func(k K, v V) bool {
			if condition(k) && !yield(k, v) {
				return false
			}
			return true
		})
	}
}

func byKeyAndValue[K comparable, V any](seq iter.Seq2[K, V], condition func(K, V) bool) iter.Seq2[K, V] {
	return func(yield func(K, V) bool) {
		seq(func(k K, v V) bool {
			if condition(k, v) && !yield(k, v) {
				return false
			}
			return true
		})
	}
}

func values[K comparable, V any](seq iter.Seq2[K, V]) iter.Seq[V] {
	return func(yield func(V) bool) {
		seq(func(_ K, v V) bool {
			if !yield(v) {
				return false
			}
			return true
		})
	}
}
