package userMapper

import (
	userDomain "github.com/quinntas/go-rest-template/api/user/domain"
	userValueObjects "github.com/quinntas/go-rest-template/api/user/domain/valueObjects"
	userModel "github.com/quinntas/go-rest-template/api/user/model"
	"github.com/quinntas/go-rest-template/internal/api/types"
	"github.com/quinntas/go-rest-template/internal/api/utils/timeUtils"
)

func ToDomain(model userModel.UserModel) userDomain.UserDomain {
	createdAt, _ := timeUtils.StringToTime(model.CreatedAt)
	updatedAt, _ := timeUtils.StringToTime(model.UpdatedAt)

	return userDomain.NewUserDomain(
		types.NewID(model.ID),
		types.NewUUID(model.PID),
		types.NewDate(createdAt),
		types.NewDate(updatedAt),
		userValueObjects.NewEmail(model.Email),
		userValueObjects.NewPassword(model.Password),
		types.NewID(model.RoleId),
	)
}
