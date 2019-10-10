package svc

import "context"

type GetBeerRequest struct {
}

type UnaryBeerResponse struct {
}

type BeerSvcHandler interface {
	GetBeer(context.Context, *GetBeerRequest, *UnaryBeerResponse) error
}

type BeerServer struct {
}

func NewBeerServer() BeerSvcHandler {
	return &BeerServer{}
}

func (BeerServer) GetBeer(ctx context.Context, request *GetBeerRequest, response *UnaryBeerResponse) error {
	panic("implement me")
}
