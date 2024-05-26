package ch13

import "testing"

func TestSquare(t *testing.T) {
	var inputParam = [...]int{1, 2, 3}
	excepted := [...]int{1, 4, 10}
	for i, v := range inputParam {
		result := Square(v)
		if result != excepted[i] {
			t.Errorf("Square(%d) = %d, want %d", v, result, excepted[i])
		}
	}

}
