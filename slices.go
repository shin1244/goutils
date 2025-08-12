package goutils

// SET은 source 슬라이스의 요소들의 중복을 제거합니다.
// 결과값은 source 슬라이스의 요소들을 키로 하고, 빈 구조체를 값으로 가지는 맵입니다.
func Set[T comparable](source []T) map[T]struct{} {
	result := make(map[T]struct{}, len(source))
	for _, item := range source {
		result[item] = struct{}{}
	}

	return result
}

// ContainsAll은 target 슬라이스의 모든 요소가 source 슬라이스에 포함되어 있는지 확인합니다.
func ContainsAll[T comparable](source, target []T) bool {
	sourceSet := Set(source)
	for _, item := range target {
		if _, ok := sourceSet[item]; !ok {
			return false
		}
	}

	return true
}

// ContainsAny는 target 슬라이스의 요소 중 하나라도 source 슬라이스에 포함되어 있는지 확인합니다.
func ContainsAny[T comparable](source, target []T) bool {
	sourceSet := Set(source)
	for _, item := range target {
		if _, ok := sourceSet[item]; ok {
			return true
		}
	}

	return false
}

// Union은 여러 개의 슬라이스를 입력받아 모든 고유한 요소들을 포함하는 단일 슬라이스를 반환합니다.
// 순서가 보장되지 않습니다.
func Union[T comparable](slices ...[]T) []T {
	unionSet := make(map[T]struct{})
	for _, slice := range slices {
		for _, item := range slice {
			unionSet[item] = struct{}{}
		}
	}

	result := make([]T, 0, len(unionSet))
	for item := range unionSet {
		result = append(result, item)
	}

	return result
}

// Intersect는 여러 개의 슬라이스를 입력받아 모든 슬라이스에 공통적으로 존재하는 요소들을 포함하는 단일 슬라이스를 반환합니다.
// 순서가 보장되지 않습니다.
func Intersect[T comparable](slices ...[]T) []T {
	intersectSet := Set(slices[0])

	for _, slice := range slices[1:] {
		currentSet := Set(slice)
		for item := range intersectSet {
			if _, ok := currentSet[item]; !ok {
				delete(intersectSet, item)
			}
		}
	}
	result := make([]T, 0, len(intersectSet))
	for item := range intersectSet {
		result = append(result, item)
	}

	return result
}

// RemoveMany는 source 슬라이스에서 target 슬라이스의 모든 요소를 제거한 새로운 슬라이스를 반환합니다.
// target 슬라이스의 요소들은 source 슬라이스에 존재하지 않을 수도 있습니다.
func RemoveMany[T comparable](source []T, target []T) []T {
	targetSet := Set(target)
	result := make([]T, 0, len(source))

	for _, item := range source {
		if _, ok := targetSet[item]; !ok {
			result = append(result, item)
		}
	}

	return result
}
