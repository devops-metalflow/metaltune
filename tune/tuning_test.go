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
	base64Decode = "abc123!?$*&()'-=@~"
	base64Encode = "YWJjMTIzIT8kKiYoKSctPUB+"
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
			"profile": base64Encode,
		})
		_, _ = w.Write(b)
	}))
	defer ts.Close()

	tn.Url = ts.URL
	buf, err := tn.auto(ctx)

	assert.Equal(t, nil, err)
	assert.Equal(t, base64Decode, buf)
}

func TestProfile(t *testing.T) {
	tn := Tuning{}
	ctx := context.Background()

	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
	}))
	defer ts.Close()

	tn.Url = ts.URL
	err := tn.profile(ctx, base64Decode)

	assert.Equal(t, nil, err)
}

func TestDecode(t *testing.T) {
	tn := Tuning{}
	ctx := context.Background()

	buf := tn.decode(ctx, base64Encode)
	assert.Equal(t, base64Decode, buf)
}

func TestEncode(t *testing.T) {
	tn := Tuning{}
	ctx := context.Background()

	buf := tn.encode(ctx, base64Decode)
	assert.Equal(t, base64Encode, buf)
}
