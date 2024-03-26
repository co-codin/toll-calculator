package main

import (
	"context"

	"github.com/co-codin/toll-calculator/types"
)

type GPRCAggregatorServer struct {
	types.UnimplementedAggregatorServer
	svc Aggregator
}

func NewAggregatorGRPCServer(svc Aggregator) *GPRCAggregatorServer {
	return &GPRCAggregatorServer{
		svc: svc,
	}
}

// transport layer
// JSON -> types.Distance -> all done (same type)
// GRPC -> types.AggregateRequest -> type.Distance
// Webpack => types.WEBpack -> types.Distance

func (s *GPRCAggregatorServer) Aggregate(ctx context.Context, req *types.AggregateRequest) (*types.None, error) {
	distance := types.Distance{
		OBUID: int(req.ObuID),
		Value: req.Value,
		Unix:  req.Unix,
	}
	return &types.None{}, s.svc.AggregateDistance(distance)
}
