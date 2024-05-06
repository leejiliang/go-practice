package condition

import "testing"

func TestIfSec(t *testing.T) {
	if a := 1 == 1; a {
		t.Log("1==1")
	} else {
		t.Log("1!=1")
	}
}

func TestSwitch(t *testing.T) {
	for i := 0; i < 10; i++ {
		switch i {
		case 0, 1, 2:
			t.Log("aaaaa")
		case 3, 4, 5:
			t.Log("bbbbb")
		default:
			t.Log("ccccc")

		}
	}
}
