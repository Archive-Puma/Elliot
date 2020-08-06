package elliot

import (
	"net/http"
	"strings"

	"github.com/cosasdepuma/elliot/app/modules"
)

func mDomain(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if domain := r.FormValue("domain"); len(domain) > 0 {
			Backend.DB.Purge()
			modules.RunDomain(strings.TrimSpace(strings.ToLower(domain)), Backend.DB)
		}
		next.ServeHTTP(w, r)
	})
}
