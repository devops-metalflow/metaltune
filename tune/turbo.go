package tune

import (
	"context"

	"github.com/devops-metalflow/metaltune/config"
)

type Turbo struct {
}

func (t *Turbo) Run(_ context.Context, _ *config.Config) error {
	return nil
}
