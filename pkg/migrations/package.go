package migrations

import (
	"context"

	"github.com/opoccomaxao/gopkg/pkg/services/gormdb"
)

func MakePackage(
	appCtx context.Context,
	db *gormdb.Package,
) error {
	err := AutoMigrate(
		appCtx,
		db.DB,
	)
	if err != nil {
		return err
	}

	return nil
}
