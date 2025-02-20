package role

import (
	"net/http"

	"github.com/suyuan32/simple-admin-core/api/internal/logic/role"
	"github.com/suyuan32/simple-admin-core/api/internal/svc"
	"github.com/suyuan32/simple-admin-core/api/internal/types"
	"github.com/zeromicro/go-zero/rest/httpx"
)

// swagger:route post /role/status role SetRoleStatus
//
// Set role status | 设置角色状态
//
// Set role status | 设置角色状态
//
// Parameters:
//  + name: body
//    require: true
//    in: body
//    type: SetStatusReq
//
// Responses:
//  200: SimpleMsg
//  401: SimpleMsg
//  500: SimpleMsg

func SetRoleStatusHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.SetStatusReq
		if err := httpx.Parse(r, &req); err != nil {
			httpx.Error(w, err)
			return
		}

		l := role.NewSetRoleStatusLogic(r.Context(), svcCtx)
		resp, err := l.SetRoleStatus(&req)
		if err != nil {
			httpx.Error(w, err)
		} else {
			httpx.OkJson(w, resp)
		}
	}
}
