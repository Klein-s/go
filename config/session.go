package config

import "github.com/klein/go-mvc/pkg/config"

func init()  {
	config.Add("session", config.StrMap{

		//目前只支持 cookie

		"driver":config.Env("SESSION_DRIVER", "cookie"),

		//会话的 cookie 名称
		"session_name":config.Env("SESSION_NAME", "blog-session"),
	})
}
