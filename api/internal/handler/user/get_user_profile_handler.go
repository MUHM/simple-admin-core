package user

import (
	"net/http"

	"github.com/suyuan32/simple-admin-core/api/internal/logic/user"
	"github.com/suyuan32/simple-admin-core/api/internal/svc"
	"github.com/zeromicro/go-zero/rest/httpx"
)

// swagger:route get /user/profile user GetUserProfile
//
// Get user's profile | 获取用户个人信息
//
// Get user's profile | 获取用户个人信息
//
// Responses:
//  200: ProfileResp
//  401: SimpleMsg
//  500: SimpleMsg

func GetUserProfileHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		l := user.NewGetUserProfileLogic(r.Context(), svcCtx)
		resp, err := l.GetUserProfile()
		if err != nil {
			httpx.Error(w, err)
		} else {
			httpx.OkJson(w, resp)
		}
	}
}
