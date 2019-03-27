package utils

import (
	ginsessions "github.com/gin-contrib/sessions"
	"github.com/gorilla/sessions"
)

type store struct {
	*sessions.CookieStore
}

func (s *store) Options(options ginsessions.Options) {
	options = ginsessions.Options{
		Path:     "/",
		HttpOnly: false,
		Secure:   false,
		MaxAge:   3600,
	}
}

func NewStore(cookieStore *sessions.CookieStore) ginsessions.Store {
	return &store{cookieStore}
}
