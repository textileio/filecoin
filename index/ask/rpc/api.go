package rpc

import (
	"context"

	pb "github.com/textileio/powergate/index/ask/pb"
	"github.com/textileio/powergate/index/ask/runner"
)

// Service implements the gprc service
type Service struct {
	UnimplementedAPIServer

	index *runner.Runner
}

// New is a helper to create a new Service.
func New(ai *runner.Runner) *Service {
	return &Service{
		index: ai,
	}
}

// Get returns the current Ask Storage index.
func (s *Service) Get(ctx context.Context, req *pb.GetRequest) (*pb.GetReply, error) {
	index := s.index.Get()
	storage := make(map[string]*StorageAsk, len(index.Storage))
	for key, ask := range index.Storage {
		storage[key] = &StorageAsk{
			Price:        ask.Price,
			MinPieceSize: ask.MinPieceSize,
			Miner:        ask.Miner,
			Timestamp:    ask.Timestamp,
			Expiry:       ask.Expiry,
		}
	}
	pbIndex := &Index{
		LastUpdated:        index.LastUpdated.Unix(),
		StorageMedianPrice: index.StorageMedianPrice,
		Storage:            storage,
	}
	return &GetReply{Index: pbIndex}, nil
}

// Query executes a query on the current Ask Storage index.
func (s *Service) Query(ctx context.Context, req *pb.QueryRequest) (*pb.QueryReply, error) {
	q := runner.Query{
		MaxPrice:  req.GetQuery().GetMaxPrice(),
		PieceSize: req.GetQuery().GetPieceSize(),
		Limit:     int(req.GetQuery().GetLimit()),
		Offset:    int(req.GetQuery().GetOffset()),
	}
	asks, err := s.index.Query(q)
	if err != nil {
		return nil, err
	}
	replyAsks := make([]*StorageAsk, len(asks))
	for i, ask := range asks {
		replyAsks[i] = &StorageAsk{
			Price:        ask.Price,
			MinPieceSize: ask.MinPieceSize,
			Miner:        ask.Miner,
			Timestamp:    ask.Timestamp,
			Expiry:       ask.Expiry,
		}
	}
	return &QueryReply{Asks: replyAsks}, nil
}
