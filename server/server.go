package server

import (
	"context"
	"log"
	"math"
	"net"

	"google.golang.org/grpc"

	"github.com/devops-metalflow/metaltune/config"
	pb "github.com/devops-metalflow/metaltune/server/proto"
)

const (
	KIND = "metaltune"
)

type Server interface {
	Init(context.Context) error
	Deinit(context.Context) error
	Run(context.Context) error
}

type Config struct {
	Addr   string
	Config config.Config
}

type server struct {
	cfg *Config
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

func (s *server) Init(_ context.Context) error {
	return nil
}

func (s *server) Deinit(_ context.Context) error {
	return nil
}

func (s *server) Run(_ context.Context) error {
	options := []grpc.ServerOption{grpc.MaxRecvMsgSize(math.MaxInt32), grpc.MaxSendMsgSize(math.MaxInt32)}

	g := grpc.NewServer(options...)
	pb.RegisterServerProtoServer(g, s)

	lis, _ := net.Listen("tcp", s.cfg.Addr)

	return g.Serve(lis)
}

func (s *server) SendServer(_ context.Context, in *pb.ServerRequest) (*pb.ServerReply, error) {
	log.Printf("Received: %v", in.GetSpec().GetTuning().GetAuto())
	return &pb.ServerReply{Output: "Hello " + in.GetSpec().GetTuning().GetProfile()}, nil
}
