package xtract

// compatibility checks
var (
	_ Extractor[int, string] = (*SliceExtractor[string])(nil)
	_ Extractor[string, int] = (*MapExtractor[string, int])(nil)
)

// Extractor provides features to extract values from a collection.
type Extractor[K comparable, V any] interface {
	// ByValue filters the values of the collection by their values.
	ByValue(condition func(V) bool) func(yield func(V) bool)
	// ByKey filters the values of the collection by their keys.
	ByKey(condition func(K) bool) func(yield func(V) bool)
	// ByKeyAndValue filters the values of the collection by their keys and values.
	ByKeyAndValue(condition func(K, V) bool) func(yield func(V) bool)
}

// SliceExtractor is implementation of Extractor for slice.
type SliceExtractor[T any] struct {
	s []T
}

// ByValue See: Extractor.ByValue
func (x *SliceExtractor[T]) ByValue(condition func(T) bool) func(yield func(T) bool) {
	return func(yield func(T) bool) {
		for _, v := range x.s {
			if condition(v) && !yield(v) {
				return
			}
		}
	}
}

// ByKey See: Extractor.ByKey
func (x *SliceExtractor[T]) ByKey(condition func(int) bool) func(yield func(T) bool) {
	return func(yield func(T) bool) {
		for i, v := range x.s {
			if condition(i) && !yield(v) {
				return
			}
		}
	}
}

// ByKeyAndValue See: Extractor.ByKeyAndValue
func (x *SliceExtractor[T]) ByKeyAndValue(condition func(int, T) bool) func(yield func(T) bool) {
	return func(yield func(T) bool) {
		for i, v := range x.s {
			if condition(i, v) && !yield(v) {
				return
			}
		}
	}
}

// FromSlice returns Extractor for a slice.
func FromSlice[T any](s []T) Extractor[int, T] {
	cp := make([]T, 0, len(s))
	cp = append(cp, s...)
	return &SliceExtractor[T]{s: cp}
}

// MapExtractor is implementation of Extractor for map.
type MapExtractor[T comparable, U any] struct {
	m map[T]U
}

// ByValue See: Extractor.ByValue
func (x MapExtractor[K, V]) ByValue(condition func(V) bool) func(yield func(V) bool) {
	return func(yield func(V) bool) {
		for _, v := range x.m {
			if condition(v) && !yield(v) {
				return
			}
		}
	}
}

// ByKey See: Extractor.ByKey
func (x MapExtractor[K, V]) ByKey(condition func(K) bool) func(yield func(V) bool) {
	return func(yield func(V) bool) {
		for i, v := range x.m {
			if condition(i) && !yield(v) {
				return
			}
		}
	}
}

// ByKeyAndValue See: Extractor.ByKeyAndValue
func (x MapExtractor[K, V]) ByKeyAndValue(condition func(K, V) bool) func(yield func(V) bool) {
	return func(yield func(V) bool) {
		for k, v := range x.m {
			if condition(k, v) && !yield(v) {
				return
			}
		}
	}
}

// FromMap returns Extractor for a map.
func FromMap[K comparable, V any](m map[K]V) Extractor[K, V] {
	cp := make(map[K]V)
	for k, v := range m {
		cp[k] = v
	}
	return &MapExtractor[K, V]{m: cp}
}
