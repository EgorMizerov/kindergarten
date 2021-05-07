package app

import (
	"fmt"
	"github.com/EgorMizerov/kindergarten/internal/delivery/http"
	"github.com/EgorMizerov/kindergarten/internal/repository"
	"github.com/EgorMizerov/kindergarten/internal/service"
	"github.com/EgorMizerov/kindergarten/pkg/database"
	"github.com/joho/godotenv"
	"log"
	"os"
)

func Run() {
	fmt.Println("Start")

	// load environment
	err := godotenv.Load()
	if err != nil {
		log.Fatalf("error loading .env: %s", err.Error())
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

	fmt.Println(hand)

	fmt.Println("Finish")
}
