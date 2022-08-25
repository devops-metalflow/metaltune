package tune

import (
	"context"
	"os"

	"github.com/pkg/errors"
	"github.com/shirou/gopsutil/v3/host"

	"github.com/devops-metalflow/metaltune/config"
	"github.com/devops-metalflow/metaltune/pipe"
)

const (
	OS_DEBIAN = "debian"
)

var (
	envHome = os.Getenv("HOME")

	debianCmds = []string{
		"rm -rf " + envHome + "/.cache/thumbnails/*",
		"apt-get autoremove --purge",
		"apt-get clean",
		"journalctl --vacuum-time=1s",
		"rm -rf /var/lib/apt/lists/*",
		"rm -rf /var/log/*",
		"rm -rf /tmp/*",
		"sync",
		`echo "3" | tee /proc/sys/vm/drop_caches`,
	}
)

type Cleanup struct {
}

func (c *Cleanup) Run(ctx context.Context, cfg *config.Config) error {
	p, err := c.new(ctx, cfg)
	if err != nil {
		return errors.Wrap(err, "failed to new")
	}

	if err := p.Init(ctx); err != nil {
		return errors.Wrap(err, "failed to init")
	}

	defer func(p pipe.Pipe, ctx context.Context) {
		_ = p.Deinit(ctx)
	}(p, ctx)

	buf := c.fetch(ctx)
	if buf == nil || len(buf) == 0 {
		return errors.New("failed to fetch")
	}

	if err := p.Run(ctx, buf); err != nil {
		return errors.Wrap(err, "failed to run")
	}

	return nil
}

func (c *Cleanup) new(ctx context.Context, cfg *config.Config) (pipe.Pipe, error) {
	p := pipe.DefaultConfig()
	if p == nil {
		return nil, errors.New("failed to config")
	}

	p.Config = *cfg

	return pipe.New(ctx, p), nil
}

func (c *Cleanup) fetch(ctx context.Context) []string {
	var buf []string

	info, err := host.InfoWithContext(ctx)
	if err != nil {
		return nil
	}

	if info.PlatformFamily == OS_DEBIAN {
		buf = debianCmds
	}

	return buf
}
