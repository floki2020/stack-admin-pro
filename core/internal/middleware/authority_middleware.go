package middleware

import (
	"github.com/casbin/casbin/v2"
	"github.com/zeromicro/go-zero/core/errorx"
	"github.com/zeromicro/go-zero/core/logx"
	"github.com/zeromicro/go-zero/core/stores/redis"
	"github.com/zeromicro/go-zero/rest/httpx"
	"net/http"
	"strings"
)

type AuthorityMiddleware struct {
	Cbn *casbin.Enforcer
	Rds *redis.Redis
}

func NewAuthorityMiddleware(cbn *casbin.Enforcer, rds *redis.Redis) *AuthorityMiddleware {
	return &AuthorityMiddleware{
		Cbn: cbn,
		Rds: rds,
	}
}

func (m *AuthorityMiddleware) Handle(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		obj := r.URL.Path

		act := r.Method

		roleIds := r.Context().Value("roleId").(string)

		// check jwt blacklist
		jwtResult, err := m.Rds.Get("token_" + r.Header.Get("Authorization"))
		if err != nil {
			logx.Errorw("redis error in jwt", logx.Field("detail", err.Error()))
			httpx.Error(w, errorx.NewApiError(http.StatusInternalServerError, err.Error()))
			return
		}
		if jwtResult == "1" {
			logx.Errorw("token in blacklist", logx.Field("detail", r.Header.Get("Authorization")))
			httpx.Error(w, errorx.NewApiErrorWithoutMsg(http.StatusUnauthorized))
			return
		}
	}
}

func batchCheck(cbn *casbin.Enforcer, roleIds, act, obj string) bool {
	var checkReq [][]any
	for _, v := range strings.Split(roleIds, ",") {
		checkReq = append(checkReq, []any{v, obj, act})
	}

	result, err := cbn.BatchEnforce(checkReq)
	if err != nil {
		logx.Errorw("Casbin enforce error", logx.Field("detail", err.Error()))
		return false
	}

	for _, v := range result {
		if v {
			return true
		}
	}

	return false
}
