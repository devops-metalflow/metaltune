package tune

import (
	"context"

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
	return r.Run(ctx, &t.cfg.Config)
}

func (t *tune) Tuning(ctx context.Context) error {
	tn := Tuning{}
	return tn.Run(ctx, &t.cfg.Config)
}

func (t *tune) Turbo(ctx context.Context) error {
	tb := Turbo{}
	return tb.Run(ctx, &t.cfg.Config)
}
