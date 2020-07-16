package handler

import (
	"context"

	"github.com/shadracnicholas/home-automation/libraries/go/database"
	"github.com/shadracnicholas/home-automation/libraries/go/oops"
	"github.com/shadracnicholas/home-automation/libraries/go/slog"
	scenedef "github.com/shadracnicholas/home-automation/service.scene/def"
	"github.com/shadracnicholas/home-automation/service.scene/domain"
)

// DeleteScene deletes a scene and associated actions
func (c *Controller) DeleteScene(ctx context.Context, body *scenedef.DeleteSceneRequest) (*scenedef.DeleteSceneResponse, error) {
	if body.SceneId == 0 {
		return nil, oops.BadRequest("scene_id empty")
	}

	// Delete the scene
	if err := database.Delete(&domain.Scene{}, body.SceneId); err != nil {
		return nil, err
	}

	slog.Infof("Deleted scene %d", body.SceneId)
	return &scenedef.DeleteSceneResponse{}, nil
}
