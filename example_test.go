package example

import (
	"context"
	"fmt"
	"os/exec"
	"time"
)

func ExampleCommandContext() {
	ctx, cancel := context.WithTimeout(context.Background(), 100*time.Millisecond)
	defer cancel()

	cmd := exec.CommandContext(ctx, "sleep", "5")
	started := time.Now()
	err := cmd.Run()

	fmt.Printf("aborted after %.1f seconds\n", time.Since(started).Seconds())
	fmt.Printf("run error: %v\n", err)
	fmt.Printf("ctx error: %v\n", ctx.Err())
	// aborted after 0.1 seconds
	// run error: signal: killed
	// ctx error: context deadline exceeded
}
