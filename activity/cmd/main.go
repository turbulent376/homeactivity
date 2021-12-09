// this file is generated by servgen util based on a template at 2021-06-26 10:37:24 +0300 MSK
package main

import (
	"context"
	kitContext "git.jetbrains.space/orbi/fcsd/kit/context"
	"git.jetbrains.space/orbi/fcsd/timesheet"
	"git.jetbrains.space/orbi/fcsd/timesheet/internal/logger"
	"google.golang.org/api/tasks/v1"
	"os"
	"os/signal"
	"syscall"
)

func main() {

	// init context
	ctx := kitContext.NewRequestCtx().Empty().WithNewRequestId().ToContext(context.Background())

	// create a new service
	s := activity.New()

	l := logger.L().Mth("main").Inf("created")

	// init service
	if err := s.Init(ctx); err != nil {
		l.E(err).St().Err("initialization")
		os.Exit(1)
	}

	l.Inf("initialized")

	// start listening
	if err := s.Start(ctx); err != nil {
		l.E(err).St().Err("listen")
		os.Exit(1)
	}

	l.Inf("listening")

	// handle app close
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt, syscall.SIGTERM)
	<-quit
	l.Inf("graceful shutdown")
	s.Close(ctx)
	os.Exit(0)
}