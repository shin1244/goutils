package goutils

// ReverseMap은 문자열을 키로 가지는 맵을 입력 받아 그 값을 키로, 키를 값으로 가지는 맵을 반환합니다.
func ReverseMap[T comparable](m map[string]T) map[T]string {
	res := make(map[T]string)
	for k, v := range m {
		res[v] = k
	}
	return res
}
