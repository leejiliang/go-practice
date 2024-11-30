package main

import (
	"math/rand"
	"testing"
)

func TestAAA(t *testing.T) {
	str := []string{"爱人", "保姆", "傻蛋"}
	t.Log(str[rand.Intn(len(str))])
}
