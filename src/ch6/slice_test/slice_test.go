package slice_test

import "testing"

func TestSliceInit(t *testing.T) {
	var s1 []int
	t.Log(len(s1), cap(s1))
	s1 = append(s1, 100)
	t.Log(len(s1), cap(s1))

	var s2 = []int{10, 20, 30}
	t.Log(len(s2), cap(s2))

	var s3 = make([]int, 3, 5)
	t.Log(len(s3), cap(s3))

	t.Log(s3[0], s3[1], s3[2])
	s3 = append(s3, 22)
	t.Log(s3[0], s3[1], s3[2], s3[3])
	t.Log(s3[0], s3[1], s3[2], s3[3], s3[4])
}

func TestSliceTravel(t *testing.T) {
	var s1 []int

	for i := 0; i < 10; i++ {
		s1 = append(s1, i)
		t.Log(len(s1), cap(s1))
	}
}
