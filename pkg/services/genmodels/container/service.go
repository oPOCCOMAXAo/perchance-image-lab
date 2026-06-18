package container

import (
	"context"

	"github.com/opoccomaxao/perchance-image-lab/pkg/services/genmodels/structs"
)

type Container struct {
	modelsByCode map[string]structs.GenModel
}

func NewContainer() *Container {
	return &Container{
		modelsByCode: make(map[string]structs.GenModel),
	}
}

func (c *Container) Register(
	ctx context.Context,
	model structs.GenModel,
) error {
	c.modelsByCode[model.Code()] = model

	return nil
}
