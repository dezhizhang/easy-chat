package middleware

import "net/http"

type LoginVerifyMiddleware struct {
}

func NewLoginVerifyMiddleware() *LoginVerifyMiddleware {
	return &LoginVerifyMiddleware{}
}

func (m *LoginVerifyMiddleware) Handle(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// TODO generate middleware implement function, delete after code implementation

		if r.Header.Get("token") == "tom" {
			next(w, r)
			return
		}

		w.Write([]byte("权限不足"))

	}
}
