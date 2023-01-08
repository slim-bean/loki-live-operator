package reconciler

import (
	"context"

	"sigs.k8s.io/controller-runtime/pkg/reconcile"
)

type Reconciler struct {
}

func (r *Reconciler) Reconcile(ctx context.Context, req reconcile.Request) (reconcile.Result, error) {

	return reconcile.Result{}, nil
}
