package go_test

import (
	"context"
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
}
