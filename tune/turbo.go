package tune

import (
	"context"

	"github.com/pkg/errors"
	"github.com/shirou/gopsutil/v3/host"

	"github.com/devops-metalflow/metaltune/config"
	"github.com/devops-metalflow/metaltune/pipe"
)

var (
	turboCmds = []string{}
)

type Turbo struct {
}

func (t *Turbo) Run(ctx context.Context, cfg *config.Config) error {
	p, err := t.new(ctx, cfg)
	if err != nil {
		return errors.Wrap(err, "failed to new")
	}

	if err := p.Init(ctx); err != nil {
		return errors.Wrap(err, "failed to init")
	}

	defer func(p pipe.Pipe, ctx context.Context) {
		_ = p.Deinit(ctx)
	}(p, ctx)

	buf := t.fetch(ctx)
	if buf == nil || len(buf) == 0 {
		return errors.New("failed to fetch")
	}

	if err := p.Run(ctx, buf); err != nil {
		return errors.Wrap(err, "failed to run")
	}

	return nil
}

func (t *Turbo) new(ctx context.Context, cfg *config.Config) (pipe.Pipe, error) {
	p := pipe.DefaultConfig()
	if p == nil {
		return nil, errors.New("failed to config")
	}

	p.Config = *cfg

	return pipe.New(ctx, p), nil
}

func (t *Turbo) fetch(ctx context.Context) []string {
	var buf []string

	info, err := host.InfoWithContext(ctx)
	if err != nil {
		return nil
	}

	if info.PlatformFamily == DEBIAN {
		buf = turboCmds
	}

	return buf
}
