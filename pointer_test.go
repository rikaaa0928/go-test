package go_test

import "testing"
import "github.com/rikaaa0928/rtools/generic"

type testPointer struct {
	X []int
	Y *testPointer
}

func TestPointPaas(t *testing.T) {
	x := &testPointer{X: []int{1}, Y: generic.TakeAddr(testPointer{X: []int{1}})}
	t.Log(x.X, *x.Y)
	y := *x
	y.X = append(y.X, 2)
	y.Y.X[0] = 2
	x.X = append(x.X, 3)
	x.Y.X[0] = 3
	t.Log(x.X, *x.Y, y.X, *y.Y)
}

func TestPointerPrint(t *testing.T) {
	a := generic.TakeAddr(1)
	//t.Logf("%p", a)
	b := a
	t.Logf("%p", a)
	t.Logf("%p", &a)
	t.Logf("%p", b)
	t.Logf("%p", &b)
	l := []int{1, 2}
	t.Logf("%p", l)
	t.Logf("%p", &l[0])
	for i, v := range l {
		t.Logf("%p,%p", &v, &l[i])
	}
	pl := []*int{generic.TakeAddr(1), generic.TakeAddr(2)}
	t.Logf("%p", pl)
	t.Logf("%p", &pl)
	t.Logf("%p", pl[0])
	t.Logf("%p", &pl[0])
	for i, v := range pl {
		t.Logf("%p,%p,%p,%p", v, pl[i], &v, &l[i])
	}
}
