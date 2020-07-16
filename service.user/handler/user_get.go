package handler

import (
	"context"

	"github.com/shadracnicholas/home-automation/libraries/go/database"
	"github.com/shadracnicholas/home-automation/libraries/go/oops"
	userdef "github.com/shadracnicholas/home-automation/service.user/def"
)

// GetUser reads a user by ID
func (c *Controller) GetUser(ctx context.Context, body *userdef.GetUserRequest) (*userdef.GetUserResponse, error) {
	user := &userdef.User{}
	if err := database.Find(user, body.UserId); err != nil {
		return nil, oops.WithMessage(err, "failed to find")
	}

	return &userdef.GetUserResponse{
		User: user,
	}, nil
}
