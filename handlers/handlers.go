package handlers

import (
	"context"

	"github.com/MarketScrapperAPI/ItemAPI/domain"
	pb "github.com/MarketScrapperAPI/ItemAPI/proto/gen"
	"github.com/MarketScrapperAPI/ItemAPI/repositories"
	"github.com/gofrs/uuid"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"gorm.io/gorm"
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

func (i *ItemHandler) CreateItem(ctx context.Context, req *pb.CreateItemRequest) (*pb.CreateItemResponse, error) {

	// Add req data Validation

	// Create new Item
	newItem := domain.NewItem()
	newItem.FromCreateItemRequest(req)

	// Add to db
	dbItem, err := i.repo.CreateItem(newItem)
	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	return dbItem.ToCreateItemResponse(), nil
}

func (i *ItemHandler) GetItem(ctx context.Context, req *pb.GetItemRequest) (*pb.GetItemResponse, error) {

	uuid, err := uuid.FromString(req.Uuid)
	if err != nil {
		return nil, status.Error(codes.InvalidArgument, "invalid uuid")
	}

	item, err := i.repo.GetItemById(uuid)
	if err == gorm.ErrRecordNotFound {
		return nil, status.Error(codes.NotFound, err.Error())
	}
	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	return item.ToGetItemResponse(), nil
}

func (i *ItemHandler) ListItems(ctx context.Context, req *pb.ListItemsRequest) (*pb.ListItemsResponse, error) {

	filters := make(map[string]interface{})

	// add filters from req
	if req.Name != "" {
		filters["name"] = req.Name
	}

	if req.PricePerItemGT != 0 {
		filters["pricePerItemGT"] = req.PricePerItemGT
	}

	if req.PricePerItemLT != 0 {
		filters["pricePerItemLT"] = req.PricePerItemLT
	}

	if req.PricePerQuantityGT != 0 {
		filters["pricePerQuantityGT"] = req.PricePerQuantityGT
	}

	if req.PricePerQuantityLT != 0 {
		filters["pricePerQuantityLT"] = req.PricePerQuantityLT
	}

	// query db
	items, err := i.repo.ListItem(filters)
	if len(items) == 0 {
		return nil, status.Error(codes.NotFound, "no items found")
	}
	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	rsp := pb.ListItemsResponse{}
	for _, v := range items {
		rsp.ListItemsResponse = append(rsp.ListItemsResponse, v.ToGetItemResponse())
	}

	return &rsp, nil
}

func (i *ItemHandler) UpdateItem(context.Context, *pb.UpdateItemRequest) (*pb.UpdateItemResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateItem not implemented")
}

func (i *ItemHandler) DeleteItem(context.Context, *pb.DeleteItemRequest) (*pb.DeleteItemResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteItem not implemented")
}
