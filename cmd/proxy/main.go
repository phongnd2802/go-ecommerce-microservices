package main

import (
	"context"
	"net/http"
	"os"

	_ "github.com/phongnd2802/go-ecommerce-microservices/docs/statik"
	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"github.com/phongnd2802/go-ecommerce-microservices/pb"
	"github.com/phongnd2802/go-ecommerce-microservices/pkg/config"
	"github.com/phongnd2802/go-ecommerce-microservices/pkg/logger"
	"github.com/rakyll/statik/fs"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/protobuf/encoding/protojson"
)

const (
	DEVELOPMENT = "development"
	PRODUCTION  = "production"
)

func newGateway(
	cfg *ProxyConfig,
	ctx context.Context,
	opts []runtime.ServeMuxOption,
) (http.Handler, error) {
	userEnpoint := cfg.UserAddr()
	mux := runtime.NewServeMux(opts...)

	dialOpts := []grpc.DialOption{grpc.WithTransportCredentials(insecure.NewCredentials())}

	err := pb.RegisterUserServiceHandlerFromEndpoint(ctx, mux, userEnpoint, dialOpts)
	if err != nil {
		return nil, err
	}

	return mux, nil
}

func main() {
	cfg := &ProxyConfig{}
	err := config.LoadConfig("./configs", "proxy", "yaml", cfg)
	if err != nil {
		log.Fatal().Msgf("failed to read config %v", err)
	}
	if cfg.Http.Mode == DEVELOPMENT {
		log.Logger = log.Output(zerolog.ConsoleWriter{Out: os.Stderr})
	}

	ctx := context.Background()
	ctx, cancel := context.WithCancel(ctx)
	defer cancel()

	mux := http.NewServeMux()

	opts := []runtime.ServeMuxOption{
		runtime.WithMarshalerOption(runtime.MIMEWildcard, &runtime.JSONPb{
			MarshalOptions: protojson.MarshalOptions{
				UseProtoNames: true,
			},
			UnmarshalOptions: protojson.UnmarshalOptions{
				DiscardUnknown: true,
			},
		}),
	}	
	gw, err := newGateway(cfg, ctx, opts)
	if err != nil {
		log.Fatal().Msgf("failed to create a new gateway: %v", err)
	}

	mux.Handle("/", gw)

	statikFS, err := fs.New()
	if err != nil {
		log.Fatal().Msgf("cannot create statik fs: %v", err)
	}
	swaggerHandler := http.StripPrefix("/docs/", http.FileServer(statikFS))
	mux.Handle("/docs/", swaggerHandler)

	server := &http.Server{
		Addr: cfg.HttpAddr(),
		Handler: logger.HttpLogger(mux),
	}
	log.Info().Msgf("Start API Gateway at: %s", cfg.HttpAddr())
	if err := server.ListenAndServe(); err != nil {
		log.Fatal().Msgf("failed to listen and serve %v", err)
	}
}