package goutils

import "strings"

// SplitAndTrim는 문자열을 쉼표로 분리하고 공백을 제거하여 슬라이스로 반환합니다.
func SplitAndTrim(s string) []string {
	parts := strings.Split(s, ",")
	var result []string
	for _, part := range parts {
		trimmed := strings.TrimSpace(part)
		if trimmed != "" {
			result = append(result, trimmed)
		}
	}
	return result
}
