package app_test

import (
	"log"
	"testing"

	"github.com/sergionunezgo/go-reuse/pkg/app"
	"github.com/sergionunezgo/go-reuse/pkg/service"
	"github.com/sergionunezgo/go-reuse/pkg/service/http"
)

func initService(cfg *service.Config) (service.Service, error) {
	httpSrv := http.NewService(cfg.Port)
	return httpSrv, nil
}

func TestRunApp(t *testing.T) {
	_ = app.Create(nil, initService)
	log.Print("app created")
}
