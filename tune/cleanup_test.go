package tune

import (
	"context"
	"testing"

	"github.com/shirou/gopsutil/v3/host"
	"github.com/stretchr/testify/assert"
)

func TestFetch(t *testing.T) {
	c := Cleanup{}
	ctx := context.Background()

	buf := c.fetch(ctx)

	info, _ := host.InfoWithContext(ctx)
	if info.PlatformFamily == OS_DEBIAN {
		assert.NotEqual(t, nil, buf)
		assert.NotEqual(t, 0, len(buf))
	} else {
		assert.Equal(t, 0, len(buf))
	}
}
