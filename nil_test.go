package go_test

import (
	"fmt"
	"testing"
)

type nilInter interface {
	a()
}
type nilImpl struct {
	int
}

func (i *nilImpl) a() {
	fmt.Println(i)
}
func TestNil(t *testing.T) {
	var x *nilImpl
	if x != nil {
		t.Fatal("x!=nil!!")
	}
	var y nilInter = x
	if y == nil {
		t.Fatal("y==nil!!")
	}
	x.a()
	y.a()
	var z nilInter
	if z != nil {
		t.Fatal("z!=nil!!")
	}
	z = nil
	if z != nil {
		t.Fatal("z!=nil!!")
	}
	var a nilInter = nil
	if a != nil {
		t.Fatal("a!=nil!!")
	}
}
