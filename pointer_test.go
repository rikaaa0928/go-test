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
