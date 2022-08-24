package tune

import (
	"context"
	"testing"

	"github.com/shirou/gopsutil/v3/host"
	"github.com/stretchr/testify/assert"
)

func TestCleanup(t *testing.T) {
	tn := tune{
		cfg: DefaultConfig(),
	}

	info, _ := host.InfoWithContext(context.Background())
	if info.PlatformFamily != OS_DEBIAN {
		return
	}

	err := tn.Cleanup(context.Background())
	assert.Equal(t, nil, err)
}

func TestTuning(t *testing.T) {
	tn := tune{
		cfg: DefaultConfig(),
	}

	info, _ := host.InfoWithContext(context.Background())
	if info.PlatformFamily != OS_DEBIAN {
		return
	}

	err := tn.Tuning(context.Background())
	assert.Equal(t, nil, err)
}

func TestTurbo(t *testing.T) {
	tn := tune{
		cfg: DefaultConfig(),
	}

	info, _ := host.InfoWithContext(context.Background())
	if info.PlatformFamily != OS_DEBIAN {
		return
	}

	err := tn.Turbo(context.Background())
	assert.Equal(t, nil, err)
}
