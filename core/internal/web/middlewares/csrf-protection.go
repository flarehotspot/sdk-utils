package middlewares

import (
	"core/internal/config"
	"fmt"
	"net/http"

	"github.com/gorilla/csrf"
)

var CsrfMiddleware func(http.Handler) http.Handler

func init() {
	appcfg, err := config.ReadApplicationConfig()
	if err != nil {
		panic(fmt.Errorf("Failed to read application config: %v", err))
	}
	CsrfMiddleware = csrf.Protect([]byte(appcfg.Secret))
}
