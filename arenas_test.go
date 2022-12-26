//go:build goexperiment.arenas

package go_test

import "testing"
import "arena"

func TestArenas(t *testing.T) {
	a := arena.NewArena()
	x := arena.New[int](a)
	t.Log(*x)
	a.Free()
	t.Log(*x)
}
