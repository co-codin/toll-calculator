package main

import "github.com/co-codin/toll-calculator/types"

type Aggregator interface {
	AggregateDistance(types.Distance) error
}

type Storer interface {
	Insert(types.Distance) error
}

type InvoiceAggregator struct {
	store Storer
}

func (i *InvoiceAggregator) AggregateDistance(distance types.Distance) error {
	i.store.Insert(distance)
	return nil
}
