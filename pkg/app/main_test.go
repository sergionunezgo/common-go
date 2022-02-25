package app_test

import (
	"log"
	"testing"

	"github.com/sergionunezgo/go-reuse/v2/pkg/app"
	"github.com/sergionunezgo/go-reuse/v2/pkg/http"
)

func initService(cfg *app.Config) (app.Service, error) {
	httpSrv := http.NewHttpService(cfg.Port)
	return httpSrv, nil
}

func TestRunApp(t *testing.T) {
	_ = app.Create(nil, initService)

	log.Print("app created")
}
