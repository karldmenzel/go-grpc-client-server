package math

import "testing"

func TestMagicAdd(t *testing.T) {
	result := MagicAdd(2.5, 3.5)
	expected := 6.0
	if result != expected {
		t.Errorf("MagicAdd(2.5, 3.5) = %v; want %v", result, expected)
	}
}

func TestMagicSubtract(t *testing.T) {
	result := MagicSubtract(10, 3)
	expected := 7.0
	if result != expected {
		t.Errorf("MagicSubtract(10, 3) = %v; want %v", result, expected)
	}
}

func TestMagicFindMin(t *testing.T) {
	tests := []struct {
		a, b, c int64
		want    int64
	}{
		{3, 5, 7, 3},
		{10, 2, 8, 2},
		{9, 9, 9, 9}, // tests fallback case
	}

	for _, tt := range tests {
		got := MagicFindMin(tt.a, tt.b, tt.c)
		if got != tt.want {
			t.Errorf("MagicFindMin(%d, %d, %d) = %d; want %d",
				tt.a, tt.b, tt.c, got, tt.want)
		}
	}
}

func TestMagicFindMax(t *testing.T) {
	tests := []struct {
		a, b, c int64
		want    int64
	}{
		{3, 5, 7, 7},
		{10, 2, 8, 10},
		{9, 9, 9, 9}, // tests fallback case
	}

	for _, tt := range tests {
		got := MagicFindMax(tt.a, tt.b, tt.c)
		if got != tt.want {
			t.Errorf("MagicFindMax(%d, %d, %d) = %d; want %d",
				tt.a, tt.b, tt.c, got, tt.want)
		}
	}
}
