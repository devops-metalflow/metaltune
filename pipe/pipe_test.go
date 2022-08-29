package pipe

import (
	"context"
	"testing"

	"github.com/shirou/gopsutil/v3/host"
	"github.com/stretchr/testify/assert"
)

const (
	Debian = "debian"
)

func TestCheck(t *testing.T) {
	p := pipe{
		cfg: DefaultConfig(),
	}

	ctx := context.Background()

	info, _ := host.InfoWithContext(ctx)
	if info.PlatformFamily != Debian {
		return
	}

	buf := p.check(ctx, "ls")
	assert.Equal(t, 1, len(buf))

	buf = p.check(ctx, "ls | wc -l")
	assert.Equal(t, 2, len(buf))
}

func TestRunCmd(t *testing.T) {
	p := pipe{
		cfg: DefaultConfig(),
	}

	ctx := context.Background()

	info, _ := host.InfoWithContext(ctx)
	if info.PlatformFamily != Debian {
		return
	}

	err := p.runCmd(ctx, "")
	assert.NotEqual(t, nil, err)

	err = p.runCmd(ctx, "ls")
	assert.Equal(t, nil, err)
}

func TestRunPipe(t *testing.T) {
	p := pipe{
		cfg: DefaultConfig(),
	}

	ctx := context.Background()

	info, _ := host.InfoWithContext(ctx)
	if info.PlatformFamily != Debian {
		return
	}

	err := p.runPipe(ctx, "ls", "wc -l")
	assert.Equal(t, nil, err)
}
