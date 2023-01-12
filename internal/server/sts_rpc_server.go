// Code generated by goctl. DO NOT EDIT.
// Source: sts.proto

package server

import (
	"context"

	"github.com/xh-polaris/sts-rpc/internal/logic"
	"github.com/xh-polaris/sts-rpc/internal/svc"
	"github.com/xh-polaris/sts-rpc/pb"
)

type StsRpcServer struct {
	svcCtx *svc.ServiceContext
	pb.UnimplementedStsRpcServer
}

func NewStsRpcServer(svcCtx *svc.ServiceContext) *StsRpcServer {
	return &StsRpcServer{
		svcCtx: svcCtx,
	}
}

func (s *StsRpcServer) GenCosSts(ctx context.Context, in *pb.GenCosStsReq) (*pb.GenCosStsResp, error) {
	l := logic.NewGenCosStsLogic(ctx, s.svcCtx)
	return l.GenCosSts(in)
}

func (s *StsRpcServer) GenSignedUrl(ctx context.Context, in *pb.GenSignedUrlReq) (*pb.GenSignedUrlResp, error) {
	l := logic.NewGenSignedUrlLogic(ctx, s.svcCtx)
	return l.GenSignedUrl(in)
}
