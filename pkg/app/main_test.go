package app_test

import (
	"log"
	"testing"

	"github.com/sergionunezgo/go-reuse/pkg/app"
	"github.com/sergionunezgo/go-reuse/pkg/http_srv"
)

func initService(cfg *app.Config) (app.Service, error) {
	httpSrv := http_srv.NewService(cfg.Port)
	return httpSrv, nil
}

func TestRunApp(t *testing.T) {
	_ = app.Create(nil, initService)
	log.Print("app created")
}
