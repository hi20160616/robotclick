package main

import (
	"context"
	"errors"
	"fmt"
	"log"
	"os"
	"os/signal"
	"syscall"

	"github.com/hi20160616/robotclick/configs"
	"golang.org/x/sync/errgroup"
)

func main() {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	g, ctx := errgroup.WithContext(ctx)

	// Jobs
	for _, s := range configs.V.Snippets.Ss {
		j := NewJob(&s)
		g.Go(func() error {
			log.Printf("Job [\"%s\"] start\n", j.s.FileName)
			return j.Start(ctx)
		})
		g.Go(func() error {
			defer log.Printf("Job [\"%s\"] stop done.\n", j.s.FileName)
			<-ctx.Done() // wait for stop signal
			log.Printf("Job [\"%s\"] stop now...\n", j.s.FileName)
			return j.Stop(ctx) // TODO: stop cron and ever
		})

	}

	// Graceful stop
	sigs := make(chan os.Signal, 1)
	signal.Notify(sigs, syscall.SIGTERM, syscall.SIGINT, syscall.SIGQUIT)
	g.Go(func() error {
		select {
		case sig := <-sigs:
			fmt.Println()
			log.Printf("signal caught: %s, ready to quit...", sig.String())
			cancel()
		case <-ctx.Done():
			return ctx.Err()
		}
		return nil
	})

	if err := g.Wait(); err != nil {
		if !errors.Is(err, context.Canceled) {
			log.Printf("not canceled by context: %s", err)
		} else {
			log.Println(err)
		}
	}
}
