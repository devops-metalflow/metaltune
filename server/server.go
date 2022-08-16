package server

import (
	"context"
	"log"
	"math"
	"net"

	"google.golang.org/grpc"

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
	Addr string
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

func (s *server) Init(ctx context.Context) error {
	return nil
}

func (s *server) Deinit(ctx context.Context) error {
	return nil
}

func (s *server) Run(_ context.Context) error {
	options := []grpc.ServerOption{grpc.MaxRecvMsgSize(math.MaxInt32), grpc.MaxSendMsgSize(math.MaxInt32)}

	g := grpc.NewServer(options...)
	pb.RegisterServerProtoServer(g, s)

	lis, _ := net.Listen("tcp", s.cfg.Addr)

	return g.Serve(lis)
}

func (s *server) SendServer(ctx context.Context, in *pb.ServerRequest) (*pb.ServerReply, error) {
	log.Printf("Received: %v", in.GetSpec().GetName())
	return &pb.ServerReply{Output: "Hello " + in.GetSpec().GetName()}, nil
}
