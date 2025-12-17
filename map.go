package goutils

// ReverseMap는 주어진 맵의 키와 값을 뒤바꾼 새로운 맵을 반환합니다.
func ReverseMap[T, U comparable](m map[T]U) map[U]T {
	res := make(map[U]T)
	for k, v := range m {
		res[v] = k
	}
	return res
}
