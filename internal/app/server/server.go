package server

import (
	"errors"
	"github.com/EestiChameleon/URLShortenerService/internal/app/cfg"
	"github.com/EestiChameleon/URLShortenerService/internal/app/server/grpcserver"
	"github.com/EestiChameleon/URLShortenerService/internal/app/server/httpserver"
)

var (
	SRV                  Server
	ErrUnknownServerType = errors.New("unknown server type")
)

type Server interface {
	Start() error
	ShutDown() error
}

func InitServer() (err error) {
	switch cfg.Envs.ServerType {
	case "http":
		SRV, err = httpserver.InitHttpServer()
		if err != nil {
			return err
		}

	case "grpc":
		SRV, err = grpcserver.InitGRPCServer()
		if err != nil {
			return err
		}

	default:
		return ErrUnknownServerType
	}

	return nil
}
