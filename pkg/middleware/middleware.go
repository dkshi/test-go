package middleware

import (
	"log"
	"net/http"
	"strings"
)

func CheckAdmin(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		var header http.Header = r.Header

		userRole := header.Get("User-Role")

		if strings.Contains(userRole, "admin") {
			log.Println("red button user detected")
		}

		next.ServeHTTP(w, r)
	})
}
