package rpc

import (
	"context"

	"github.com/textileio/powergate/index/miner"
)

// Service implements the gprc service
type Service struct {
	UnimplementedAPIServer

	index *miner.Index
}

// NewService is a helper to create a new Service
func NewService(mi *miner.Index) *Service {
	return &Service{
		index: mi,
	}
}

// Get calls miner index Get
func (s *Service) Get(ctx context.Context, req *GetRequest) (*GetReply, error) {
	index := s.index.Get()

	info := make(map[string]*Meta, len(index.Meta.Info))
	for key, meta := range index.Meta.Info {
		location := &Location{
			Country:   meta.Location.Country,
			Longitude: meta.Location.Longitude,
			Latitude:  meta.Location.Latitude,
		}
		info[key] = &Meta{
			LastUpdated: meta.LastUpdated.Unix(),
			UserAgent:   meta.UserAgent,
			Location:    location,
			Online:      meta.Online,
		}
	}

	pbPower := make(map[string]*Power, len(index.Chain.Power))
	for key, power := range index.Chain.Power {
		pbPower[key] = &Power{
			Power:    power.Power,
			Relative: float32(power.Relative),
		}
	}

	meta := &MetaIndex{
		Online:  index.Meta.Online,
		Offline: index.Meta.Offline,
		Info:    info,
	}

	chain := &ChainIndex{
		LastUpdated: index.Chain.LastUpdated,
		Power:       pbPower,
	}

	pbIndex := &Index{
		Meta:  meta,
		Chain: chain,
	}

	return &GetReply{Index: pbIndex}, nil
}