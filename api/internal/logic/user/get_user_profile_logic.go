package user

import (
	"context"

	"github.com/suyuan32/simple-admin-core/rpc/types/core"

	"github.com/zeromicro/go-zero/core/logx"

	"github.com/suyuan32/simple-admin-core/api/internal/svc"
	"github.com/suyuan32/simple-admin-core/api/internal/types"
)

type GetUserProfileLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetUserProfileLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetUserProfileLogic {
	return &GetUserProfileLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetUserProfileLogic) GetUserProfile() (resp *types.ProfileResp, err error) {
	result, err := l.svcCtx.CoreRpc.GetUserById(l.ctx, &core.UUIDReq{
		UUID: l.ctx.Value("userId").(string),
	})
	if err != nil {
		return nil, err
	}

	resp = &types.ProfileResp{
		Nickname: result.Nickname,
		Avatar:   result.Avatar,
		Mobile:   result.Mobile,
		Email:    result.Email,
	}

	return resp, nil
}
