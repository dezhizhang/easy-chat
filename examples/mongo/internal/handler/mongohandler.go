package handler

import (
	"net/http"

	"easy-chat/examples/mongo/internal/logic"
	"easy-chat/examples/mongo/internal/svc"
	"easy-chat/examples/mongo/internal/types"
	"github.com/zeromicro/go-zero/rest/httpx"
)

func MongoHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.Request
		if err := httpx.Parse(r, &req); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}

		l := logic.NewMongoLogic(r.Context(), svcCtx)
		resp, err := l.Mongo(&req)
		if err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.OkJsonCtx(r.Context(), w, resp)
		}
	}
}
