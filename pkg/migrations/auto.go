package migrations

import (
	"context"

	"github.com/opoccomaxao/gopkg/pkg/utils/automigrate"
	"github.com/opoccomaxao/perchance-image-lab/pkg/models"
	"gorm.io/gorm"
)

// AutoMigrate doesn't work correctly with foreign keys in duplicate tables,
// so we recreate all relations in current package.

type GenModelType struct {
	models.GenModelType
}

type GenModel struct {
	models.GenModel

	Type *GenModelType `gorm:"foreignKey:TypeID;references:ID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
}

type APIKey struct {
	models.APIKey

	Model *GenModel `gorm:"foreignKey:ModelID;references:ID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
}

func AutoMigrate(
	ctx context.Context,
	db *gorm.DB,
) error {
	migrator := db.WithContext(ctx).Migrator()

	err := automigrate.CreateTables(
		migrator,
		&GenModelType{},
		&GenModel{},
		&APIKey{},
	)
	if err != nil {
		return err
	}

	return nil
}
