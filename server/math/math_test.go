package math

import "testing"

func TestLocalAdd(t *testing.T) {
	result := LocalAdd(2.5, 3.5)
	expected := 6.0
	if result != expected {
		t.Errorf("LocalAdd(2.5, 3.5) = %v; want %v", result, expected)
	}
}

func TestLocalSubtract(t *testing.T) {
	result := LocalSubtract(10, 3)
	expected := 7.0
	if result != expected {
		t.Errorf("LocalSubtract(10, 3) = %v; want %v", result, expected)
	}
}

func TestLocalFindMin(t *testing.T) {
	tests := []struct {
		a, b, c int64
		want    int64
	}{
		{3, 5, 7, 3},
		{10, 2, 8, 2},
		{9, 9, 9, 9}, // tests fallback case
	}

	for _, tt := range tests {
		got := LocalFindMin(tt.a, tt.b, tt.c)
		if got != tt.want {
			t.Errorf("LocalFindMin(%d, %d, %d) = %d; want %d",
				tt.a, tt.b, tt.c, got, tt.want)
		}
	}
}

func TestLocalFindMax(t *testing.T) {
	tests := []struct {
		a, b, c int64
		want    int64
	}{
		{3, 5, 7, 7},
		{10, 2, 8, 10},
		{9, 9, 9, 9}, // tests fallback case
	}

	for _, tt := range tests {
		got := LocalFindMax(tt.a, tt.b, tt.c)
		if got != tt.want {
			t.Errorf("LocalFindMax(%d, %d, %d) = %d; want %d",
				tt.a, tt.b, tt.c, got, tt.want)
		}
	}
}
