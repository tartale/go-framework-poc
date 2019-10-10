package svc

import "context"

type GetBeerRequest struct {
}

type UnaryBeerResponse struct {
}

type IBeerServer interface {
	GetBeer(ctx context.Context, request *GetBeerRequest)
}

type BeerServer struct {
}

func NewBeerServer() *BeerServer {
	return &BeerServer{}
}

func (BeerServer) GetBeer(ctx context.Context, request *GetBeerRequest, response *UnaryBeerResponse) error {
	panic("implement me")
}
