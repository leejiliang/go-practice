package ch9

import (
	"fmt"
	"testing"
)

type Programme interface {
	HelloWorld() string
}

type GoProgramme struct{}
type JavaProgramme struct{}

func (GoProgramme) HelloWorld() string {
	return "goProgramme, Hello World"
}

func (JavaProgramme) HelloWorld() string {
	return "javaProgramme, Hello World"
}

func SayHello(p Programme) {
	fmt.Printf("%T, %v  \n", p, p.HelloWorld())
}

func TestPP(t *testing.T) {
	goP := new(GoProgramme)
	javaP := new(JavaProgramme)
	SayHello(goP)
	SayHello(javaP)
}
