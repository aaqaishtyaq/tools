/*
Copyright © 2022 Aaqa Ishtyaq aaqaishtyaq@gmail.com

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/
package main

import (
	"context"
	"os"
	"os/signal"
	"syscall"

	"github.com/aaqaishtyaq/tools/git-gh/cmd"
	"github.com/sirupsen/logrus"
)

func main() {
	exitCode := 0
	defer func() {
		os.Exit(exitCode)
	}()

	ctx := context.Background()

	ctx, cancel := context.WithCancel(ctx)

	signals := make(chan os.Signal, 1)

	signal.Notify(signals, os.Interrupt, syscall.SIGHUP, syscall.SIGTERM)

	defer func() {
		logrus.Info("Done")
		signal.Stop(signals)
		cancel()
	}()

	go func() {
		select {
		case <-signals:
			logrus.Info("Got signal, propagating...")
			cancel()
		case <-ctx.Done():
		}
	}()

	if err := cmd.ExecuteContext(ctx); err != nil {
		logrus.Error(err)
		exitCode = 1
		return
	}
}
