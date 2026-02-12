package sshx

import (
	"context"
	"time"

	"github.com/ashrafinamdar23/coe-sshx/core"
)

type Runner interface {
	Run(ctx context.Context, spec core.RunSpec) (core.Result, error)
}

// NoopRunner is a placeholder so downstream apps can compile against the module
// before real SSH support is implemented.
type NoopRunner struct{}

func NewNoop() Runner { return &NoopRunner{} }

func (r *NoopRunner) Run(ctx context.Context, spec core.RunSpec) (core.Result, error) {
	now := time.Now().UTC()
	return core.Result{
		ObservedAt: now,
		TotalTime:  0,
		Status:     "ok",
		Commands:   nil,
	}, nil
}
