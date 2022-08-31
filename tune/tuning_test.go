package tune

import (
	"context"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
)

const (
	profile = "content"
)

func TestAddress(t *testing.T) {
	tn := Tuning{}
	ctx := context.Background()

	buf := tn.address(ctx)
	assert.NotEqual(t, "", buf)
}

func TestAuto(t *testing.T) {
	tn := Tuning{}
	ctx := context.Background()

	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		b, _ := json.Marshal(map[string]string{
			"address": "127.0.0.1",
			"msg":     "",
			"profile": profile,
			"suc":     "true",
		})
		_, _ = w.Write(b)
	}))
	defer ts.Close()

	tn.Url = ts.URL
	buf, err := tn.auto(ctx)

	assert.Equal(t, nil, err)
	assert.Equal(t, profile, buf)
}

func TestProfile(t *testing.T) {
	tn := Tuning{}
	ctx := context.Background()

	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		b, _ := json.Marshal(map[string]string{
			"msg": "",
			"suc": "true",
		})
		_, _ = w.Write(b)
	}))
	defer ts.Close()

	tn.Url = ts.URL
	err := tn.profile(ctx, profile)

	assert.Equal(t, nil, err)
}
