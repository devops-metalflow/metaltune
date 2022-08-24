package tune

import (
	"context"
	"testing"

	"github.com/shirou/gopsutil/v3/host"
	"github.com/stretchr/testify/assert"
)

func TestFetchCmds(t *testing.T) {
	c := Cleanup{}
	ctx := context.Background()

	buf := c.fetchCmds(ctx)
	info, _ := host.InfoWithContext(ctx)
	if info.PlatformFamily == OS_DEBIAN {
		assert.NotEqual(t, nil, buf)
		assert.NotEqual(t, 0, len(buf))
	} else {
		assert.Equal(t, 0, len(buf))
	}
}

func TestCheckPipe(t *testing.T) {
	c := Cleanup{}
	ctx := context.Background()

	info, _ := host.InfoWithContext(ctx)
	if info.PlatformFamily != OS_DEBIAN {
		return
	}

	buf := c.checkPipe(ctx, "ls")
	assert.Equal(t, 1, len(buf))

	buf = c.checkPipe(ctx, "ls | wc -l")
	assert.Equal(t, 2, len(buf))
}

func TestRunPipe(t *testing.T) {
	c := Cleanup{}
	ctx := context.Background()

	info, _ := host.InfoWithContext(ctx)
	if info.PlatformFamily != OS_DEBIAN {
		return
	}

	err := c.runPipe(ctx, "ls", "wc -l")
	assert.Equal(t, nil, err)
}

func TestRunCmd(t *testing.T) {
	c := Cleanup{}
	ctx := context.Background()

	info, _ := host.InfoWithContext(ctx)
	if info.PlatformFamily != OS_DEBIAN {
		return
	}

	err := c.runCmd(ctx, "")
	assert.NotEqual(t, nil, err)

	err = c.runCmd(ctx, "ls")
	assert.Equal(t, nil, err)
}
