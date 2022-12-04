package main

import (
	"context"

	"github.com/go-cmd/cmd"
	"go.uber.org/zap"
)

func runCmd(ctx context.Context, cancel context.CancelFunc, log *zap.SugaredLogger, env []string, exec string, args ...string) error {
	runCmd := cmd.NewCmdOptions(
		cmd.Options{
			Streaming: true,
			Buffered:  true,
		},
		exec,
		args...,
	)
	runCmd.Env = env

	statusChan := runCmd.Start()

	for {
		select {
		case out := <-runCmd.Stdout:
			log.Info(out)
		case out := <-runCmd.Stderr:
			log.Error(out)
		case <-statusChan:
			cancel()
		case <-ctx.Done():
			runCmd.Stop()

			status := <-statusChan

			return status.Error
		}
	}
}
