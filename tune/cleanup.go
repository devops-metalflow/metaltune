package tune

import (
	"context"
	"io"
	"os"
	"os/exec"
	"strings"

	"github.com/pkg/errors"
	"github.com/shirou/gopsutil/v3/host"
)

const (
	OS_DEBIAN = "debian"
	PIPE_MAX  = 2
)

var (
	envHome = os.Getenv("HOME")

	debianCmds = []string{
		"rm -rf " + envHome + "/.cache/thumbnails/*",
		"apt-get autoremove --purge",
		"apt-get clean",
		"journalctl --vacuum-time=1s",
		"rm -rf /var/lib/apt/lists/*",
		"rm -rf /var/log/*",
		"rm -rf /tmp/*",
		"sync",
		`echo "3" | tee /proc/sys/vm/drop_caches`,
	}
)

type Cleanup struct {
}

func (c *Cleanup) Run(ctx context.Context) error {
	buf := c.fetchCmds(ctx)
	if buf == nil || len(buf) == 0 {
		return errors.New("OS not supported")
	}

	for _, item := range buf {
		p := c.checkPipe(ctx, item)
		if p == nil || len(p) == 0 || len(p) > 2 {
			return errors.New("pipe (>1) not supported")
		}
		if len(p) == 1 {
			if err := c.runCmd(ctx, p[0]); err != nil {
				return errors.Wrap(err, "failed to run command")
			}
		} else if len(p) == PIPE_MAX {
			if err := c.runPipe(ctx, p[0], p[1]); err != nil {
				return errors.Wrap(err, "failed to run pipe")
			}
		} else {
			// PASS
		}
	}

	return nil
}

func (c *Cleanup) fetchCmds(ctx context.Context) []string {
	var buf []string

	info, err := host.InfoWithContext(ctx)
	if err != nil {
		return nil
	}

	if info.PlatformFamily == OS_DEBIAN {
		buf = debianCmds
	}

	return buf
}

func (c *Cleanup) checkPipe(_ context.Context, cmd string) []string {
	var buf []string

	if strings.Contains(cmd, "|") {
		buf = strings.Split(cmd, "|")
	} else {
		buf = append(buf, cmd)
	}

	return buf
}

// nolint: gosec
func (c *Cleanup) runPipe(_ context.Context, cmd0, cmd1 string) error {
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

// nolint: gosec
func (c *Cleanup) runCmd(_ context.Context, cmd string) error {
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
