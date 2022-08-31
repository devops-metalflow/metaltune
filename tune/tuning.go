package tune

import (
	"bytes"
	"context"
	"encoding/json"
	"io"
	"net"
	"net/http"
	"strings"

	"github.com/pkg/errors"

	"github.com/devops-metalflow/metaltune/config"
)

const (
	statusOK = "true"
)

type Tuning struct {
	Address string
	Pass    string
	Url     string
	User    string
}

func (t *Tuning) Init(ctx context.Context, cfg *config.Config) error {
	t.Address = t.address(ctx)
	t.Pass = cfg.Spec.Tuning.Pass
	t.Url = cfg.Spec.Tuning.Url
	t.User = cfg.Spec.Tuning.User

	return nil
}

func (t *Tuning) Deinit(_ context.Context) error {
	return nil
}

func (t *Tuning) Run(ctx context.Context, auto bool, profile string) (string, error) {
	var buf string
	var err error

	if auto {
		buf, err = t.auto(ctx)
		if err != nil {
			return "", errors.Wrap(err, "failed to runAuto")
		}
	}

	if profile != "" {
		if err = t.profile(ctx, profile); err != nil {
			return "", errors.Wrap(err, "failed to runProfile")
		}
		buf = ""
	}

	return buf, nil
}

func (t *Tuning) address(_ context.Context) string {
	conn, _ := net.Dial("udp", "8.8.8.8:8")
	defer func(conn net.Conn) {
		_ = conn.Close()
	}(conn)

	buf := conn.LocalAddr().(*net.UDPAddr)
	addr := strings.Split(buf.String(), ":")[0]

	return addr
}

func (t *Tuning) auto(_ context.Context) (string, error) {
	buf := map[string]string{
		"address": t.Address,
	}

	body, err := json.Marshal(buf)
	if err != nil {
		return "", errors.Wrap(err, "failed to marshal")
	}

	req, err := http.NewRequest(http.MethodPost, t.Url+"/tuning/auto", bytes.NewBuffer(body))
	if err != nil {
		return "", errors.Wrap(err, "failed to request")
	}

	req.Header.Add("Content-Type", "application/json")

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return "", errors.Wrap(err, "failed to post")
	}

	defer func() {
		_ = resp.Body.Close()
	}()

	if resp.StatusCode != http.StatusOK {
		return "", errors.New("invalid status")
	}

	ret, err := io.ReadAll(resp.Body)
	if err != nil {
		return "", errors.Wrap(err, "failed to read")
	}

	data := make(map[string]interface{})
	if err := json.Unmarshal(ret, &data); err != nil {
		return "", errors.Wrap(err, "failed to unmarshal")
	}

	if data["suc"].(string) != statusOK {
		return "", errors.New(data["msg"].(string))
	}

	return data["profile"].(string), nil
}

func (t *Tuning) profile(_ context.Context, profile string) error {
	buf := map[string]string{
		"address": t.Address,
		"profile": profile,
	}

	body, err := json.Marshal(buf)
	if err != nil {
		return errors.Wrap(err, "failed to marshal")
	}

	req, err := http.NewRequest(http.MethodPost, t.Url+"/tuning/profile", bytes.NewBuffer(body))
	if err != nil {
		return errors.Wrap(err, "failed to request")
	}

	req.Header.Add("Content-Type", "application/json")

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return errors.Wrap(err, "failed to post")
	}

	defer func() {
		_ = resp.Body.Close()
	}()

	if resp.StatusCode != http.StatusOK {
		return errors.New("invalid status")
	}

	ret, err := io.ReadAll(resp.Body)
	if err != nil {
		return errors.Wrap(err, "failed to read")
	}

	data := make(map[string]interface{})
	if err := json.Unmarshal(ret, &data); err != nil {
		return errors.Wrap(err, "failed to unmarshal")
	}

	if data["suc"].(string) != statusOK {
		return errors.New(data["msg"].(string))
	}

	return nil
}
