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
	s *configs.Snippet
	c *cron.Cron
}

func NewJob(s *configs.Snippet) *Job {
	return &Job{
		s: s,
		c: cron.New(
			cron.WithLogger(
				cron.VerbosePrintfLogger(
					log.New(os.Stdout, "cron: ", log.LstdFlags),
				),
			),
		)}
}

func (j *Job) Start(ctx context.Context) error {
	do := func() {
		if err := NewTrip(j.s).treatSnippet(); err != nil {
			log.Println(err)
		}
	}
	if configs.V.Debug {
		do() // just working at started.
	}

	j.c.AddFunc(j.s.Cron, do)
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
