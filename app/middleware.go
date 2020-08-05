package elliot

import "net/http"

func mDomain(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if domain := r.FormValue("target"); len(domain) > 0 {
			RunDomainOSINT(domain)
		}
		next.ServeHTTP(w, r)
	})
}
