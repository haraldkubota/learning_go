package rotLeft

import (
	"reflect"
	"testing"
)

func Test0RotLeft(t *testing.T) {
	source := []int{3, 4, 5, 6, 7, 8}
	want := []int{3, 4, 5, 6, 7, 8}
	res := rotLeft(source, 0)
	if !reflect.DeepEqual(res, want) {
		t.Fatalf(`Wanted %v, got %v`, source, res)
	}
}

func Test6RotLeft(t *testing.T) {
	source := []int{3, 4, 5, 6, 7, 8}
	want := []int{3, 4, 5, 6, 7, 8}
	res := rotLeft(source, 6)
	if !reflect.DeepEqual(res, want) {
		t.Fatalf(`Wanted %v, got %v`, source, res)
	}
}
func Test1RotLeft(t *testing.T) {
	source := []int{3, 4, 5, 6, 7, 8}
	rotCount := 1
	want := []int{4, 5, 6, 7, 8, 3}
	res := rotLeft(source, rotCount)
	if !reflect.DeepEqual(res, want) {
		t.Fatalf(`Wanted %v, got %v`, want, res)
	}
}

func Test4RotLeft(t *testing.T) {
	source := []int{3, 4, 5, 6, 7, 8}
	rotCount := 4
	want := []int{7, 8, 3, 4, 5, 6}
	res := rotLeft(source, rotCount)
	if !reflect.DeepEqual(res, want) {
		t.Fatalf(`Wanted %v, got %v`, want, res)
	}
}
