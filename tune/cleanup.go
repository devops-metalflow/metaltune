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
	Platform = "debian"
)

var (
	Home = os.Getenv("HOME")
	Cmds = []string{
		"rm -rf " + Home + "/.cache/thumbnails/*",
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
	pipe pipe.Pipe
}

func (c *Cleanup) Init(ctx context.Context, cfg *config.Config) error {
	var err error

	c.pipe, err = c.new(ctx, cfg)
	if err != nil {
		return errors.Wrap(err, "failed to new")
	}

	if err := c.pipe.Init(ctx); err != nil {
		return errors.Wrap(err, "failed to init")
	}

	return nil
}

func (c *Cleanup) Deinit(ctx context.Context) error {
	return c.pipe.Deinit(ctx)
}

func (c *Cleanup) Run(ctx context.Context) error {
	buf := c.fetch(ctx)
	if buf == nil || len(buf) == 0 {
		return errors.New("failed to fetch")
	}

	if err := c.pipe.Run(ctx, buf); err != nil {
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

	if info.PlatformFamily == Platform {
		buf = Cmds
	}

	return buf
}
