package server

import (
	"github.com/dimfeld/httptreemux"
	"mb_api/internal/pkg/router"
	"mb_api/internal/pkg/settings"
	"sync"

	prodncatsDelivery "mb_api/internal/pkg/prodncats/delivery"
)

var routesMap = map[string][]settings.MapHandler{
	"/api/getprods": {{
		Type:         "GET",
		Handler:      prodncatsDelivery.GetDelivery().GetProducts,
		CORS:         true,
		CSRF:         false,
	}},
	"/api/getcats": {{
		Type:         "GET",
		Handler:      prodncatsDelivery.GetDelivery().GetCategories,
		CORS:         true,
		CSRF:         false,
	}},
	"/api/getpairs": {{
		Type:         "GET",
		Handler:      prodncatsDelivery.GetDelivery().GetPairs,
		CORS:         true,
		CSRF:         false,
	}},
}

// Env variables which must to be set before running server
var Secrets = []string{
	"POSTGRES_DB",
	"POSTGRES_PASSWORD",
	"POSTGRES_USER",
}

var doOnce sync.Once
var conf settings.ServerSettings

func GetConfig() *settings.ServerSettings {
	doOnce.Do(func() {
		conf = settings.ServerSettings{
			Port:   3001,
			Ip:     "0.0.0.0",
			Routes: routesMap,
		}
		settings.SecureSettings = settings.GlobalSecure{
			CORSMethods: "",
			CORSMap:     map[string]struct{}{},
			AllowedHosts: map[string]struct{}{
				"http://localhost":           {},
				"http://localhost:8080":      {},
				"http://localhost:3000":      {},
				"http://127.0.0.1":           {},
				"http://127.0.0.1:8080":      {},
				"http://127.0.0.1:3000":      {},
			},
			// referring to https://security.stackexchange.com/questions/6957/length-of-csrf-token
			// it's correct length of CSRF token for Base64 (in bytes)
			CSRFTokenLen: 20,
			CSRFTokenTTL: 1, // one hour
			EnableCSRF:   true,
		}
		settings.UseCaseConf = settings.GlobalConfig{
			PageLimit: 10,
			InHDD:     true,
		}
		conf.InitSecure(&settings.SecureSettings)
		conf.InitConf(&settings.UseCaseConf)
		router.InitRouter(&conf, httptreemux.New())
	})
	return &conf
}