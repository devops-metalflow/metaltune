package tune

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestTune(t *testing.T) {
	tn := tune{
		cfg: DefaultConfig(),
	}

	assert.NotEqual(t, nil, tn)
}
