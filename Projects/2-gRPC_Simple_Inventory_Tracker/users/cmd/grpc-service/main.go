package main

import (
	"context"
	"fmt"
	"net"
	"net/url"
	"os"
	"os/signal"
	"sync"
	"syscall"

	"github.com/I-Frostbyte/Github-Repos-New-2026/Projects/2-gRPC_Simple_Inventory_Tracker/grpc/go/usersgrpc"
	"github.com/I-Frostbyte/Github-Repos-New-2026/Projects/2-gRPC_Simple_Inventory_Tracker/users/public"
	"github.com/I-Frostbyte/Github-Repos-New-2026/Projects/2-gRPC_Simple_Inventory_Tracker/users/users"
	grpc_recovery "github.com/grpc-ecosystem/go-grpc-middleware/recovery"
	"github.com/rs/zerolog"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

func main() {
	// Setting up context.
	// A context.Context is an object used to carry info across function calls, primarily:
	// cancellations, timeouts, deadlines & request-scoped values.
	// Think of it as a control signal that many parts of the program can listen to.
	// Here we are creating a background context that CAN be cancelled
	// specifically if the system sends syscall.SIGINT (Ctrl + C) or syscall.SIGTERM.
	// SIGTERM is usually sent when another program asks yours to shut down.
	// e.g. Docker stopping a container, Kubernetes terminating a pod, kill <pid>,
	// systemd stopping a service
	ctx, stop := signal.NotifyContext(context.Background(), syscall.SIGINT, syscall.SIGTERM)
	defer stop()

	// Creating the logger.
	// zerolog.New() creates a new logger i.e. writing simple JSON logs to standard errors.
	// .With() adds additional fields to this new logger like Timestamp & Caller.
	// Timestamp adds time to every log.
	// Caller automatically records file name, line number.
	// Logger finishes building the logger
	logger := zerolog.New(os.Stderr).With().Timestamp().Caller().Logger()

	err := run(ctx, logger)
	if err != nil {
		logger.Err(err).Msg("failed to run grpc service")
		os.Exit(1)
	}
}

func run(ctx context.Context, logger zerolog.Logger) error {
	logger.Info().Msg("Starting grpc service...")

	config := model.Config{}

	err := config.LoadConfig()
	if err != nil {
		logger.Err(err).Msg("failed to load config")
		return err
	}

	logger.Info().Msgf("Successfully loaded config...: %+v", config)

	// Configuring the new loglevel (Debug) and using it
	// to create a new logger that begins at that level.
	logLevel, err := zerolog.ParseLevel(config.LogLevel)
	if err != nil {
		return fmt.Errorf("failed to parse log level: %w", err)
	}
	logger = logger.Level(logLevel)

	/*
		dbConnectionURL := getPostgresConnectionURL(config.DB)
		db, err := pgxpool.New(ctx, dbConnectionURL)
		if err != nil {
			return fmt.Errorf("failed to connect to database: %w", err)
		}

		defer db.Close()
	*/

	// Starting the gRPC server in a separate goroutine

	// ServerOption is configuration object for the server.
	// Here we create a single slice of the ServerOption
	// So that when we declare, we can make way for dynamic adding.
	// grpc_recovery.UnaryServerInterceptor() acts as middleware.
	// It intercepts any panics and returns them as internal errors,
	// so that the server may continue running.
	svrOpts := []grpc.ServerOption{
		grpc.ChainUnaryInterceptor(
			grpc_recovery.UnaryServerInterceptor(),
		),
	}

	// Initiating a new gRPC server.
	// Reflection allows for things like Postman to
	// automatically pick it up.
	grpcServer := grpc.NewServer(svrOpts...)
	reflection.Register(grpcServer)

	usersgrpc.RegisterUsersServiceServer(grpcServer, users.NewUsersService())
	logger.Info().Msg("Successfully registered UsersServiceServer...")

	listener, err := net.Listen("tcp", fmt.Sprintf(":%d", config.ListenPort))
	if err != nil {
		return fmt.Errorf("net.Listen: %w", err)
	}

	logger.Info().Msgf(`grpc service is listening on port: %s`, listener.Addr().String())

	var startupErr error
	wg := &sync.WaitGroup{}
	wg.Add(1)
	go func() {
		defer wg.Done()
		err = grpcServer.Serve(listener)
		if err != nil {
			startupErr = fmt.Errorf("error starting gRPC server: %w", err)
		}
	}()

	// Do a graceful shutdown if the context is cancelled.
	go func() {
		<-ctx.Done() // sleep until context is cancelled.
		logger.Info().Msg("Shutting down gRPC server...")
		grpcServer.GracefulStop()
		logger.Info().Msg("gRPC server stopped.")
	}()

	logger.Info().Msgf(`HTTP server running on %s`, listener.Addr().String())

	// Graceful shutdown logic
	// wait for the context to finish
	wg.Wait()
	logger.Info().Msg("gRPC server has shut down gracefully...")

	return startupErr
}

func getPostgresConnectionURL(config model.DBConfig) string {
	queryValues := url.Values{}
	if config.TLSDisabled {
		queryValues.Add("sslmode", "disable")
	} else {
		queryValues.Add("sslmode", "require")
	}

	dbURL := url.URL{
		Scheme:   "postgres",
		User:     url.UserPassword(config.DBUser, config.DBPassword),
		Host:     fmt.Sprintf("%s:%d", config.DBHost, config.DBPort),
		Path:     config.DBName,
		RawQuery: queryValues.Encode(),
	}

	return dbURL.String()
}

/*

The Lifecycle of run()

Start
  │
  ▼
Log startup
  │
  ▼
Load configuration
  │
  ▼
Configure logger
  │
  ▼
Build PostgreSQL connection URL
  │
  ▼
Connect to PostgreSQL
  │
  ▼
Create gRPC server
  │
  ▼
Register RPC services
  │
  ▼
Open TCP listener
  │
  ▼
Start serving requests (goroutine)
  │
  ├───────────────┐
  │               │
  ▼               ▼
Wait for      Wait for
RPC traffic   context cancellation (Ctrl+C/SIGTERM)
                  │
                  ▼
          Gracefully stop server
                  │
                  ▼
          Serve() returns
                  │
                  ▼
           WaitGroup reaches zero
                  │
                  ▼
          Return to `main()`


*/

/*

MENTAL MODEL OF SERVER:

main()

    │
    ▼

run()

    │
    ├──────────────► Goroutine 1
    │               grpcServer.Serve()
    │               Accept client requests
    │
    ├──────────────► Goroutine 2
    │               Wait for Ctrl+C / SIGTERM
    │               Call GracefulStop()
    │
    ▼
WaitGroup.Wait()
    │
    ▼
Return to main()

*/