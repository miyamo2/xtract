package xtract_test

import (
	"fmt"
	"github.com/miyamo2/xtract"
	"strings"
)

func ExampleSliceExtractor_ByValue() {
	s := []string{"gopher", "iterator", "range over func"}
	for v := range xtract.FromSlice(s).ByValue(func(v string) bool { return len(v) < 9 }) {
		fmt.Println(v)
	}
	// Output: gopher
	//iterator
}

func ExampleSliceExtractor_ByKey() {
	s := []string{"gopher", "iterator", "range over func"}
	for v := range xtract.FromSlice(s).ByKey(func(i int) bool { return i < 2 }) {
		fmt.Println(v)
	}
	// Output: gopher
	//iterator
}

func ExampleSliceExtractor_ByKeyAndValue() {
	s := []string{"gopher", "iterator", "range over func"}
	for v := range xtract.FromSlice(s).ByKeyAndValue(func(i int, v string) bool { return i > 1 && len(v) > 6 }) {
		fmt.Println(v)
	}
	// Output: range over func
}

func ExampleMapExtractor_ByValue() {
	m := map[string]string{"language": "gopher", "design pattern": "iterator", "implementation": "range over func"}
	for v := range xtract.FromMap(m).ByValue(func(v string) bool { return len(v) < 8 }) {
		fmt.Println(v)
	}
	// Output: gopher
}

func ExampleMapExtractor_ByKey() {
	m := map[string]string{"language": "gopher", "design pattern": "iterator", "implementation": "range over func"}
	for v := range xtract.FromMap(m).ByKey(func(k string) bool { return strings.Contains(k, " ") }) {
		fmt.Println(v)
	}
	// Output: iterator
}

func ExampleMapExtractor_ByKeyAndValue() {
	m := map[string]string{"language": "gopher", "design pattern": "iterator", "implementation": "range over func"}
	for v := range xtract.FromMap(m).ByKeyAndValue(func(k, v string) bool { return strings.Contains(k, "e") && len(v) < 8 }) {
		fmt.Println(v)
	}
	// Output: gopher
}
