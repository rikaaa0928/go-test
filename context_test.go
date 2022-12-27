package go_test

import (
	"context"
	"github.com/rikaaa0928/rtools/generic"
	"testing"
)

func TestContext(t *testing.T) {
	ctx := context.Background()
	ctx = context.WithValue(ctx, "key", "1")
	t.Log(ctx.Value("key"))
	for i := 0; i < 10; i++ {
		ctx = context.WithValue(ctx, "key", i)
		t.Log(ctx.Value("key"))
	}
	ctxInner(ctx, t)
	t.Log(ctx.Value("key"))
	ctx = context.WithValue(ctx, "pointer", generic.TakeAddr(10))
	t.Log(*ctx.Value("pointer").(*int))
	ctxPointerInner(ctx, t)
	t.Log(*ctx.Value("pointer").(*int))
}

func ctxInner(ctx context.Context, t *testing.T) {
	ctx = context.WithValue(ctx, "key", "x")
	t.Log(ctx.Value("key"))
}

func ctxPointerInner(ctx context.Context, t *testing.T) {
	p, ok := ctx.Value("pointer").(*int)
	if ok {
		*p = 11
		t.Log(*ctx.Value("pointer").(*int))
	}
	_, ok = ctx.Value("key").(*int)
	if ok {
		t.Error("should not ok")
	}
}
