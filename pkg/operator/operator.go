package operator

import (
	"context"
	"fmt"

	"github.com/go-kit/log"
	"github.com/go-kit/log/level"
	apps_v1 "k8s.io/api/apps/v1"
	"k8s.io/client-go/rest"
	controller "sigs.k8s.io/controller-runtime"
	"sigs.k8s.io/controller-runtime/pkg/healthz"

	"loki-live-controller/pkg/reconciler"
)

// Operator is the Loki Live operator
type Operator struct {
	logger log.Logger
}

type Config struct {
	Controller controller.Options
	// RestConfig used to connect to cluster. One will be generated based on the
	// environment if not set.
	RestConfig *rest.Config
}

func New(l log.Logger, c *Config) (*Operator, error) {

	restConfig := c.RestConfig
	if restConfig == nil {
		restConfig = controller.GetConfigOrDie()
	}
	manager, err := controller.NewManager(restConfig, c.Controller)
	if err != nil {
		return nil, fmt.Errorf("failed to create manager: %w", err)
	}

	if err := manager.AddReadyzCheck("running", healthz.Ping); err != nil {
		level.Warn(l).Log("msg", "failed to set up 'running' readyz check", "err", err)
	}
	if err := manager.AddHealthzCheck("running", healthz.Ping); err != nil {
		level.Warn(l).Log("msg", "failed to set up 'running' healthz check", "err", err)
	}

	r := reconciler.Reconciler{}

	err = controller.NewControllerManagedBy(manager).
		Owns(&apps_v1.DaemonSet{}).
		Complete(&r)
	if err != nil {
		return nil, fmt.Errorf("failed to create GrafanaAgent controller: %w", err)
	}

	return &Operator{
		logger: l,
	}, nil
}

func (o *Operator) Start(ctx context.Context) error {
	return nil
}
