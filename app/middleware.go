package elliot

import (
	"net/http"

	"github.com/cosasdepuma/elliot/app/modules"
)

func mDomain(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		Backend.DB.Purge()
		if domain := r.FormValue("domain"); len(domain) > 0 {
			modules.RunDomain(domain, Backend.DB)
		}
		next.ServeHTTP(w, r)
	})
}
