package api

import (
	"log"
	"net"

	"github.com/MarketScrapperAPI/ItemAPI/domain"
	"github.com/MarketScrapperAPI/ItemAPI/handlers"
	pb "github.com/MarketScrapperAPI/ItemAPI/proto/gen"
	"github.com/MarketScrapperAPI/ItemAPI/repositories"
	"google.golang.org/grpc"
	"gorm.io/gorm"
)

type Api struct {
	srv         *grpc.Server
	lis         net.Listener
	itemHandler *handlers.ItemHandler
}

func NewApi(db *gorm.DB, port string) *Api {

	err := db.AutoMigrate(domain.Item{})
	if err != nil {
		panic(err)
	}

	// create repositories
	itemRepo := repositories.NewItemRepository(db)

	// create listener
	lis, err := net.Listen("tcp", ":"+port)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	// create handler
	itemHandler := handlers.NewItemHandler(itemRepo)

	//create server
	grpcServer := grpc.NewServer()
	pb.RegisterItemApiServer(grpcServer, itemHandler)

	return &Api{
		srv:         grpcServer,
		lis:         lis,
		itemHandler: itemHandler,
	}
}

func (a *Api) Start() {
	log.Println("Start server on: ", a.lis.Addr().String())

	if err := a.srv.Serve(a.lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
