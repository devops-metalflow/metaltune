package server

import (
	"context"
	"math"
	"net"

	"github.com/pkg/errors"
	"google.golang.org/grpc"

	"github.com/devops-metalflow/metaltune/config"
	pb "github.com/devops-metalflow/metaltune/server/proto"
	"github.com/devops-metalflow/metaltune/tune"
)

const (
	Kind = "metaltune"
)

type Server interface {
	Init(context.Context) error
	Deinit(context.Context) error
	Run(context.Context) error
}

type Config struct {
	Address string
	Config  config.Config
	Tune    tune.Tune
}

type server struct {
	cfg *Config
	srv *grpc.Server
	pb.UnimplementedServerProtoServer
}

func New(_ context.Context, cfg *Config) Server {
	return &server{
		cfg: cfg,
	}
}

func DefaultConfig() *Config {
	return &Config{}
}

func (s *server) Init(ctx context.Context) error {
	if err := s.cfg.Tune.Init(ctx); err != nil {
		return errors.Wrap(err, "failed to init tune")
	}

	options := []grpc.ServerOption{grpc.MaxRecvMsgSize(math.MaxInt32), grpc.MaxSendMsgSize(math.MaxInt32)}

	s.srv = grpc.NewServer(options...)
	pb.RegisterServerProtoServer(s.srv, s)

	return nil
}

func (s *server) Deinit(ctx context.Context) error {
	s.srv.Stop()
	_ = s.cfg.Tune.Deinit(ctx)

	return nil
}

func (s *server) Run(_ context.Context) error {
	lis, _ := net.Listen("tcp", s.cfg.Address)
	return s.srv.Serve(lis)
}

func (s *server) SendServer(ctx context.Context, in *pb.ServerRequest) (*pb.ServerReply, error) {
	if in.GetKind() != Kind {
		return &pb.ServerReply{Error: "invalid kind"}, nil
	}

	if in.GetSpec().GetCleanup() {
		if err := s.cfg.Tune.Cleanup(ctx); err != nil {
			return &pb.ServerReply{Error: "failed to cleanup"}, nil
		}
	}

	auto := in.GetSpec().GetTuning().GetAuto()
	profile := in.GetSpec().GetTuning().GetProfile()

	if auto || profile != "" {
		if err := s.cfg.Tune.Tuning(ctx, auto, profile); err != nil {
			return &pb.ServerReply{Error: "failed to tuning"}, nil
		}
	}

	if in.GetSpec().GetTurbo() {
		if err := s.cfg.Tune.Turbo(ctx); err != nil {
			return &pb.ServerReply{Error: "failed to turbo"}, nil
		}
	}

	return &pb.ServerReply{Output: "completed"}, nil
}
