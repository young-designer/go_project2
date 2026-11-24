package main

import "testing"

func TestGenerateRandomElements(t *testing.T) {
	tests := []struct {
		name    string
		size    int
		wantLen int
	}{
		{name: "обычный размер", size: 10, wantLen: 10},
		{name: "большой размер", size: 1000, wantLen: 1000},
		{name: "нулевой размер", size: 0, wantLen: 0},
		{name: "отрицательный размер", size: -5, wantLen: 0},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := generateRandomElements(tt.size)
			if len(got) != tt.wantLen {
				t.Fatalf("generateRandomElements(%d) вернула слайс длиной %d, ожидалось %d", tt.size, len(got), tt.wantLen)
			}

			for i, value := range got {
				if value <= 0 {
					t.Fatalf("элемент с индексом %d не является положительным: %d", i, value)
				}
			}
		})
	}
}

func TestMaximum(t *testing.T) {
	tests := []struct {
		name string
		data []int
		want int
	}{
		{name: "пустой слайс", data: []int{}, want: 0},
		{name: "один элемент", data: []int{42}, want: 42},
		{name: "максимум в середине", data: []int{1, 5, 3, 2}, want: 5},
		{name: "максимум в конце", data: []int{1, 2, 3, 100}, want: 100},
		{name: "максимум в начале", data: []int{100, 2, 3, 4}, want: 100},
		{name: "все элементы равны", data: []int{7, 7, 7}, want: 7},
		{name: "отрицательные числа", data: []int{-10, -3, -42}, want: -3},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := maximum(tt.data); got != tt.want {
				t.Fatalf("maximum(%v) = %d, ожидалось %d", tt.data, got, tt.want)
			}
		})
	}
}

func TestMaxChunks(t *testing.T) {
	tests := []struct {
		name string
		data []int
	}{
		{name: "пустой слайс", data: []int{}},
		{name: "один элемент", data: []int{42}},
		{name: "меньше восьми элементов", data: []int{3, 1, 4, 1, 5, 9, 2}},
		{name: "ровно восемь элементов", data: []int{8, 7, 6, 5, 4, 3, 2, 1}},
		{name: "девять элементов", data: []int{9, 8, 7, 6, 5, 4, 3, 2, 1}},
		{name: "много элементов", data: generateRandomElements(1000)},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			want := maximum(tt.data)
			if got := maxChunks(tt.data); got != want {
				t.Fatalf("maxChunks вернула %d, ожидалось %d", got, want)
			}
		})
	}
}
