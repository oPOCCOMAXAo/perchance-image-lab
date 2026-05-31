package app

import (
	"context"
	"os"
	"os/signal"

	"github.com/opoccomaxao/gopkg/pkg/services/ginserver"
	"github.com/opoccomaxao/gopkg/pkg/services/gormdb"
	"github.com/opoccomaxao/gopkg/pkg/services/lifecycle"
	"github.com/opoccomaxao/gopkg/pkg/services/logger"
	"github.com/opoccomaxao/gopkg/pkg/services/tasks"
	"github.com/opoccomaxao/perchance-image-lab/pkg/api"
	"github.com/opoccomaxao/perchance-image-lab/pkg/config"
	"github.com/opoccomaxao/perchance-image-lab/pkg/migrations"
)

func Run() error {
	appCtx, appCancelCause := context.WithCancelCause(context.Background())
	defer appCancelCause(nil)

	appCtx, cancel := signal.NotifyContext(appCtx, os.Interrupt)
	defer cancel()

	config, err := config.Load()
	if err != nil {
		return err
	}

	logger := logger.MakePackage(config.Logger)

	lifecycle := lifecycle.MakePackage(
		config.Lifecycle,
		appCancelCause,
		logger,
	)

	db, err := gormdb.MakePackage(
		config.DB,
		logger,
	)
	if err != nil {
		logger.Error(err)

		return err
	}

	err = migrations.AutoMigrate(appCtx, db.DB)
	if err != nil {
		logger.Error(err)

		return err
	}

	tasks := tasks.MakePackage(
		config.Tasks,
		lifecycle,
		logger,
	)

	server := ginserver.MakePackage(
		config.Server,
		lifecycle,
		logger,
	)

	err = api.MakePackage(
		config.API,
		server,
		tasks,
	)
	if err != nil {
		logger.Error(err)

		return err
	}

	err = lifecycle.Service.Serve(appCtx, appCancelCause)
	if err != nil {
		logger.Error(err)

		return err
	}

	<-appCtx.Done()

	err = lifecycle.Service.Shutdown(appCtx)
	if err != nil {
		logger.Error(err)

		return err
	}

	return nil
}
