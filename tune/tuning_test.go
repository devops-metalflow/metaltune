package tune

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestTuning(t *testing.T) {
	tn := Tuning{}
	assert.NotEqual(t, nil, tn)
}
