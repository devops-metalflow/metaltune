package pipe

import (
	"context"
	"io"
	"os"
	"os/exec"
	"strings"

	"github.com/pkg/errors"

	"github.com/devops-metalflow/metaltune/config"
)

const (
	PIPE_MAX = 2
)

type Pipe interface {
	Init(context.Context) error
	Deinit(context.Context) error
	Run(context.Context, []string) error
}

type Config struct {
	Config config.Config
}

type pipe struct {
	cfg *Config
}

func New(_ context.Context, cfg *Config) Pipe {
	return &pipe{
		cfg: cfg,
	}
}

func DefaultConfig() *Config {
	return &Config{}
}

func (p *pipe) Init(_ context.Context) error {
	return nil
}

func (p *pipe) Deinit(_ context.Context) error {
	return nil
}

func (p *pipe) Run(ctx context.Context, cmds []string) error {
	for _, item := range cmds {
		b := p.check(ctx, item)
		if b == nil || len(b) == 0 || len(b) > 2 {
			return errors.New("pipe (>1) not supported")
		}
		if len(b) == 1 {
			if err := p.runCmd(ctx, b[0]); err != nil {
				return errors.Wrap(err, "failed to run command")
			}
		} else if len(b) == PIPE_MAX {
			if err := p.runPipe(ctx, b[0], b[1]); err != nil {
				return errors.Wrap(err, "failed to run pipe")
			}
		} else {
			// PASS
		}
	}

	return nil
}

func (p *pipe) check(_ context.Context, cmd string) []string {
	var buf []string

	if strings.Contains(cmd, "|") {
		buf = strings.Split(cmd, "|")
	} else {
		buf = append(buf, cmd)
	}

	return buf
}

// nolint: gosec
func (p *pipe) runCmd(_ context.Context, cmd string) error {
	var err error

	buf := strings.Split(cmd, " ")
	if len(buf) == 0 {
		return errors.New("invalid command")
	}

	if len(buf) == 1 {
		_, err = exec.Command(buf[0]).Output()
	} else {
		_, err = exec.Command(buf[0], buf[1:]...).Output()
	}

	return err
}

// nolint: gosec
func (p *pipe) runPipe(_ context.Context, cmd0, cmd1 string) error {
	helper := func(cmd string) *exec.Cmd {
		buf := strings.Split(cmd, " ")
		if len(buf) == 0 {
			return nil
		}
		if len(buf) == 1 {
			return exec.Command(buf[0])
		}
		return exec.Command(buf[0], buf[1:]...)
	}

	c0 := helper(cmd0)
	c1 := helper(cmd1)

	reader, writer := io.Pipe()

	defer func(writer *io.PipeWriter) {
		_ = writer.Close()
	}(writer)

	defer func(reader *io.PipeReader) {
		_ = reader.Close()
	}(reader)

	c0.Stdout = writer
	c1.Stdin = reader
	c1.Stdout = os.Stdout

	if err := c0.Start(); err != nil {
		return errors.Wrap(err, "failed to start cmd0")
	}

	if err := c1.Start(); err != nil {
		return errors.Wrap(err, "failed to start cmd1")
	}

	go func() {
		defer func(writer *io.PipeWriter) {
			_ = writer.Close()
		}(writer)
		_ = c0.Wait()
	}()

	if err := c1.Wait(); err != nil {
		return errors.Wrap(err, "failed to wait cmd1")
	}

	return nil
}
