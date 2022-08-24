package tune

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestTuning(t *testing.T) {
	tn := tune{
		cfg: DefaultConfig(),
	}

	assert.NotEqual(t, nil, tn)
}
