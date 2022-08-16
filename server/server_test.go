package server

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestServer(t *testing.T) {
	s := server{
		cfg: DefaultConfig(),
	}

	assert.NotEqual(t, nil, s)
}
