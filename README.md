# xtract

[![Go Reference](https://pkg.go.dev/badge/github.com/miyamo2/xtract.svg)](https://pkg.go.dev/github.com/miyamo2/xtract)
[![GitHub go.mod Go version](https://img.shields.io/github/go-mod/go-version/miyamo2/xtract)](https://img.shields.io/github/go-mod/go-version/miyamo2/xtract)
[![GitHub release (latest by date)](https://img.shields.io/github/v/release/miyamo2/xtract)](https://img.shields.io/github/v/release/miyamo2/xtract)
[![codecov](https://codecov.io/gh/miyamo2/xtract/graph/badge.svg?token=PXU3HXGBWQ)](https://codecov.io/gh/miyamo2/xtract)
[![Go Report Card](https://goreportcard.com/badge/github.com/miyamo2/xtract)](https://goreportcard.com/report/github.com/miyamo2/xtract)
[![GitHub License](https://img.shields.io/github/license/miyamo2/xtract?&color=blue)](https://img.shields.io/github/license/miyamo2/xtract?&color=blue)

Extract from collection and build iterators.

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
> Also, if Go1.21 or lower, you will need to update to Go1.22.

```sh
go env -w GOEXPERIMENT=rangefunc
```

### Usage

#### With `SliceExtractor.ByValue`

```go
s := []string{"gopher", "iterator", "range over func"}
for v := range xtract.FromSlice(s).ByValue(func(v string) bool { return len(v) < 9 }) {
    fmt.Println(v)
}
// Output: gopher
//iterator
```

#### With `SliceExtractor.ByKey`

```go
s := []string{"gopher", "iterator", "range over func"}
for v := range xtract.FromSlice(s).ByKey(func(i int) bool { return i < 2 }) {
    fmt.Println(v)
}
// Output: gopher
//iterator
```

#### With `SliceExtractor.ByKeyAndValue`

```go
s := []string{"gopher", "iterator", "range over func"}
for v := range xtract.FromSlice(s).ByKeyAndValue(func(i int, v string) bool { return i > 1 && len(v) > 6 }) {
    fmt.Println(v)
}
// Output: range over func
```

#### With `MapExtractor.ByValue`

```go
m := map[string]string{"language": "gopher", "design pattern": "iterator", "implementation": "range over func"}
for v := range xtract.FromMap(m).ByValue(func(v string) bool { return len(v) < 8 }) {
    fmt.Println(v)
}
// Output: gopher
```

#### With `MapExtractor.ByKey`

```go
m := map[string]string{"language": "gopher", "design pattern": "iterator", "implementation": "range over func"}
for v := range xtract.FromMap(m).ByKey(func(k string) bool { return strings.Contains(k, " ") }) {
    fmt.Println(v)
}
// Output: iterator
```

#### With `MapExtractor.ByKeyAndValue`

```go
m := map[string]string{"language": "gopher", "design pattern": "iterator", "implementation": "range over func"}
for v := range xtract.FromMap(m).ByKeyAndValue(func(k, v string) bool { return strings.Contains(k, "e") && len(v) < 8 }) {
    fmt.Println(v)
}
// Output: gopher
```

## Contributing

Feel free to open a PR or an Issue.

## License

**xtract** released under the [MIT License](https://github.com/miyamo2/xtract/blob/main/LICENSE)
