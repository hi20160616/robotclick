package main

import (
	"context"
	"fmt"
	"log"
	"os"
	"time"

	"github.com/hi20160616/robotclick/configs"
	"github.com/robfig/cron/v3"
)

type Job struct {
	c *cron.Cron
}

func NewJob() *Job {
	return &Job{cron.New(
		cron.WithLogger(
			cron.VerbosePrintfLogger(
				log.New(os.Stdout, "cron: ", log.LstdFlags),
			),
		),
	)}
}

func (j *Job) Start(ctx context.Context) error {
	do := func() {
		log.Println("Job start.")
		if err := NewTrip().working(); err != nil {
			log.Println(err)
		}
	}
	if configs.V.Debug {
		do() // just working at started.
	}

	j.c.AddFunc(configs.V.Cron, do)
	j.c.Start()
	return ctx.Err()
}

func (j *Job) Stop(ctx context.Context) error {
	ctx = j.c.Stop()
	select {
	case <-ctx.Done():
		return ctx.Err()
	case <-time.After(30 * time.Second):
		return fmt.Errorf("context was not done immediately.")
	}
}
