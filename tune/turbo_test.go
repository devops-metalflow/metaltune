package tune

import (
	"context"
	"os"
	"path/filepath"
	"strconv"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

const (
	BaseTest = "/tmp/metaltune"
	SaveGov  = "powersave"
)

const (
	Count = 3
	Cur   = 500000
	Max   = 1000000
	Min   = 100000
)

// nolint: funlen
func initTurboTest() {
	_ = os.MkdirAll(BaseTest, os.ModePerm)

	// online
	file, _ := os.Create(filepath.Join(BaseTest, Range))
	_ = file.Close()

	buf := "0-2"
	_ = os.WriteFile(filepath.Join(BaseTest, Range), []byte(buf), Perm)

	// cpu0/cpufreq
	_ = os.MkdirAll(filepath.Join(BaseTest, "cpu0", "cpufreq"), os.ModePerm)

	// cpu0/cpufreq/scaling_governor
	file, _ = os.Create(filepath.Join(BaseTest, "cpu0", "cpufreq", ScalGov))
	_ = file.Close()

	buf = SaveGov
	_ = os.WriteFile(filepath.Join(BaseTest, "cpu0", "cpufreq", ScalGov), []byte(buf), Perm)

	// cpu0/cpufreq/scaling_cur_freq
	file, _ = os.Create(filepath.Join(BaseTest, "cpu0", "cpufreq", CurFreq))
	_ = file.Close()

	buf = strconv.Itoa(Cur)
	_ = os.WriteFile(filepath.Join(BaseTest, "cpu0", "cpufreq", CurFreq), []byte(buf), Perm)

	// cpu0/cpufreq/scaling_max_freq
	file, _ = os.Create(filepath.Join(BaseTest, "cpu0", "cpufreq", MaxFreq))
	_ = file.Close()

	buf = strconv.Itoa(Max)
	_ = os.WriteFile(filepath.Join(BaseTest, "cpu0", "cpufreq", MaxFreq), []byte(buf), Perm)

	// cpu0/cpufreq/scaling_min_freq
	file, _ = os.Create(filepath.Join(BaseTest, "cpu0", "cpufreq", MinFreq))
	_ = file.Close()

	buf = strconv.Itoa(Min)
	_ = os.WriteFile(filepath.Join(BaseTest, "cpu0", "cpufreq", MinFreq), []byte(buf), Perm)

	// cpu0/cpufreq/scaling_setspeed
	file, _ = os.Create(filepath.Join(BaseTest, "cpu0", "cpufreq", SetSpeed))
	_ = file.Close()

	buf = strconv.Itoa(Cur)
	_ = os.WriteFile(filepath.Join(BaseTest, "cpu0", "cpufreq", SetSpeed), []byte(buf), Perm)

	// cpu1/cpufreq
	_ = os.MkdirAll(filepath.Join(BaseTest, "cpu1", "cpufreq"), os.ModePerm)

	// cpu1/cpufreq/scaling_governor
	file, _ = os.Create(filepath.Join(BaseTest, "cpu1", "cpufreq", ScalGov))
	_ = file.Close()

	buf = SaveGov
	_ = os.WriteFile(filepath.Join(BaseTest, "cpu1", "cpufreq", ScalGov), []byte(buf), Perm)

	// cpu1/cpufreq/scaling_cur_freq
	file, _ = os.Create(filepath.Join(BaseTest, "cpu1", "cpufreq", CurFreq))
	_ = file.Close()

	buf = strconv.Itoa(Cur)
	_ = os.WriteFile(filepath.Join(BaseTest, "cpu1", "cpufreq", CurFreq), []byte(buf), Perm)

	// cpu1/cpufreq/scaling_max_freq
	file, _ = os.Create(filepath.Join(BaseTest, "cpu1", "cpufreq", MaxFreq))
	_ = file.Close()

	buf = strconv.Itoa(Max)
	_ = os.WriteFile(filepath.Join(BaseTest, "cpu1", "cpufreq", MaxFreq), []byte(buf), Perm)

	// cpu1/cpufreq/scaling_min_freq
	file, _ = os.Create(filepath.Join(BaseTest, "cpu1", "cpufreq", MinFreq))
	_ = file.Close()

	buf = strconv.Itoa(Min)
	_ = os.WriteFile(filepath.Join(BaseTest, "cpu1", "cpufreq", MinFreq), []byte(buf), Perm)

	// cpu1/cpufreq/scaling_setspeed
	file, _ = os.Create(filepath.Join(BaseTest, "cpu1", "cpufreq", SetSpeed))
	_ = file.Close()

	buf = strconv.Itoa(Cur)
	_ = os.WriteFile(filepath.Join(BaseTest, "cpu1", "cpufreq", SetSpeed), []byte(buf), Perm)

	// cpu2/cpufreq
	_ = os.MkdirAll(filepath.Join(BaseTest, "cpu2", "cpufreq"), os.ModePerm)

	// cpu2/cpufreq/scaling_governor
	file, _ = os.Create(filepath.Join(BaseTest, "cpu2", "cpufreq", ScalGov))
	_ = file.Close()

	buf = SaveGov
	_ = os.WriteFile(filepath.Join(BaseTest, "cpu2", "cpufreq", ScalGov), []byte(buf), Perm)

	// cpu2/cpufreq/scaling_cur_freq
	file, _ = os.Create(filepath.Join(BaseTest, "cpu2", "cpufreq", CurFreq))
	_ = file.Close()

	buf = strconv.Itoa(Cur)
	_ = os.WriteFile(filepath.Join(BaseTest, "cpu2", "cpufreq", CurFreq), []byte(buf), Perm)

	// cpu2/cpufreq/scaling_max_freq
	file, _ = os.Create(filepath.Join(BaseTest, "cpu2", "cpufreq", MaxFreq))
	_ = file.Close()

	buf = strconv.Itoa(Max)
	_ = os.WriteFile(filepath.Join(BaseTest, "cpu2", "cpufreq", MaxFreq), []byte(buf), Perm)

	// cpu2/cpufreq/scaling_min_freq
	file, _ = os.Create(filepath.Join(BaseTest, "cpu2", "cpufreq", MinFreq))
	_ = file.Close()

	buf = strconv.Itoa(Min)
	_ = os.WriteFile(filepath.Join(BaseTest, "cpu2", "cpufreq", MinFreq), []byte(buf), Perm)

	// cpu2/cpufreq/scaling_setspeed
	file, _ = os.Create(filepath.Join(BaseTest, "cpu2", "cpufreq", SetSpeed))
	_ = file.Close()

	buf = strconv.Itoa(Cur)
	_ = os.WriteFile(filepath.Join(BaseTest, "cpu2", "cpufreq", SetSpeed), []byte(buf), Perm)
}

func deinitTurboTest() {
	_ = os.RemoveAll(BaseTest)
}

func TestGetFreq(t *testing.T) {
	initTurboTest()

	tb := Turbo{base: BaseTest}
	ctx := context.Background()

	buf, err := tb.getFreq(ctx)
	assert.Equal(t, nil, err)
	assert.NotEqual(t, nil, buf)
	assert.Equal(t, Count, len(buf))

	for _, val := range buf {
		assert.Equal(t, Cur, val)
	}

	deinitTurboTest()
}

func TestMaxFreq(t *testing.T) {
	initTurboTest()

	tb := Turbo{base: BaseTest}
	ctx := context.Background()

	buf, err := tb.getMaxFreq(ctx)
	assert.Equal(t, nil, err)
	assert.NotEqual(t, nil, buf)
	assert.Equal(t, Count, len(buf))

	for _, val := range buf {
		assert.Equal(t, Max, val)
	}

	deinitTurboTest()
}

func TestMinFreq(t *testing.T) {
	initTurboTest()

	tb := Turbo{base: BaseTest}
	ctx := context.Background()

	buf, err := tb.getMinFreq(ctx)
	assert.Equal(t, nil, err)
	assert.NotEqual(t, nil, buf)
	assert.Equal(t, Count, len(buf))

	for _, val := range buf {
		assert.Equal(t, Min, val)
	}

	deinitTurboTest()
}

func TestSetFreq(t *testing.T) {
	initTurboTest()

	tb := Turbo{base: BaseTest}
	ctx := context.Background()

	err := tb.setFreq(ctx, Min-1)
	assert.NotEqual(t, nil, err)

	err = tb.setFreq(ctx, Max+1)
	assert.NotEqual(t, nil, err)

	err = tb.setFreq(ctx, Max)
	assert.Equal(t, nil, err)

	deinitTurboTest()
}

func TestGetGovernor(t *testing.T) {
	initTurboTest()

	tb := Turbo{base: BaseTest}
	ctx := context.Background()

	buf, err := tb.getGovernor(ctx)
	assert.Equal(t, nil, err)
	assert.NotEqual(t, nil, buf)
	assert.Equal(t, Count, len(buf))

	for _, val := range buf {
		assert.Equal(t, SaveGov, val)
	}

	deinitTurboTest()
}

func TestSetGovernor(t *testing.T) {
	initTurboTest()

	tb := Turbo{base: BaseTest}
	ctx := context.Background()

	err := tb.setGovernor(ctx, PerfGov)
	assert.Equal(t, nil, err)

	buf, err := tb.getGovernor(ctx)
	assert.Equal(t, nil, err)
	assert.NotEqual(t, nil, buf)
	assert.Equal(t, Count, len(buf))

	for _, val := range buf {
		assert.Equal(t, PerfGov, val)
	}

	deinitTurboTest()
}

func TestGetOnline(t *testing.T) {
	initTurboTest()

	tb := Turbo{base: BaseTest}
	ctx := context.Background()

	buf, err := tb.getOnline(ctx)
	assert.Equal(t, nil, err)
	assert.NotEqual(t, nil, buf)
	assert.Equal(t, Count, len(buf))

	deinitTurboTest()
}

func TestGetVariable(t *testing.T) {
	initTurboTest()

	tb := Turbo{base: BaseTest}
	ctx := context.Background()

	buf, err := tb.getVariable(ctx, CurFreq)
	assert.Equal(t, nil, err)
	assert.NotEqual(t, nil, buf)
	assert.Equal(t, Count, len(buf))

	deinitTurboTest()
}

func TestGetRange(t *testing.T) {
	initTurboTest()

	tb := Turbo{base: BaseTest}
	ctx := context.Background()

	buf, err := tb.getRange(ctx, filepath.Join(tb.base, Range))
	assert.Equal(t, nil, err)
	assert.NotEqual(t, nil, buf)
	assert.Equal(t, Count, len(buf))

	deinitTurboTest()
}

func TestReadFile(t *testing.T) {
	initTurboTest()

	tb := Turbo{base: BaseTest}
	ctx := context.Background()

	p := filepath.Join(tb.base, "cpu0", "cpufreq", SetSpeed)
	buf, err := tb.readFile(ctx, p)
	assert.Equal(t, nil, err)

	buf = strings.Trim(strings.Trim(buf, "\n"), " ")
	b, err := strconv.ParseInt(strings.Split(buf, " ")[0], IntBase, IntBit)
	assert.Equal(t, nil, err)
	assert.Equal(t, Cur, int(b))

	deinitTurboTest()
}

func TestWriteFile(t *testing.T) {
	initTurboTest()

	tb := Turbo{base: BaseTest}
	ctx := context.Background()

	p := filepath.Join(tb.base, "cpu0", "cpufreq", SetSpeed)
	err := tb.writeFile(ctx, p, strconv.Itoa(Max))
	assert.Equal(t, nil, err)

	buf, err := tb.readFile(ctx, p)
	assert.Equal(t, nil, err)

	buf = strings.Trim(strings.Trim(buf, "\n"), " ")
	b, err := strconv.ParseInt(strings.Split(buf, " ")[0], IntBase, IntBit)
	assert.Equal(t, nil, err)
	assert.Equal(t, Max, int(b))

	deinitTurboTest()
}
