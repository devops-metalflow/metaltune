package cmd

import (
	"context"
	"io"
	"log"
	"os"
	"os/signal"
	"syscall"

	"github.com/pkg/errors"
	"gopkg.in/alecthomas/kingpin.v2"
	"gopkg.in/yaml.v3"

	"github.com/devops-metalflow/metaltune/config"
	"github.com/devops-metalflow/metaltune/server"
	"github.com/devops-metalflow/metaltune/tune"
)

var (
	app        = kingpin.New("metaltune", "metaltune").Version(config.Version + "-build-" + config.Build)
	configFile = app.Flag("config-file", "Config file (.yml)").Required().String()
	listenUrl  = app.Flag("listen-url", "Listen URL (host:port)").Required().String()
)

func Run(ctx context.Context) error {
	kingpin.MustParse(app.Parse(os.Args[1:]))

	c, err := initConfig(ctx, *configFile)
	if err != nil {
		return errors.Wrap(err, "failed to init config")
	}

	t, err := initTune(ctx, c)
	if err != nil {
		return errors.Wrap(err, "failed to init tune")
	}

	s, err := initServer(ctx, c, t)
	if err != nil {
		return errors.Wrap(err, "failed to init server")
	}

	if err := runServer(ctx, t, s); err != nil {
		return errors.Wrap(err, "failed to run server")
	}

	return nil
}

func initConfig(_ context.Context, name string) (*config.Config, error) {
	c := config.New()

	fi, err := os.Open(name)
	if err != nil {
		return c, errors.Wrap(err, "failed to open")
	}

	defer func() {
		_ = fi.Close()
	}()

	buf, _ := io.ReadAll(fi)

	if err := yaml.Unmarshal(buf, c); err != nil {
		return c, errors.Wrap(err, "failed to unmarshal")
	}

	return c, nil
}

func initTune(ctx context.Context, cfg *config.Config) (tune.Tune, error) {
	c := tune.DefaultConfig()
	if c == nil {
		return nil, errors.New("failed to config")
	}

	c.Config = *cfg

	return tune.New(ctx, c), nil
}

func initServer(ctx context.Context, cfg *config.Config, tn tune.Tune) (server.Server, error) {
	c := server.DefaultConfig()
	if c == nil {
		return nil, errors.New("failed to config")
	}

	c.Address = *listenUrl
	c.Config = *cfg
	c.Tune = tn

	return server.New(ctx, c), nil
}

func runServer(ctx context.Context, tn tune.Tune, srv server.Server) error {
	if err := srv.Init(ctx); err != nil {
		return errors.New("failed to init")
	}

	go func() {
		if err := srv.Run(ctx); err != nil {
			log.Fatalf("failed to run: %v", err)
		}
	}()

	s := make(chan os.Signal, 1)

	// kill (no param) default send syscanll.SIGTERM
	// kill -2 is syscall.SIGINT
	// kill -9 is syscall.SIGKILL but can"t be caught, so don't need add it
	signal.Notify(s, syscall.SIGINT, syscall.SIGTERM)

	done := make(chan bool, 1)

	go func() {
		<-s
		_ = srv.Deinit(ctx)
		done <- true
	}()

	<-done

	return nil
}
