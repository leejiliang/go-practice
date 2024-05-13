package map_test

import "testing"

func TestMapInit(t *testing.T) {
	m := map[string]int{"A": 1, "B": 2, "C": 3, "D": 4, "E": 5}
	t.Log(len(m), m["A"])

	m2 := map[string]int{}
	m2["A"] = 4
	t.Log(len(m2), m2["A"])

	m3 := make(map[string]int, 10) // cap 10
	m3["A"] = 5
	t.Logf("m3 len:%d", len(m3))
	//t.Logf("m3 len:%d", cap(m3))// 无法通过cap()获取cap
}

func TestMapExistsKey(t *testing.T) {
	map1 := map[int]string{1: "a", 2: "b", 3: "c"}

	if v, ok := map1[4]; ok {
		t.Logf("key:%s exists:%t", v, ok)
	} else {
		t.Logf("key:%s not exists:%t", v, ok)
	}
}

func TestMapWithFunc(t *testing.T) {
	map2 := map[int]func(int) int{}
	map2[1] = func(i int) int { return i }
	map2[2] = func(i int) int { return i * i }
	map2[3] = func(i int) int { return i * i * i }

	t.Log(map2[1](2), map2[2](2), map2[3](2))

}
