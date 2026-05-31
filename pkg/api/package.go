package api

import (
	"github.com/opoccomaxao/gopkg/pkg/services/ginserver"
	"github.com/opoccomaxao/gopkg/pkg/services/tasks"
	"github.com/opoccomaxao/perchance-image-lab/pkg/api/internal"
	"github.com/opoccomaxao/perchance-image-lab/pkg/api/middleware"
	"github.com/opoccomaxao/perchance-image-lab/pkg/api/swagger"
	"github.com/opoccomaxao/perchance-image-lab/pkg/api/system"
)

type Config struct{}

func MakePackage(
	config Config,
	server *ginserver.Package,
	tasks *tasks.Package,
) error {
	mw := middleware.NewMiddleware(
		tasks.Service,
	)

	server.Engine.Use(mw.Init)

	regs := []internal.EndpointsService{
		system.New(),
		swagger.New(),
	}
	for _, reg := range regs {
		err := reg.Register(server.Router)
		if err != nil {
			return err
		}
	}

	return nil
}
