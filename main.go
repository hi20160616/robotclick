package main

import (
	"context"
	"log"

	"golang.org/x/sync/errgroup"
)

func main() {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	g, ctx := errgroup.WithContext(ctx)

	// Jobs
	j := NewJob()
	g.Go(func() error {
		log.Println("Jobs start")
		return j.Start(ctx)
	})
	g.Go(func() error {
		defer log.Println("Jobs stop done.")
		<-ctx.Done() // wait for stop signal
		log.Println("Jobs stop now...")
		return j.Stop(ctx) // TODO: stop cron and ever
	})

}
