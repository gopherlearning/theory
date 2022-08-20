package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestAbs(t *testing.T) {
	tests := []struct { // добавился слайс тестов
		name  string
		value float64
		want  float64
	}{
		{
			name: "Test Task 1 Example 1 -3", value: -3, want: 3,
		},
		{
			name: "Test Task 1 Example 1 3", value: 3, want: 3,
		},
		{
			name: "Test Task 1 Example 1 -2.000001", value: -2.000001, want: 2.000001,
		},
		{
			name: "Test Task 1 Example 1 -0.000000003", value: -0.000000003, want: 0.000000003,
		},
	}
	for _, tt := range tests { // цикл по всем тестам
		t.Run(tt.name, func(t *testing.T) {
			v := Abs(tt.value)
			assert.Equal(t, tt.want, v)
		})
	}
}
