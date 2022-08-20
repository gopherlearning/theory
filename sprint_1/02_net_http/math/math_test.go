package math

import "testing"

// Sum возвращает сумму элементов.
func Sum(values ...int) int {
	var sum int
	for _, v := range values {
		sum += v
	}
	return sum
}

func TestSum(t *testing.T) {
	tests := []struct { // добавился слайс тестов
		name   string
		values []int
		want   int
	}{
		{
			name:   "simple test #1", // описывается каждый тест
			values: []int{1, 2},      // значения, которые будет принимать функция
			want:   3,                // ожидаемое значение
		},
		{
			name:   "one",
			values: []int{1},
			want:   1,
		},
		{
			name:   "with negative values",
			values: []int{-1, -2, 3},
			want:   0,
		},
		{
			name:   "with negative zero",
			values: []int{-0, 3},
			want:   3,
		},
		{
			name:   "a lot of values",
			values: []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16, 17, 18, 18},
			want:   189,
		},
	}
	for _, tt := range tests { // цикл по всем тестам
		t.Run(tt.name, func(t *testing.T) {
			if sum := Sum(tt.values...); sum != tt.want {
				t.Errorf("Add() = %v, want %v", sum, tt.want)
			}
		})
	}
}
