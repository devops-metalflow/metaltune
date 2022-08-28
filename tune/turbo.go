package tune

import (
	"context"
	"os"
	"path/filepath"
	"strconv"
	"strings"

	"github.com/pkg/errors"
	"github.com/shirou/gopsutil/v3/host"

	"github.com/devops-metalflow/metaltune/config"
)

const (
	OS   = "linux"
	Perm = 0644
)

const (
	Base  = "/sys/devices/system/cpu"
	Range = "online"
	Len   = 2
)

const (
	PerfGov = "performance"
	ScalGov = "scaling_governor"
)

const (
	CurFreq  = "scaling_cur_freq"
	MaxFreq  = "scaling_max_freq"
	MinFreq  = "scaling_min_freq"
	SetSpeed = "scaling_setspeed"
)

const (
	IntBase = 10
	IntBit  = 64
)

type Turbo struct {
	base string
}

func (t *Turbo) Init(_ context.Context, _ *config.Config) error {
	t.base = Base
	return nil
}

func (t *Turbo) Deinit(_ context.Context) error {
	return nil
}

func (t *Turbo) Run(ctx context.Context) error {
	info, err := host.InfoWithContext(ctx)
	if err != nil {
		return errors.Wrap(err, "failed to info")
	}

	if info.OS != OS {
		return errors.New("OS not supported")
	}

	if err := t.setGovernor(ctx, PerfGov); err != nil {
		return errors.Wrap(err, "failed to set")
	}

	return nil
}

func (t *Turbo) getFreq(ctx context.Context) (map[int]int, error) {
	var freqs map[int]int

	buf, err := t.getVariable(ctx, CurFreq)
	if err != nil {
		return nil, errors.Wrap(err, "failed to get")
	}

	for key, val := range buf {
		b, err := strconv.ParseInt(val, IntBase, IntBit)
		if err != nil {
			return nil, errors.Wrap(err, "failed to parse")
		}
		freqs[key] = int(b)
	}

	return freqs, nil
}

func (t *Turbo) getMaxFreq(ctx context.Context) (map[int]int, error) {
	var freqs map[int]int

	cpus, err := t.getRange(ctx, filepath.Join(t.base, Range))
	if err != nil {
		return nil, errors.Wrap(err, "failed to get")
	}

	for _, item := range cpus {
		p := filepath.Join(t.base, "cpu"+strconv.Itoa(item), "cpufreq", MaxFreq)
		b, err := t.readFile(ctx, p)
		if err != nil {
			return nil, errors.Wrap(err, "failed to read")
		}
		b = strings.Trim(strings.Trim(b, "\n"), " ")
		f, err := strconv.ParseInt(strings.Split(b, " ")[0], IntBase, IntBit)
		if err != nil {
			return nil, errors.Wrap(err, "failed to parse")
		}
		freqs[item] = int(f)
	}

	return freqs, nil
}

func (t *Turbo) getMinFreq(ctx context.Context) (map[int]int, error) {
	var freqs map[int]int

	cpus, err := t.getRange(ctx, filepath.Join(t.base, Range))
	if err != nil {
		return nil, errors.Wrap(err, "failed to get")
	}

	for _, item := range cpus {
		p := filepath.Join(t.base, "cpu"+strconv.Itoa(item), "cpufreq", MinFreq)
		b, err := t.readFile(ctx, p)
		if err != nil {
			return nil, errors.Wrap(err, "failed to read")
		}
		b = strings.Trim(strings.Trim(b, "\n"), " ")
		f, err := strconv.ParseInt(strings.Split(b, " ")[0], IntBase, IntBit)
		if err != nil {
			return nil, errors.Wrap(err, "failed to parse")
		}
		freqs[item] = int(f)
	}

	return freqs, nil
}

func (t *Turbo) setFreq(ctx context.Context, freq int) error {
	max, err := t.getMaxFreq(ctx)
	if err != nil {
		return errors.Wrap(err, "failed to get")
	}

	min, err := t.getMinFreq(ctx)
	if err != nil {
		return errors.Wrap(err, "failed to get")
	}

	cpus, err := t.getRange(ctx, filepath.Join(t.base, Range))
	if err != nil {
		return errors.Wrap(err, "failed to get")
	}

	for _, item := range cpus {
		if freq < min[item] || freq > max[item] {
			return errors.Wrap(err, "invalid frequency")
		}
		p := filepath.Join(t.base, "cpu"+strconv.Itoa(item), "cpufreq", SetSpeed)
		if err := t.writeFile(ctx, p, strconv.Itoa(freq)); err != nil {
			return errors.Wrap(err, "failed to write")
		}
	}

	return nil
}

func (t *Turbo) getGovernor(ctx context.Context) (map[int]string, error) {
	return t.getVariable(ctx, ScalGov)
}

func (t *Turbo) setGovernor(ctx context.Context, name string) error {
	r, err := t.getRange(ctx, filepath.Join(t.base, Range))
	if err != nil {
		return errors.Wrap(err, "failed to get")
	}

	for _, item := range r {
		p := filepath.Join(t.base, "cpu"+strconv.Itoa(item), "cpufreq", ScalGov)
		if err := t.writeFile(ctx, p, name); err != nil {
			return errors.Wrap(err, "failed to write")
		}
	}

	return nil
}

func (t *Turbo) getOnline(ctx context.Context) ([]int, error) {
	return t.getRange(ctx, filepath.Join(t.base, Range))
}

func (t *Turbo) getVariable(ctx context.Context, name string) (map[int]string, error) {
	var freqs map[int]string

	r, err := t.getRange(ctx, filepath.Join(t.base, Range))
	if err != nil {
		return nil, errors.Wrap(err, "failed to get")
	}

	for _, item := range r {
		p := filepath.Join(t.base, "cpu"+strconv.Itoa(item), "cpufreq", name)
		b, err := t.readFile(ctx, p)
		if err != nil {
			return nil, errors.Wrap(err, "failed to read")
		}
		b = strings.Trim(strings.Trim(b, "\n"), " ")
		freqs[item] = strings.Split(b, " ")[0]
	}

	return freqs, nil
}

func (t *Turbo) getRange(ctx context.Context, name string) ([]int, error) {
	var cpus []int

	buf, err := t.readFile(ctx, name)
	if err != nil {
		return nil, errors.Wrap(err, "failed to read")
	}

	buf = strings.Trim(strings.Trim(buf, "\n"), " ")
	if buf == "" {
		return nil, errors.Wrap(err, "failed to trim")
	}

	for _, item := range strings.Split(buf, ",") {
		b := strings.Split(item, "-")
		if len(b) == Len {
			s, err := strconv.ParseInt(b[0], IntBase, IntBit)
			if err != nil {
				return nil, errors.Wrap(err, "failed to parse")
			}
			e, err := strconv.ParseInt(b[1], IntBase, IntBit)
			if err != nil {
				return nil, errors.Wrap(err, "failed to parse")
			}
			for i := int(s); i <= int(e); i++ {
				cpus = append(cpus, i)
			}
		} else {
			c, err := strconv.ParseInt(b[0], IntBase, IntBit)
			if err != nil {
				return nil, errors.Wrap(err, "failed to parse")
			}
			cpus = append(cpus, int(c))
		}
	}

	return cpus, nil
}

func (t *Turbo) readFile(_ context.Context, name string) (string, error) {
	buf, err := os.ReadFile(name)
	if err != nil {
		return "", errors.Wrap(err, "failed to read")
	}

	return string(buf), nil
}

func (t *Turbo) writeFile(_ context.Context, name, data string) error {
	return os.WriteFile(name, []byte(data), Perm)
}
