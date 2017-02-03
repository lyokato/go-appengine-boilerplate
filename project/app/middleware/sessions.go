package middleware

import (
	"app/config"

	s "github.com/gin-contrib/sessions"
	gin "gopkg.in/gin-gonic/gin.v1"
)

func Sessions(cnf *config.SessionConfig) gin.HandlerFunc {
	store := s.NewCookieStore([]byte(cnf.Secret))
	store.Options(s.Options{
		Path:     cnf.Path,
		Domain:   cnf.Domain,
		MaxAge:   cnf.MaxAge,
		Secure:   cnf.Secure,
		HttpOnly: false,
	})
	return s.Sessions(cnf.Name, store)
}
