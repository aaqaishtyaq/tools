package main

import (
	"context"
	"fmt"
	"os"
	"os/signal"
	"syscall"
	"time"

	"go.uber.org/zap"
	"golang.org/x/sync/errgroup"
)

const (
	acme   = "acme"
	acmego = "acmego"
	font   = "/mnt/font/GoMono/16a/font"
	rc     = "bin/rc"
	ahead  = 5000
)

func main() {
	ctx := context.Background()

	// wait for the interrupt signal to gracefully shutdown the server
	// kill (no param) sends syscall.SIGTERM (default)
	// kill  -2	sends syscall.SIGINT
	// kill  -9 sends syscall.SIGKILL (but can't catch it, so don't need it)
	ctx, stop := signal.NotifyContext(ctx, syscall.SIGINT, syscall.SIGTERM)
	defer stop()

	log := SetUpLogs()

	e, err := env()
	if err != nil {
		log.Fatal(err)
	}

	arg := os.Args[1]

	g, gCtx := errgroup.WithContext(ctx)

	g.Go(func() error {
		return runAcme(gCtx, stop, arg, e, log)
	})

	g.Go(func() error {
		time.Sleep(time.Millisecond * ahead)

		return runAcmeGo(gCtx, stop, arg, e, log)
	})

	<-gCtx.Done()

	log.Info("Received interrupt signal")
}

func runAcme(ctx context.Context, stop context.CancelFunc, arg string, e []string, log *zap.SugaredLogger) error {
	return runCmd(
		ctx,
		stop,
		log.Named("acme"),
		e,
		acme,
		"-a",
		"-f",
		font,
		arg,
	)
}

func runAcmeGo(ctx context.Context, stop context.CancelFunc, arg string, e []string, log *zap.SugaredLogger) error {
	return runCmd(
		ctx,
		stop,
		log.Named("acmego"),
		e,
		acmego,
		"-f",
		arg,
	)
}

func env() ([]string, error) {
	envs := os.Environ()
	r, err := plan9Env()
	if err != nil {
		return nil, err
	}

	envs = append(envs, r)

	return envs, nil
}

func plan9Env() (string, error) {
	e := os.Getenv("PLAN9")
	if e == "" {
		return "", fmt.Errorf("plan9 path not found")
	}

	return fmt.Sprintf("SHELL=%s/%s", e, rc), nil
}
