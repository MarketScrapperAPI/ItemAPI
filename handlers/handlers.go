package handlers

import (
	"context"

	pb "github.com/MarketScrapperAPI/ItemAPI/proto/gen"
	"github.com/MarketScrapperAPI/ItemAPI/repositories"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type ItemHandler struct {
	pb.UnimplementedItemApiServer
	repo *repositories.ItemRepository
}

func NewItemHandler(repo *repositories.ItemRepository) *ItemHandler {
	return &ItemHandler{
		repo: repo,
	}
}

func (i *ItemHandler) CreateItem(context.Context, *pb.CreateItemRequest) (*pb.CreateItemResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetItem not implemented")
}

func (i *ItemHandler) GetItem(context.Context, *pb.GetItemRequest) (*pb.GetItemResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetItem not implemented")
}

func (i *ItemHandler) ListItems(context.Context, *pb.ListItemsRequest) (*pb.ListItemsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListItems not implemented")
}

func (i *ItemHandler) UpdateItem(context.Context, *pb.UpdateItemRequest) (*pb.UpdateItemResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateItem not implemented")
}

func (i *ItemHandler) DeleteItem(context.Context, *pb.DeleteItemRequest) (*pb.DeleteItemResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteItem not implemented")
}
