package array_test

import "testing"

func TestArray(t *testing.T) {
	var arr [3]int
	//arr1 := [4]int{1, 2, 3, 4}
	//arr2 := [...]int{1, 2, 3, 4, 5}
	t.Log(arr[0], arr[1], arr[2])
}

func TestIteratorArray(t *testing.T) {
	arr := [...]int{11, 12, 13, 14, 15}
	for i := range arr {
		t.Log(arr[i])
	}

	for i, e := range arr {
		t.Log(i, e)
	}
}

func TestArraySlice(t *testing.T) { // 规则, 含头部含尾, 不支持负数
	arr := [...]int{11, 12, 13, 14, 15}
	arr1 := arr[1:2]
	arr2 := arr[:3]
	arr3 := arr[3:]
	arr4 := arr[:]
	t.Log(arr1, arr2, arr3, arr4)
}
