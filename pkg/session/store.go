package session

import (
	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/cookie"
	"github.com/klein/go-mvc/pkg/config"
)

func Store() sessions.Store {

	var store sessions.Store
	switch config.GetString("session.driver") {
	case "cookie":
		store = cookie.NewStore([]byte(config.GetString("app.key")))
	}
	return store
}
