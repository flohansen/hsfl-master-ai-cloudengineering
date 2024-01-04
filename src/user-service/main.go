package main

import (
	"context"
	"errors"
	"fmt"
	"google.golang.org/grpc"
	proto "hsfl.de/group6/hsfl-master-ai-cloud-engineering/lib/rpc/user"
	"hsfl.de/group6/hsfl-master-ai-cloud-engineering/user-service/api/http/handler"
	"hsfl.de/group6/hsfl-master-ai-cloud-engineering/user-service/api/http/middleware"
	"hsfl.de/group6/hsfl-master-ai-cloud-engineering/user-service/api/http/router"
	"hsfl.de/group6/hsfl-master-ai-cloud-engineering/user-service/api/rpc"
	"hsfl.de/group6/hsfl-master-ai-cloud-engineering/user-service/auth"
	"hsfl.de/group6/hsfl-master-ai-cloud-engineering/user-service/auth/utils"
	"hsfl.de/group6/hsfl-master-ai-cloud-engineering/user-service/crypto"
	"hsfl.de/group6/hsfl-master-ai-cloud-engineering/user-service/user"
	"hsfl.de/group6/hsfl-master-ai-cloud-engineering/user-service/user/model"
	"log"
	"net"
	"net/http"
	"os"
	"os/signal"
	"sync"
	"syscall"
)

func main() {
	var usersRepository user.Repository = user.NewDemoRepository()
	var usersController user.Controller = user.NewDefaultController(usersRepository)

	var tokenGenerator = createTokenGenerator()

	var wg sync.WaitGroup

	ctx, cancel := context.WithCancel(context.Background())

	wg.Add(1)
	go startHTTPServer(ctx, &wg, &usersController, &usersRepository, tokenGenerator)

	wg.Add(1)
	go startGRPCServer(ctx, &wg, &usersRepository, tokenGenerator)

	stopChan := make(chan os.Signal, 1)
	signal.Notify(stopChan, syscall.SIGINT, syscall.SIGTERM)

	<-stopChan
	cancel()

	wg.Wait()
}

func startHTTPServer(ctx context.Context, wg *sync.WaitGroup, usersController *user.Controller, usersRepository *user.Repository, tokenGenerator auth.TokenGenerator) {
	defer wg.Done()

	var loginHandler = createLoginHandler(*usersRepository, tokenGenerator)
	var registerHandler = createRegisterHandler(*usersRepository)

	authMiddleware := middleware.CreateLocalAuthMiddleware(usersRepository, tokenGenerator)
	handler := router.New(loginHandler, registerHandler, usersController, authMiddleware)
	server := &http.Server{Addr: ":3001", Handler: handler}

	go func() {
		if err := server.ListenAndServe(); err != nil && !errors.Is(err, http.ErrServerClosed) {
			log.Fatalf("Failed to start HTTP server: %v", err)
		}
	}()

	<-ctx.Done()

	if err := server.Shutdown(context.Background()); err != nil {
		log.Fatalf("HTTP Server Shutdown Failed:%v", err)
	}
}

func startGRPCServer(ctx context.Context, wg *sync.WaitGroup, usersRepository *user.Repository, tokenGenerator auth.TokenGenerator) {
	defer wg.Done()

	lis, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatalf("Failed to listen: %v", err)
	}

	grpcServer := grpc.NewServer()
	userServiceServer := rpc.NewUserServiceServer(usersRepository, tokenGenerator)
	proto.RegisterUserServiceServer(grpcServer, userServiceServer)

	go func() {
		if err := grpcServer.Serve(lis); err != nil {
			log.Fatalf("Failed to start gRPC server: %v", err)
		}
	}()

	<-ctx.Done()
	grpcServer.GracefulStop()
}

func createTokenGenerator() auth.TokenGenerator {
	privateKey := utils.GenerateRandomECDSAPrivateKeyAsPEM()
	tokenGenerator, err := auth.NewJwtTokenGenerator(auth.JwtConfig{PrivateKey: privateKey})
	if err != nil {
		panic(fmt.Sprintf("Can't generate token generator: %v", err))
	}
	return tokenGenerator
}

func createLoginHandler(userRepository user.Repository, tokenGenerator auth.TokenGenerator) *handler.LoginHandler {
	return handler.NewLoginHandler(createMockRepository(userRepository),
		crypto.NewBcryptHasher(), tokenGenerator)
}

func createRegisterHandler(userRepository user.Repository) *handler.RegisterHandler {
	return handler.NewRegisterHandler(createMockRepository(userRepository),
		crypto.NewBcryptHasher())
}

func createMockRepository(userRepository user.Repository) user.Repository {
	userSlice := createDemoUserSlice()
	for _, newUser := range userSlice {
		_, _ = userRepository.Create(newUser)
	}

	return userRepository
}

func createDemoUserSlice() []*model.User {
	bcryptHasher := crypto.NewBcryptHasher()
	hashedPassword, _ := bcryptHasher.Hash([]byte("12345"))

	return []*model.User{
		{
			Id:       1,
			Email:    "ada.lovelace@gmail.com",
			Password: hashedPassword,
			Name:     "Ada Lovelace",
			Role:     model.Customer,
		},
		{
			Id:       2,
			Email:    "info-aldi@gmail.com",
			Password: hashedPassword,
			Name:     "Aldi",
			Role:     model.Merchant,
		},
		{
			Id:       3,
			Email:    "info-edeka@gmail.com",
			Password: hashedPassword,
			Name:     "Edeka",
			Role:     model.Merchant,
		},
	}
}
