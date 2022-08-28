package tune

import (
	"context"

	"github.com/devops-metalflow/metaltune/config"
)

type Tuning struct {
}

func (t *Tuning) Init(_ context.Context, _ *config.Config) error {
	return nil
}

func (t *Tuning) Deinit(_ context.Context) error {
	return nil
}

func (t *Tuning) Run(_ context.Context) error {
	return nil
}
