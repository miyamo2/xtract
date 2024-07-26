package xtract_test

import (
	"fmt"
	"github.com/miyamo2/xtract"
	"strings"
)

func Example() {
	s := []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 100, 101}
	even := xtract.FromSlice(s).ByValue(func(i int) bool { return i%2 == 0 })
	odd := xtract.FromSlice(s).ByValue(func(i int) bool { return i%2 != 0 })

	fmt.Println("---even---")
	for v := range even.Values() {
		fmt.Println(v)
	}

	fmt.Println("---odd---")
	for v := range odd.Values() {
		fmt.Println(v)
	}

	evenAndTwoDigits := even.ByValue(func(i int) bool { return i > 9 && i < 100 })
	fmt.Println("---even and two digits---")
	for v := range evenAndTwoDigits.Values() {
		fmt.Println(v)
	}

	oddAndTwoDigits := odd.ByValue(func(i int) bool { return i > 9 && i < 100 })
	fmt.Println("---odd and two digits---")
	for v := range oddAndTwoDigits.Values() {
		fmt.Println(v)
	}
	// Output: ---even---
	//0
	//2
	//4
	//6
	//8
	//10
	//100
	//---odd---
	//1
	//3
	//5
	//7
	//9
	//11
	//101
	//---even and two digits---
	//10
	//---odd and two digits---
	//11
}

func ExampleSliceExtractor_ByValue() {
	s := []string{"gopher", "iterator", "range over func"}
	xt := xtract.FromSlice(s).ByValue(func(v string) bool { return len(v) < 9 })
	for v := range xt.Values() {
		fmt.Println(v)
	}
	// Output: gopher
	//iterator
}

func ExampleSliceExtractor_ByKey() {
	s := []string{"gopher", "iterator", "range over func"}
	xt := xtract.FromSlice(s).ByKey(func(i int) bool { return i > 0 })
	for v := range xt.Values() {
		fmt.Println(v)
	}
	// Output: iterator
	//range over func
}

func ExampleSliceExtractor_ByKeyAndValue() {
	s := []string{"gopher", "iterator", "range over func"}
	xt := xtract.FromSlice(s).ByKeyAndValue(func(i int, v string) bool { return i > 1 && len(v) > 6 })
	for v := range xt.Values() {
		fmt.Println(v)
	}
	// Output: range over func
}

func ExampleSliceExtractor_Limit() {
	s := []string{"gopher", "iterator", "range over func"}
	xt := xtract.FromSlice(s).Limit(1)
	for v := range xt.Values() {
		fmt.Println(v)
	}
	// Output: gopher
}

func ExampleSliceExtractor_Offset() {
	s := []string{"gopher", "iterator", "range over func"}
	xt := xtract.FromSlice(s).Offset(1)
	for v := range xt.Values() {
		fmt.Println(v)
	}
	// Output: iterator
	//range over func
}

func ExampleMapExtractor_ByValue() {
	m := map[string]string{"language": "gopher", "design pattern": "iterator", "implementation": "range over func"}
	xt := xtract.FromMap(m).ByValue(func(v string) bool { return len(v) < 8 })
	for v := range xt.Values() {
		fmt.Println(v)
	}
	// Output: gopher
}

func ExampleMapExtractor_ByKey() {
	m := map[string]string{"language": "gopher", "design pattern": "iterator", "implementation": "range over func"}
	xt := xtract.FromMap(m).ByKey(func(k string) bool { return strings.Contains(k, " ") })
	for v := range xt.Values() {
		fmt.Println(v)
	}
	// Output: iterator
}

func ExampleMapExtractor_ByKeyAndValue() {
	m := map[string]string{"language": "gopher", "design pattern": "iterator", "implementation": "range over func"}
	xt := xtract.FromMap(m).ByKeyAndValue(func(k, v string) bool { return strings.Contains(k, "e") && len(v) < 8 })
	for v := range xt.Values() {
		fmt.Println(v)
	}
	// Output: gopher
}

func ExampleMapExtractor_Limit() {
	m := map[string]string{"language": "gopher", "design pattern": "iterator", "implementation": "range over func"}
	xt := xtract.FromMap(m).Limit(1)

	values := make([]string, 0, 1)
	for v := range xt.Values() {
		values = append(values, v)
	}
	fmt.Println(len(values))
	// Output: 1
}

func ExampleMapExtractor_Offset() {
	m := map[string]string{"language": "gopher", "design pattern": "iterator", "implementation": "range over func"}
	xt := xtract.FromMap(m).Offset(1)

	values := make([]string, 0, 1)
	for v := range xt.Values() {
		values = append(values, v)
	}
	fmt.Println(len(values))
	// Output: 2
}
