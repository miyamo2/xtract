# xtract

[![Go Reference](https://pkg.go.dev/badge/github.com/miyamo2/xtract.svg)](https://pkg.go.dev/github.com/miyamo2/xtract)
[![GitHub go.mod Go version](https://img.shields.io/github/go-mod/go-version/miyamo2/xtract)](https://img.shields.io/github/go-mod/go-version/miyamo2/xtract)
[![GitHub release (latest by date)](https://img.shields.io/github/v/release/miyamo2/xtract)](https://img.shields.io/github/v/release/miyamo2/xtract)
[![codecov](https://codecov.io/gh/miyamo2/xtract/graph/badge.svg?token=PXU3HXGBWQ)](https://codecov.io/gh/miyamo2/xtract)
[![Go Report Card](https://goreportcard.com/badge/github.com/miyamo2/xtract)](https://goreportcard.com/report/github.com/miyamo2/xtract)
[![GitHub License](https://img.shields.io/github/license/miyamo2/xtract?&color=blue)](https://img.shields.io/github/license/miyamo2/xtract?&color=blue)

Extract from collection and build iterator.

## Quick Start

### Install

```sh
go get github.com/miyamo2/xtract
```

### Setup `GOEXPERIMENT`

> [!IMPORTANT]
> 
> If your Go project is Go 1.23 or higher, this section is not necessary.
> 
> Also, if Go 1.21 or lower, you will need to update to Go 1.22.

```sh
go env -w GOEXPERIMENT=rangefunc
```

### Usage

```go
s := []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 100, 101}
xt := xtract.FromSlice(s)
even := xt.ByValue(func(i int) bool { return i%2 == 0 })
odd := xt.ByValue(func(i int) bool { return i%2 != 0 })

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
```

#### With `SliceExtractor.ByValue`

```go
s := []string{"go", "iterator", "range over func"}
xt := xtract.FromSlice(s).ByValue(func(v string) bool { return len(v) < 9 })
for v := range xt.Values() {
    fmt.Println(v)
}
// Output: go
//iterator
```

#### With `SliceExtractor.ByKey`

```go
s := []string{"go", "iterator", "range over func"}
xt := xtract.FromSlice(s).ByKey(func(i int) bool { return i > 0 })
for v := range xt.Values() {
    fmt.Println(v)
}
// Output: go
//iterator
```

#### With `SliceExtractor.ByKeyAndValue`

```go
s := []string{"go", "iterator", "range over func"}
xt := xtract.FromSlice(s).ByKeyAndValue(func(i int, v string) bool { return i > 1 && len(v) > 6 })
for v := range xt.Values()
    fmt.Println(v)
}
// Output: range over func
```

#### With `MapExtractor.ByValue`

```go
m := map[string]string{"language": "go", "design pattern": "iterator", "implementation": "range over func"}
xt := xtract.FromMap(m).ByValue(func(v string) bool { return len(v) < 8 })
for v := range xt.Values() {
    fmt.Println(v)
}
// Output: go
```

#### With `MapExtractor.ByKey`

```go
m := map[string]string{"language": "go", "design pattern": "iterator", "implementation": "range over func"}
xt := xtract.FromMap(m).ByKey(func(k string) bool { return strings.Contains(k, " ") })
for v := range xt.Values() {
    fmt.Println(v)
}
// Output: iterator
```

#### With `MapExtractor.ByKeyAndValue`

```go
m := map[string]string{"language": "go", "design pattern": "iterator", "implementation": "range over func"}
xt := xtract.FromMap(m).ByKeyAndValue(func(k, v string) bool { return strings.Contains(k, "e") && len(v) < 8 })
for v := range xt.Values() {
    fmt.Println(v)
}
// Output: go
```

## Contributing

Feel free to open a PR or an Issue.

## License

**xtract** released under the [MIT License](https://github.com/miyamo2/xtract/blob/main/LICENSE)
