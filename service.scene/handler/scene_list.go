package handler

import (
	"context"

	"github.com/shadracnicholas/home-automation/libraries/go/database"
	scenedef "github.com/shadracnicholas/home-automation/service.scene/def"
	"github.com/shadracnicholas/home-automation/service.scene/domain"
)

// ListScenes lists all scenes in the database
func (c *Controller) ListScenes(ctx context.Context, body *scenedef.ListScenesRequest) (*scenedef.ListScenesResponse, error) {
	where := make(map[string]interface{})
	if body.OwnerId > 0 {
		where["owner_id"] = body.OwnerId
	}

	var scenes []*domain.Scene
	if err := database.Find(&scenes, where); err != nil {
		return nil, err
	}

	protos := make([]*scenedef.Scene, len(scenes))
	for i, s := range scenes {
		protos[i] = s.ToProto()
	}

	return &scenedef.ListScenesResponse{
		Scenes: protos,
	}, nil
}
