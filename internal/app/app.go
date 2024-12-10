package app

import (
	"context"
	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/grpc/reflection"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"job-service/internal/data"
	"job-service/internal/loger"
	"job-service/internal/proto/pb"
	"job-service/internal/server"
	"net"
	"net/http"
)

type App struct {
	DB     data.DataBase
	Logger loger.Logger
	gRPC   *grpc.Server
	server *server.Server
	Lis    net.Listener
}

func Init() App {
	logger := loger.New()
	lis, err := net.Listen("tcp", "0.0.0.0:50051")
	if err != nil {
		logger.Error().Err(err)
	}
	db, err := gorm.Open(sqlite.Open("storage/main.db"))
	if err != nil {
		logger.Error().Err(err)
	}

	return App{
		DB:     data.DataBase{db},
		Logger: logger,
		gRPC:   grpc.NewServer(),
		server: &server.Server{
			DB: data.DataBase{db},
		},
		Lis: lis,
	}
}

func (a App) Run() {
	reflection.Register(a.gRPC)

	reflection.Register(a.gRPC)
	pb.RegisterJobServer(a.gRPC, a.server)

	go func() {
		a.Logger.Info().Msg("Starting gRPC Server on 0.0.0.0:50052")
		if err := a.gRPC.Serve(a.Lis); err != nil {
			a.Logger.Logger.Error().Err(err)
		}
	}()

	conn, err := grpc.DialContext(
		context.Background(),
		"0.0.0.0:50052",
		grpc.WithBlock(),
		grpc.WithTransportCredentials(insecure.NewCredentials()),
	)
	if err != nil {
		a.Logger.Logger.Error().Err(err)
	}

	mx := runtime.NewServeMux()
	err = pb.RegisterJobHandler(context.Background(), mx, conn)
	if err != nil {
		a.Logger.Logger.Error().Err(err)
	}

	s := &http.Server{
		Addr:    "0.0.0.0:50052",
		Handler: mx,
	}

	if err := s.ListenAndServe(); err != nil {
		a.Logger.Logger.Error().Err(err)
	}
}
