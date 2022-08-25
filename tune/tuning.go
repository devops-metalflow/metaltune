package tune

import (
	"context"

	"github.com/devops-metalflow/metaltune/config"
)

type Tuning struct {
}

func (t *Tuning) Run(_ context.Context, _ *config.Config) error {
	return nil
}
