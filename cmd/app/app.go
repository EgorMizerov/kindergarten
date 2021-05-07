package app

import (
	"context"
	"github.com/EgorMizerov/kindergarten/internal/delivery/http"
	"github.com/EgorMizerov/kindergarten/internal/repository"
	"github.com/EgorMizerov/kindergarten/internal/service"
	"github.com/EgorMizerov/kindergarten/pkg/database"
	pkgserver "github.com/EgorMizerov/kindergarten/pkg/server"
	"github.com/joho/godotenv"
	"github.com/spf13/viper"
	"log"
	"os"
	"os/signal"
	"syscall"
)

func Run() { // load environment
	err := godotenv.Load()
	if err != nil {
		log.Fatalf("error loading .env: %s", err.Error())
	}

	// read config
	viper.SetConfigFile("./config.yaml")
	err = viper.ReadInConfig()
	if err != nil {
		log.Fatalf("error reading config: %s", err.Error())
	}

	// connect to MongoDB
	query := os.Getenv("MONGODB_CONNECTION_QUERY")
	client, err := database.ConnectClient(query)
	if err != nil {
		log.Fatalf("error connection to mongodb: %s", err.Error())
	}

	// initializing layers
	repo := repository.NewRepository(client)
	serv := service.NewService(repo)
	hand := http.NewHandler(serv)

	server := new(pkgserver.Server)

	// get addr
	host := viper.GetString("server.host")
	port := viper.GetString("server.port")
	if port == "" {
		port = ":8080"
	}
	addr := host + port

	// running server
	go func() {
		err = server.RunServer(addr, hand.Init())
	}()
	log.Printf("http://%s", addr)

	// create chan for notify unix signals
	quit := make(chan os.Signal)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit

	// shutting server
	err = server.Shutdown(context.Background())
	if err != nil {
		log.Fatalf("error shutting server: %s", err.Error())
	}
}
