# GOUTILS

Go 언어를 위한 가볍고 실용적인 유틸리티 라이브러리입니다.

## Installation

```bash
go get github.com/shin1244/goutils

```

## Features

### 🫠Slices

**`Set`**

```go
func Set[T comparable](source []T) map[T]struct{}

```

> 슬라이스의 중복 요소를 제거합니다. `source`의 요소를 키로 하고 빈 구조체를 값으로 가지는 맵을 반환합니다.

**`ContainsAll`**

```go
func ContainsAll[T comparable](source, target []T) bool

```

> `target` 슬라이스의 모든 요소가 `source` 슬라이스에 포함되어 있는지 확인합니다.

**`Union`**

```go
func Union[T comparable](slices ...[]T) []T

```

> 여러 슬라이스를 입력받아 모든 고유한 요소를 포함하는 단일 슬라이스를 반환합니다. (순서 미보장)

**`Intersect`**

```go
func Intersect[T comparable](slices ...[]T) []T

```

> 입력받은 모든 슬라이스에 공통적으로 존재하는 요소들만 포함하는 단일 슬라이스를 반환합니다. (순서 미보장)

**`RemoveMany`**

```go
func RemoveMany[T comparable](source, target []T) []T

```

> `source` 슬라이스에서 `target` 슬라이스에 있는 모든 요소를 제거한 새로운 슬라이스를 반환합니다.

### 🫠Strings

**`SplitAndTrim`**

```go
func SplitAndTrim(s string) []string

```

> 문자열을 쉼표(`,`)로 분리하고, 각 요소의 앞뒤 공백을 제거하여 슬라이스로 반환합니다.

### 🫠Map

**`ReverseMap`**

```go
func ReverseMap[T, U comparable](m map[T]U) map[U]T

```

> 주어진 맵의 키(Key)와 값(Value)을 뒤바꾼 새로운 맵을 반환합니다.
