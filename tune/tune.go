package tune

import (
	"context"

	"github.com/pkg/errors"

	"github.com/devops-metalflow/metaltune/config"
)

type Tune interface {
	Init(context.Context) error
	Deinit(context.Context) error
	Cleanup(context.Context) error
	Tuning(context.Context) error
	Turbo(context.Context) error
}

type Config struct {
	Config config.Config
}

type tune struct {
	cfg *Config
}

func New(_ context.Context, cfg *Config) Tune {
	return &tune{
		cfg: cfg,
	}
}

func DefaultConfig() *Config {
	return &Config{}
}

func (t *tune) Init(_ context.Context) error {
	return nil
}

func (t *tune) Deinit(_ context.Context) error {
	return nil
}

func (t *tune) Cleanup(ctx context.Context) error {
	r := Cleanup{}

	if err := r.Init(ctx, &t.cfg.Config); err != nil {
		return errors.Wrap(err, "failed to init")
	}

	defer func(r *Cleanup, ctx context.Context) {
		_ = r.Deinit(ctx)
	}(&r, ctx)

	return r.Run(ctx)
}

func (t *tune) Tuning(ctx context.Context) error {
	tn := Tuning{}

	if err := tn.Init(ctx, &t.cfg.Config); err != nil {
		return errors.Wrap(err, "failed to init")
	}

	defer func(tn *Tuning, ctx context.Context) {
		_ = tn.Deinit(ctx)
	}(&tn, ctx)

	return tn.Run(ctx)
}

func (t *tune) Turbo(ctx context.Context) error {
	tb := Turbo{}

	if err := tb.Init(ctx, &t.cfg.Config); err != nil {
		return errors.Wrap(err, "failed to init")
	}

	defer func(tb *Turbo, ctx context.Context) {
		_ = tb.Deinit(ctx)
	}(&tb, ctx)

	return tb.Run(ctx)
}
