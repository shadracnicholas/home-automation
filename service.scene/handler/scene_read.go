package handler

import (
	"context"

	"github.com/shadracnicholas/home-automation/libraries/go/database"
	"github.com/shadracnicholas/home-automation/libraries/go/oops"
	scenedef "github.com/shadracnicholas/home-automation/service.scene/def"
	"github.com/shadracnicholas/home-automation/service.scene/domain"
)

// ReadScene returns the scene with the given ID
func (c *Controller) ReadScene(ctx context.Context, body *scenedef.ReadSceneRequest) (*scenedef.ReadSceneResponse, error) {
	scene := &domain.Scene{}
	if err := database.Find(&scene, body.SceneId); err != nil {
		return nil, oops.WithMessage(err, "failed to find")
	}

	return &scenedef.ReadSceneResponse{
		Scene: scene.ToProto(),
	}, nil
}
