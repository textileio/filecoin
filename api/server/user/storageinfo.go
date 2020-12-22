package user

import (
	"context"

	"github.com/ipfs/go-cid"
	userPb "github.com/textileio/powergate/api/gen/powergate/user/v1"
	"github.com/textileio/powergate/ffs/api"
	"github.com/textileio/powergate/util"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

// StorageInfo returns the information about a stored Cid. If no information is available,
// since the Cid was never stored, it returns an error with codes.NotFound.
func (s *Service) StorageInfo(ctx context.Context, req *userPb.StorageInfoRequest) (*userPb.StorageInfoResponse, error) {
	i, err := s.getInstanceByToken(ctx)
	if err != nil {
		return nil, status.Errorf(codes.PermissionDenied, "getting instance: %v", err)
	}
	cid, err := util.CidFromString(req.Cid)
	if err != nil {
		return nil, status.Errorf(codes.InvalidArgument, "parsing cid: %v", err)
	}
	info, err := i.StorageInfo(cid)
	if err == api.ErrNotFound {
		return nil, status.Errorf(codes.NotFound, "getting storage info: %v", err)
	}
	if err != nil {
		return nil, status.Errorf(codes.Internal, "getting storage info: %v", err)
	}
	pbInfo := toRPCStorageInfo(info)
	return &userPb.StorageInfoResponse{StorageInfo: pbInfo}, nil
}

// QueryStorageInfo returns a list of infomration about all stored cids, filtered by cids if provided.
func (s *Service) QueryStorageInfo(ctx context.Context, req *userPb.QueryStorageInfoRequest) (*userPb.QueryStorageInfoResponse, error) {
	i, err := s.getInstanceByToken(ctx)
	if err != nil {
		return nil, status.Errorf(codes.PermissionDenied, "getting instance: %v", err)
	}
	cids := make([]cid.Cid, len(req.Cids))
	for i, s := range req.Cids {
		c, err := util.CidFromString(s)
		if err != nil {
			return nil, status.Errorf(codes.InvalidArgument, "parsing cid: %v", err)
		}
		cids[i] = c
	}
	infos, err := i.QueryStorageInfo(cids...)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "querying storage info: %v", err)
	}
	res := make([]*userPb.StorageInfo, len(infos))
	for i, info := range infos {
		res[i] = toRPCStorageInfo(info)
	}
	return &userPb.QueryStorageInfoResponse{StorageInfo: res}, nil
}
