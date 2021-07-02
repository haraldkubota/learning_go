package medianSorted

import (
	// "fmt"
	"reflect"
	"testing"
)

func TestSingleItem(t *testing.T) {
	source1 := []int{}
	source2 := []int{2}
	want := 2.0
	res := findMedianSortedArrays(source1, source2)
	if want != res {
		t.Fatalf(`Wanted %f, got %f`, want, res)
	}
}

func TestSixItem(t *testing.T) {
	source1 := []int{4, 7, 9}
	source2 := []int{3, 4, 5}
	want := 4.5
	res := findMedianSortedArrays(source1, source2)
	if want != res {
		t.Fatalf(`Wanted %f, got %f`, want, res)
	}
}

func TestFiveItem(t *testing.T) {
	source1 := []int{1, 9}
	source2 := []int{2, 3, 4}
	want := 3.0
	res := findMedianSortedArrays(source1, source2)
	if want != res {
		t.Fatalf(`Wanted %f, got %f`, want, res)
	}
}

func TestLeftEmptyItem(t *testing.T) {
	source1 := []int{}
	source2 := []int{2, 3, 4}
	want := 3.0
	res := findMedianSortedArrays(source1, source2)
	if want != res {
		t.Fatalf(`Wanted %f, got %f`, want, res)
	}
}

func TestRightEmptyItem(t *testing.T) {
	source1 := []int{1, 9}
	source2 := []int{}
	want := 5.0
	res := findMedianSortedArrays(source1, source2)
	if want != res {
		t.Fatalf(`Wanted %f, got %f`, want, res)
	}
}

func TestGetNextLong(t *testing.T) {

	type testMerge struct {
		s1, s2, want []int
	}
	tests := []testMerge{
		testMerge{
			[]int{1, 9}, []int{}, []int{1, 9},
		},
		testMerge{
			[]int{1, 3}, []int{2, 4, 6, 8}, []int{1, 2, 3, 4, 6, 8},
		},
		testMerge{
			[]int{}, []int{}, []int{},
		},
		testMerge{
			[]int{2, 3}, []int{1, 4, 6, 8}, []int{1, 2, 3, 4, 6, 8},
		},
	}
	// fmt.Printf("tests=%v\n", tests)
	for i := 0; i < len(tests); i += 1 {
		res := []int{}
		i1 := 0
		i2 := 0
		var val int
		var err error
		for j := 0; j < len(tests[i].s1)+len(tests[i].s2); j += 1 {
			i1, i2, val, err = getNext(i1, i2, tests[i].s1, tests[i].s2)
			if err != nil {
				t.Fatalf("Running out of elements")
			}
			res = append(res, val)
		}
		if !reflect.DeepEqual(tests[i].want, res) {
			t.Fatalf(`Wanted %v, got %v`, tests[i].want, res)
		}
	}
}

func TestFindMedianFast(t *testing.T) {

	type testMedian struct {
		s1, s2 []int
		want   float64
	}
	tests := []testMedian{
		testMedian{
			[]int{1, 9}, []int{}, 5.0,
		},
		testMedian{
			[]int{1, 3}, []int{2, 4, 6, 8}, 3.5,
		},
		testMedian{
			[]int{2, 3}, []int{1, 4, 6}, 3.0,
		},
		testMedian{
			[]int{1, 2, 3}, []int{1}, 1.5,
		},
	}
	for i := 0; i < len(tests); i += 1 {
		res := findMedianSortedArraysFast(tests[i].s1, tests[i].s2)
		if tests[i].want != res {
			t.Fatalf(`Wanted %v, got %v`, tests[i].want, res)
		}
	}
}
