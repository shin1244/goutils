package goutils

// source 슬라이스의 요소들의 중복을 제거합니다.
// 결과값은 source 슬라이스의 요소들을 키로 하고, 빈 구조체를 값으로 가지는 맵입니다.
func Set[T comparable](source []T) map[T]struct{} {
	result := make(map[T]struct{}, len(source))
	for _, item := range source {
		result[item] = struct{}{}
	}

	return result
}

// target 슬라이스의 모든 요소가 source 슬라이스에 포함되어 있는지 확인합니다.
func ContainsAll[T comparable](source, target []T) bool {
	sourceSet := Set(source)
	for _, item := range target {
		if _, ok := sourceSet[item]; !ok {
			return false
		}
	}

	return true
}

// target 슬라이스의 요소 중 하나라도 source 슬라이스에 포함되어 있는지 확인합니다.
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
