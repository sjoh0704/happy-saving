package middleware

import (
	// "log"
	"net/http"

	log "github.com/sirupsen/logrus"
	"github.com/sjoh0704/happysaving/auth"
	"github.com/sjoh0704/happysaving/util"
	// "github.com/gorilla/mux"
)

// Middleware function, which will be called for each request
func TokenAuthMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		// user 생성 요청이거나 로그인 요청일 때 token 검사를 하면 안됨
		if (r.URL.Path == "/apis/v1/users" && r.Method == "POST") ||
			r.URL.Path == "/auth" ||
			r.URL.Path == "/ready" {
			next.ServeHTTP(w, r)
			return
		}

		// access token을 받아서
		token := r.Header.Get("access-token")

		if verify, err := auth.VerifiyJWTToken(token); verify {
			next.ServeHTTP(w, r)
		} else {
			log.Error("not authenticated request")
			util.SetResponse(w, "Forbidden", err, http.StatusForbidden)
		}
	})
}

// func Example_authenticationMiddleware() {
// 	r := mux.NewRouter()
// 	r.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
// 		// Do something here
// 	})
// 	// amw := authenticationMiddleware{make(map[string]string)}
// 	// amw.Populate()
// 	r.Use(TokenAuthMiddleware)
// }
