# Changelog

## 0.3.0 - 2024-07-26

### ✨New Features

- Added `Limit` to limit the number of values in a sequence to a specified number
- Added `Offset` to skip a specified number of values in a sequence

### 📚Documentation

- Few Fixes in GoDoc and README

## 0.2.1 - 2024-07-23

### Refactor🔨

- Generation of `iter.Seq` and `iter.Seq2` is now common as an internal function

### Documentation📚

- Added more example to GoDoc and README

## 0.2.0 - 2024-06-23

### Breaking Changes💥

- `Extractor[T, U]` is now method-chain interface

## 0.1.1 - 2024-06-16

### Refactor♻️

- Modified the return value to `iter.Seq[T]`, the alias type

### Go Version🐭

- Go Version are now not specified patch version

## 0.1.0 - 2024-06-15

### Initial Release🎉

- `SliceExtractor[T].ByValue`
- `SliceExtractor[T].ByKey`
- `SliceExtractor[T].ByKeyAndValue`
- `MapExtractor[T, U].ByValue`
- `MapExtractor[T, U].ByKey`
- `MapExtractor[T, U].ByKeyAndValue`
