package api

import (
	"context"

	"github.com/suyuan32/simple-admin-core/rpc/types/core"

	"github.com/zeromicro/go-zero/core/logx"

	"github.com/suyuan32/simple-admin-core/api/internal/svc"
	"github.com/suyuan32/simple-admin-core/api/internal/types"
)

type DeleteApiLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewDeleteApiLogic(ctx context.Context, svcCtx *svc.ServiceContext) *DeleteApiLogic {
	return &DeleteApiLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *DeleteApiLogic) DeleteApi(req *types.IDReq) (resp *types.SimpleMsg, err error) {
	data, err := l.svcCtx.CoreRpc.DeleteApi(context.Background(), &core.IDReq{
		ID: uint64(req.ID),
	})
	if err != nil {
		return nil, err
	}
	resp = &types.SimpleMsg{}
	resp.Msg = data.Msg
	return resp, nil
}
