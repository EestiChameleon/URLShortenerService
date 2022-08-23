package grpcserver

import (
	"context"
	"errors"
	"fmt"
	"github.com/EestiChameleon/URLShortenerService/internal/app/service/process"
	"github.com/EestiChameleon/URLShortenerService/internal/app/storage"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"net"

	// импортируем пакет со сгенерированными protobuf-файлами
	pb "github.com/EestiChameleon/URLShortenerService/proto"
)

type GRPCServer struct {
	serv *grpc.Server
	pb.UnimplementedShortenerServer
}

func InitGRPCServer() (*GRPCServer, error) {
	// создаём gRPC-сервер без зарегистрированной службы
	s := grpc.NewServer()
	// регистрируем сервис
	pb.RegisterShortenerServer(s, &GRPCServer{})

	return &GRPCServer{serv: s}, nil
}

func (g *GRPCServer) Start() error {
	// определяем порт для сервера
	listen, err := net.Listen("tcp", ":3200")
	if err != nil {
		return err
	}

	fmt.Println("Сервер gRPC начал работу")
	// получаем запрос gRPC
	return g.serv.Serve(listen)
}

func (g *GRPCServer) ShutDown() error {
	g.serv.GracefulStop()
	return nil
}

// GetOrigURL .
func (g *GRPCServer) GetOrigURL(ctx context.Context, in *pb.GetOrigURLRequest) (*pb.GetOrigURLResponse, error) {

	origURL, err := storage.STRG.GetOrigURL(in.ShortUrl)
	if err != nil {
		return nil, err
	}

	return &pb.GetOrigURLResponse{OrigUrl: origURL}, nil
}

// GetAllPairs .
func (g *GRPCServer) GetAllPairs(ctx context.Context, in *pb.GetGetAllPairsRequest) (*pb.GetGetAllPairsResponse, error) {
	pairs, err := storage.STRG.GetUserURLs()
	if err != nil {
		return nil, err
	}

	if len(pairs) == 0 || pairs == nil {
		return nil, status.Error(codes.NotFound, `not found`)
	}

	var grpcPairs []*pb.Pair
	for _, v := range pairs {
		grpcPairs = append(grpcPairs, &pb.Pair{
			ShortUrl: v.ShortURL,
			OrigUrl:  v.OrigURL,
		})
	}
	return &pb.GetGetAllPairsResponse{Pairs: grpcPairs}, nil //???
}

// PostProvideShortURL .
func (g *GRPCServer) PostProvideShortURL(ctx context.Context, in *pb.PostProvideShortURLRequest) (*pb.PostProvideShortURLResponse, error) {

	if in.OrigUrl == "" {
		return nil, status.Error(codes.InvalidArgument, `invalid url`)
	}

	shortURL, err := process.ShortURLforOrigURL(in.OrigUrl)
	if err != nil {
		if errors.Is(err, storage.ErrDBOrigURLExists) {
			return &pb.PostProvideShortURLResponse{ShortUrl: shortURL},
				status.Error(codes.AlreadyExists, `provided url was already processed`)
		}
		return nil, status.Error(codes.InvalidArgument, `invalid url`)
	}

	return &pb.PostProvideShortURLResponse{ShortUrl: shortURL}, nil
}

func (g *GRPCServer) PostBatch(ctx context.Context, in *pb.PostBatchRequest) (*pb.PostBatchResponse, error) {
	if len(in.ReqPairs) == 0 {
		return nil, status.Error(codes.InvalidArgument, "empty data received")
	}

	// map of Correlation_id : original url pairs
	resp := &pb.PostBatchResponse{}
	for _, v := range in.ReqPairs {
		origURL := v.OrigUrl
		if origURL == "" {
			return nil, status.Error(codes.InvalidArgument, "empty value for original_url")
		}

		shortURL, err := process.ShortURLforOrigURL(origURL)
		if err != nil && !errors.Is(err, storage.ErrDBOrigURLExists) {
			return nil, status.Error(codes.Internal, "failed to process incoming data")
		}

		resp.RespPairs = append(resp.RespPairs, &pb.PostBatchResponsePair{
			CorrelationId: v.CorrelationId,
			ShortUrl:      shortURL,
		})
	}

	return resp, nil
}

// DelUser реализует интерфейс удаления информации о пользователе.
func (g *GRPCServer) DeleteBatch(ctx context.Context, in *pb.DelBatchRequest) (*pb.DelBatchResponse, error) {
	if len(in.ShortUrls) < 1 {
		return nil, status.Error(codes.InvalidArgument, "empty incoming values")
	}
	process.BatchDelete(in.ShortUrls)

	return &pb.DelBatchResponse{Error: "Accepted"}, nil
}

func (g *GRPCServer) GetStat(ctx context.Context, in *pb.GetStatRequest) (*pb.GetStatResponse, error) {
	urls, users, err := storage.STRG.GetStats()
	if err != nil {
		return nil, status.Error(codes.Internal, "failed to provide statistics")
	}
	return &pb.GetStatResponse{
		Urls:  int32(urls),
		Users: int32(users),
	}, nil
}
