package goutils

import (
	"reflect"
	"testing"
)

func TestSet(t *testing.T) {
	t.Run("정수 슬라이스(중복 포함)", func(t *testing.T) {
		input := []int{1, 2, 2, 3, 1}
		expected := map[int]struct{}{1: {}, 2: {}, 3: {}}
		result := Set(input)
		if !reflect.DeepEqual(result, expected) {
			t.Errorf("Set() = %v, want %v", result, expected)
		}
	})

	t.Run("문자열 슬라이스(중복 없음)", func(t *testing.T) {
		input := []string{"a", "b", "c"}
		expected := map[string]struct{}{"a": {}, "b": {}, "c": {}}
		result := Set(input)
		if !reflect.DeepEqual(result, expected) {
			t.Errorf("Set() = %v, want %v", result, expected)
		}
	})

	t.Run("빈 슬라이스", func(t *testing.T) {
		input := []int{}
		expected := map[int]struct{}{}
		result := Set(input)
		if !reflect.DeepEqual(result, expected) {
			t.Errorf("Set() = %v, want %v", result, expected)
		}
	})
}

func TestContainsAll(t *testing.T) {
	t.Run("정수 - 모두 포함", func(t *testing.T) {
		source := []int{1, 2, 3, 4, 5}
		target := []int{2, 4}
		if !ContainsAll(source, target) {
			t.Errorf("ContainsAll(%v, %v) = false, want true", source, target)
		}
	})

	t.Run("정수 - 일부 미포함", func(t *testing.T) {
		source := []int{1, 2, 3, 4, 5}
		target := []int{2, 6}
		if ContainsAll(source, target) {
			t.Errorf("ContainsAll(%v, %v) = true, want false", source, target)
		}
	})

	t.Run("문자열 - 모두 포함", func(t *testing.T) {
		source := []string{"a", "b", "c", "d"}
		target := []string{"a", "c"}
		if !ContainsAll(source, target) {
			t.Errorf("ContainsAll(%v, %v) = false, want true", source, target)
		}
	})

	t.Run("문자열 - 일부 미포함", func(t *testing.T) {
		source := []string{"a", "b", "c", "d"}
		target := []string{"a", "e"}
		if ContainsAll(source, target) {
			t.Errorf("ContainsAll(%v, %v) = true, want false", source, target)
		}
	})

	t.Run("빈 타겟 슬라이스", func(t *testing.T) {
		source := []int{1, 2, 3}
		target := []int{}
		if !ContainsAll(source, target) {
			t.Errorf("빈 타겟 슬라이스를 전달하면 항상 true여야 합니다")
		}
	})

	t.Run("빈 소스 슬라이스", func(t *testing.T) {
		source := []int{}
		target := []int{1}
		if ContainsAll(source, target) {
			t.Errorf("빈 소스 슬라이스를 전달하면 항상 false여야 합니다 (타겟이 비어있지 않은 경우)")
		}
	})

	t.Run("두 슬라이스 모두 비어있음", func(t *testing.T) {
		source := []int{}
		target := []int{}
		if !ContainsAll(source, target) {
			t.Errorf("두 슬라이스가 모두 비어있으면 true여야 합니다")
		}
	})
}

func TestUnion(t *testing.T) {
	t.Run("정수 슬라이스", func(t *testing.T) {
		s1 := []int{1, 2, 3}
		s2 := []int{3, 4, 5}
		s3 := []int{5, 6, 7}
		result := Union(s1, s2, s3)
		// 순서가 보장되지 않으므로, set으로 변환하여 비교합니다.
		expectedSet := map[int]struct{}{1: {}, 2: {}, 3: {}, 4: {}, 5: {}, 6: {}, 7: {}}
		resultSet := Set(result)
		if !reflect.DeepEqual(resultSet, expectedSet) {
			t.Errorf("Union() = %v, want %v", result, expectedSet)
		}
		if len(result) != len(expectedSet) {
			t.Errorf("Union() length = %d, want %d", len(result), len(expectedSet))
		}
	})

	t.Run("문자열 슬라이스", func(t *testing.T) {
		s1 := []string{"a", "b"}
		s2 := []string{"b", "c", "d"}
		result := Union(s1, s2)
		expectedSet := map[string]struct{}{"a": {}, "b": {}, "c": {}, "d": {}}
		resultSet := Set(result)
		if !reflect.DeepEqual(resultSet, expectedSet) {
			t.Errorf("Union() = %v, want %v", result, expectedSet)
		}
	})

	t.Run("단일 슬라이스", func(t *testing.T) {
		s1 := []int{1, 1, 2, 2}
		result := Union(s1)
		expectedSet := map[int]struct{}{1: {}, 2: {}}
		resultSet := Set(result)
		if !reflect.DeepEqual(resultSet, expectedSet) {
			t.Errorf("Union() with single slice = %v, want %v", result, expectedSet)
		}
	})

	t.Run("빈 슬라이스 포함", func(t *testing.T) {
		s1 := []int{1, 2}
		s2 := []int{}
		s3 := []int{2, 3}
		result := Union(s1, s2, s3)
		expectedSet := map[int]struct{}{1: {}, 2: {}, 3: {}}
		resultSet := Set(result)
		if !reflect.DeepEqual(resultSet, expectedSet) {
			t.Errorf("Union() with empty slice = %v, want %v", result, expectedSet)
		}
	})

	t.Run("모두 빈 슬라이스", func(t *testing.T) {
		s1 := []int{}
		s2 := []int{}
		result := Union(s1, s2)
		if len(result) != 0 {
			t.Errorf("Union() of empty slices should be empty, got %v", result)
		}
	})
}
