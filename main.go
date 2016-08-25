package main

import (
	"context"
	"log"
	"os/exec"
	"time"
)

func main() {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	cmd := exec.CommandContext(ctx, "sleep", "20")

	log.Println("started")

	err := cmd.Run()
	if err != nil {
		log.Printf("finished with error: %s\n", err)
		log.Printf("ctx err: %s\n", ctx.Err())
	}

}
