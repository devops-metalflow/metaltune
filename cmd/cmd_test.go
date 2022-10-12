package cmd

import (
	"context"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestInitConfig(t *testing.T) {
	ctx := context.Background()

	_, err := initConfig(ctx, "invalid.yml")
	assert.NotEqual(t, nil, err)

	_, err = initConfig(ctx, "../test/invalid.yml")
	assert.NotEqual(t, nil, err)

	_, err = initConfig(ctx, "../test/config.yml")
	assert.Equal(t, nil, err)
}

func TestInitTune(t *testing.T) {
	ctx := context.Background()

	c, err := initConfig(ctx, "../test/config.yml")
	assert.Equal(t, nil, err)

	_, err = initTune(ctx, c)
	assert.Equal(t, nil, err)
}

func TestInitServer(t *testing.T) {
	ctx := context.Background()

	c, err := initConfig(ctx, "../test/config.yml")
	assert.Equal(t, nil, err)

	tn, err := initTune(ctx, c)
	assert.Equal(t, nil, err)

	_, err = initServer(ctx, c, tn)
	assert.Equal(t, nil, err)
}
