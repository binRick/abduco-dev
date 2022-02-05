package abducoctl

import (
	"context"
	"fmt"
	"time"
)

func Loop() {
	ctx, cancel := context.WithCancel(ctx)
	defer cancel()
	fmt.Println("ok")
	time.Sleep(10 * time.Millisecond)
	Connect(ctx, "")
}
