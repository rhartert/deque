package queues

import (
	"reflect"
	"testing"
)

func TestNewWithCapa_initCapa(t *testing.T) {
	testCases := []struct {
		n    int
		want int
	}{
		{-10, 16},
		{0, 16},
		{1, 16},
		{2, 16},
		{15, 16},
		{16, 16},
		{17, 32},
		{31, 32},
		{32, 32},
		{142, 256},
	}

	for _, tc := range testCases {
		q := NewWithCapa[int](tc.n)

		if got := len(q.ring); got != tc.want {
			t.Errorf("NewWithCapa(%d): want capa %d, got %d", tc.n, tc.want, got)
		}
	}
}

func TestQueue_Enqueue_ResizeNoRotation(t *testing.T) {
	q := &Queue[int]{
		ring:  []int{1, 2, 3, 4},
		start: 0,
		end:   0,
		size:  4,
		mask:  0b11,
	}
	want := &Queue[int]{
		ring:  []int{1, 2, 3, 4, 5, 0, 0, 0},
		start: 0,
		end:   5,
		size:  5,
		mask:  0b111,
	}

	q.Enqueue(5)

	if !reflect.DeepEqual(want, q) {
		t.Errorf("Mismatch: want %#v, got %#v", want, q)
	}
}

func TestQueue_Enqueue_ResizeAndRotation(t *testing.T) {
	q := &Queue[int]{
		ring:  []int{3, 4, 1, 2},
		start: 2,
		end:   2,
		size:  4,
		mask:  0b11,
	}
	want := &Queue[int]{
		ring:  []int{1, 2, 3, 4, 5, 0, 0, 0},
		start: 0,
		end:   5,
		size:  5,
		mask:  0b111,
	}

	q.Enqueue(5)

	if !reflect.DeepEqual(want, q) {
		t.Errorf("Mismatch: want %#v, got %#v", want, q)
	}
}
