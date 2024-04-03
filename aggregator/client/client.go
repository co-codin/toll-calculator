package client

import (
	"context"

	"github.com/co-codin/toll-calculator/types"
)

type Client interface {
	Aggregate(context.Context, *types.AggregateRequest) error
}