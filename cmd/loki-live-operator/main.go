package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/go-kit/log"
	"github.com/go-kit/log/level"
	"github.com/prometheus/common/version"
	controller "sigs.k8s.io/controller-runtime"

	// Needed for clients.
	_ "k8s.io/client-go/plugin/pkg/client/auth"

	"loki-live-controller/pkg/operator"
)

func main() {
	var (
		logger = log.NewLogfmtLogger(log.NewSyncWriter(os.Stderr))
		cfg    = loadConfig(logger)

		err error
	)

	op, err := operator.New(logger, cfg)
	if err != nil {
		level.Error(logger).Log("msg", "unable to create operator", "err", err)
		os.Exit(1)
	}

	// Run the manager and wait for a signal to shut down.
	level.Info(logger).Log("msg", "starting manager")
	if err := op.Start(controller.SetupSignalHandler()); err != nil {
		level.Error(logger).Log("msg", "problem running manager", "err", err)
		os.Exit(1)
	}
}

// loadConfig will read command line flags and populate a Config. loadConfig
// will exit the program on failure.
func loadConfig(l log.Logger) *operator.Config {
	fs := flag.NewFlagSet(os.Args[0], flag.ExitOnError)

	var (
		printVersion bool
	)

	cfg, err := operator.NewConfig(fs)
	if err != nil {
		level.Error(l).Log("msg", "failed to parse flags", "err", err)
		os.Exit(1)
	}

	fs.BoolVar(&printVersion, "version", false, "Print this build's version information")

	if err := fs.Parse(os.Args[1:]); err != nil {
		level.Error(l).Log("msg", "failed to parse flags", "err", err)
		os.Exit(1)
	}

	if printVersion {
		fmt.Println(version.Print("agent-operator"))
		os.Exit(0)
	}

	return cfg
}
